package authentication_check

import (
	"context"
	"net/http"
	"strings"

	"github.com/mulan17/project-user-service/pkg/token"
)

func Authenticate(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		authToken := r.Header.Get("Authorization")
		
		if authToken == "" {
			http.Error(w, "Not authorized", http.StatusUnauthorized)
			return
		}

		// Remove "Bearer " prefix if present
		if strings.HasPrefix(authToken, "Bearer ") {
			authToken = strings.TrimPrefix(authToken, "Bearer ")
		}

		admin, userId, err := token.VerifyToken(authToken)

		if err != nil {
			http.Error(w, "Not authorized", http.StatusUnauthorized)
			return
		}

		// Add user information to request context
		ctx := context.WithValue(r.Context(), "userId", userId)
		ctx = context.WithValue(ctx, "admin", admin)

		// Call the next handler with the new context
		next.ServeHTTP(w, r.WithContext(ctx))
	})
}
