package auth

import (
	database "buycoins-raven/internal/pkg/db/firestore"
	"buycoins-raven/internal/pkg/jwt"
	"buycoins-raven/internal/users"
	"context"
	"log"
	"net/http"
)

type contextKey struct {
	name string
}

var userCtxKey = &contextKey{"user"}

func Middleware() func(handler http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			header := r.Header.Get("Authorization")
			log.Println(header)
			//allow unauthenticated users
			if header == "" {
				next.ServeHTTP(w, r)
				return
			}

			tokenStr := header
			username, err := jwt.ParseToken(tokenStr)
			log.Println(username)
			if err != nil {
				http.Error(w, "Invalid token", http.StatusForbidden)
				return
			}

			user := users.User{UserName: username}

			client, ctx, err := database.ConnectToFirebase()
			log.Println("connecting to firestore.....")
			if err != nil {
				log.Fatal(err)
			}
			id, err := users.GetUserIdByUsername(username, ctx, *client)
			if err != nil {
				next.ServeHTTP(w, r)
				return
			}
			user.ID = id
			log.Println(user)
			ctxx := context.WithValue(r.Context(), userCtxKey, &user)

			r = r.WithContext(ctxx)
			next.ServeHTTP(w, r)
		})
	}
}

// ForContext finds the user from the context. REQUIRES Middleware to have run.
func ForContext(ctx context.Context) *users.User {
	raw, _ := ctx.Value(userCtxKey).(*users.User)
	return raw
}
