package middleware

import (
	"course/src/controller"
	"net/http"
	"strings"
)

type AdminMiddleware controller.UserController

func InitAdminMiddleware(c *controller.UserController) *AdminMiddleware {
	return &AdminMiddleware{Repository: c.Repository}
}

func (m *AdminMiddleware) AdminMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		header := r.Header.Get("Authorization")
		if header == "" {
			http.Error(w, "Unauthorized", http.StatusUnauthorized)
			return
		}

		tokenParts := strings.Split(header, " ")
		if len(tokenParts) != 2 || tokenParts[0] != "Bearer" {
			http.Error(w, "Unauthorized", http.StatusUnauthorized)
			return
		}

		u, err := m.Repository.GetUserByToken(tokenParts[1])
		if u == nil || err != nil || !m.Repository.UserIsAdmin(u) {
			http.Error(w, "Unauthorized", http.StatusUnauthorized)
			return
		}

		next.ServeHTTP(w, r)
	})
}
