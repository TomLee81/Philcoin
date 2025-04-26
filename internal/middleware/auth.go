package middleware

import (
	"context"
	"net/http"

	"philcoin/internal/utils"
)

// Auth checks JWT token in Authorization header and sets userID
func Auth(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		tokenStr := r.Header.Get("Authorization")
		if tokenStr == "" {
			utils.RespondError(w, http.StatusUnauthorized, "Missing token")
			return
		}
		claims, err := utils.ValidateToken(tokenStr, utils.GetJWTSecret())
		if err != nil {
			utils.RespondError(w, http.StatusUnauthorized, "Invalid token")
			return
		}
		ctx := context.WithValue(r.Context(), "userID", claims.Subject)
		next.ServeHTTP(w, r.WithContext(ctx))
	})
}
