package middleware

import (
	"lms_system/internal/domain/common"
	"lms_system/utils"
	"net/http"
)

func AdminOnlyMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		userCtx := utils.GetUserFromContext(r.Context())
		if userCtx == nil {
			http.Error(w, "User context not found", http.StatusInternalServerError)
			return
		}

		if userCtx.Role != common.RoleAdmin {
			http.Error(w, "Admin access required", http.StatusForbidden)
			return
		}

		next.ServeHTTP(w, r)
	})
}
