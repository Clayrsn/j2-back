package middlewares

import (
	"context"
	"errors"
	"log"
	"net/http"
	"strings"

	"j2-api/models"
	"j2-api/services"

	firebase "firebase.google.com/go/v4"
	"firebase.google.com/go/v4/auth"
	"go.mongodb.org/mongo-driver/mongo"
	"google.golang.org/api/option"
)

type MiddlewareResources struct {
	UsersService services.UsersService
}

func (mr MiddlewareResources) FirebaseAuthMiddleware(client *auth.Client) func(next http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

			authHeader := r.Header.Get("Authorization")
			if authHeader == "" {
				log.Printf("Error: auth empty\n")
				http.Error(w, http.StatusText(http.StatusUnauthorized), http.StatusUnauthorized)
				return
			}

			tokenString := strings.TrimSpace(strings.Replace(authHeader, "Bearer", "", 1))
			if tokenString == "" {
				log.Printf("Error: token empty\n")
				http.Error(w, http.StatusText(http.StatusUnauthorized), http.StatusUnauthorized)
				return
			}

			user, err := mr.UsersService.GetUserByToken(tokenString)
			if err != nil {
				if errors.Is(err, mongo.ErrNoDocuments) {
					token, err := client.VerifyIDToken(context.Background(), tokenString)
					if err != nil {
						log.Printf("Error: verifying token: %v\n", err)
						http.Error(w, http.StatusText(http.StatusUnauthorized), http.StatusUnauthorized)
						return
					}

					userID := token.UID

					user, err = mr.UsersService.GetUser(userID)
					if err != nil {
						if errors.Is(err, mongo.ErrNoDocuments) {
							userCreate := models.UserCreate{
								Name:         token.Claims["name"].(string),
								Email:        token.Claims["email"].(string),
								CurrentToken: tokenString,
							}
							user, err = mr.UsersService.CreateUser(userCreate)
							if err != nil {
								log.Printf("Error: creating user: %v\n", err)
								http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
								return
							}
						} else {
							log.Printf("Error: checking user: %v\n", err)
							http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
							return
						}
					} else {
						user.CurrentToken = tokenString
						_, err = mr.UsersService.UpdateUserToken(user.ID.String(), user.CurrentToken)
						if err != nil {
							log.Printf("Error: updating user token: %v\n", err)
							http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
							return
						}
					}
				} else {
					log.Printf("Error: checking user token: %v\n", err)
					http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
					return
				}
			}

			ctx := context.WithValue(r.Context(), "user", user)

			next.ServeHTTP(w, r.WithContext(ctx))
		})
	}
}

func InitializeFirebaseApp(credentialsFilePath string) (*auth.Client, error) {
	ctx := context.Background()
	opt := option.WithCredentialsFile(credentialsFilePath)
	app, err := firebase.NewApp(ctx, nil, opt)
	if err != nil {
		return nil, err
	}

	client, err := app.Auth(ctx)
	if err != nil {
		return nil, err
	}

	return client, nil
}
