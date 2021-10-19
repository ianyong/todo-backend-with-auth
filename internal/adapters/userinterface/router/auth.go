package router

import (
	"context"
	"net/http"
	"strings"

	"github.com/ianyong/todo-backend/internal/adapters/userinterface/api"
	"github.com/ianyong/todo-backend/internal/contextkeys"
	"github.com/ianyong/todo-backend/internal/errors/externalerrors"
	"github.com/ianyong/todo-backend/internal/services"
)

func authMiddleware(s *services.Services) func(http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			authHeader := r.Header.Get("Authorization")
			authHeaderTokens := strings.Split(authHeader, "Bearer ")
			if len(authHeaderTokens) < 2 {
				api.ServeHTTPError(w, &externalerrors.AuthenticationError{})
				return
			}

			accessToken := authHeaderTokens[1]
			claims, err := s.JwtManager.Verify(accessToken)
			if err != nil {
				api.ServeHTTPError(w, &externalerrors.AuthenticationError{})
				return
			}

			ctx := context.WithValue(r.Context(), contextkeys.UserEmail, claims.Email)
			r = r.WithContext(ctx)

			next.ServeHTTP(w, r)
		})
	}
}
