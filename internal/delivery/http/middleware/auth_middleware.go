package middleware

import (
	"context"
	"lms_system/internal/domain/common"
	"lms_system/internal/domain/entity"
	"net/http"
	"strings"
)

func AuthMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		authHeader := r.Header.Get("Authorization")
		if authHeader == "" {
			http.Error(w, "Authorization header required", http.StatusUnauthorized)
			return
		}

		tokenParts := strings.Split(authHeader, " ")
		if len(tokenParts) != 2 || tokenParts[0] != "Bearer" {
			http.Error(w, "Invalid authorization header format", http.StatusUnauthorized)
			return
		}

		token := tokenParts[1]
		userCtx, err := validateToken(token)
		if err != nil {
			http.Error(w, "Invalid token", http.StatusUnauthorized)
			return
		}

		ctx := context.WithValue(r.Context(), common.UserContextKey, userCtx)
		next.ServeHTTP(w, r.WithContext(ctx))
	})
}

func validateToken(token string) (*entity.UserContext, error) {
	switch token {
	case "admin-token":
		return &entity.UserContext{
			UserID: 1,
			Role:   common.RoleAdmin,
		}, nil
	case "user-token":
		return &entity.UserContext{
			UserID: 2,
			Role:   common.RoleUser,
		}, nil
	default:
		return nil, http.ErrNoCookie // Using as generic error
	}
}
