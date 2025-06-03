package middleware

import (
	"net/http"
	"strings"

	"github.com/minab/internship-backend/internal/util"
)

func JWTAuth(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		authHeader := r.Header.Get("Authorization")
		if !strings.HasPrefix(authHeader, "Bearer ") {
			http.Error(w, "Missing or invalid Authorization header", http.StatusUnauthorized)
			return
		}
		tokenStr := strings.TrimPrefix(authHeader, "Bearer ")
		claims, err := util.ParseJWT(tokenStr)
		if err != nil {
			http.Error(w, "Invalid token", http.StatusUnauthorized)
			return
		}
		// Optionally set claims in context for downstream handlers
		r = r.WithContext(util.ContextWithClaims(r.Context(), claims))
		next.ServeHTTP(w, r)
	})
}
