package authentication_check

import (
	"context"
	"net/http"
	"strings"

	"github.com/mulan17/project-user-service/pkg/token"
)

type contextKey string

func Authenticate(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		authToken := r.Header.Get("Authorization")

		if authToken == "" {
			http.Error(w, "Not authorized", http.StatusUnauthorized)
			return
		}

		authToken = strings.TrimPrefix(authToken, "Bearer ")

		email, role, userId, err := token.VerifyToken(authToken)

		if err != nil {
			http.Error(w, "Not authorized", http.StatusUnauthorized)
			return
		}

		if role != "admin" {
			http.Error(w, "Not authorized as admin", http.StatusUnauthorized)
			return
		}

		ctx := context.WithValue(r.Context(), contextKey("userId"), userId)
		ctx = context.WithValue(ctx, contextKey("role"), role)
		ctx = context.WithValue(ctx, contextKey("email"), email)

		next.ServeHTTP(w, r.WithContext(ctx))
	})
}
