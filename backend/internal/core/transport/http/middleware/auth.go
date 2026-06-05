package core_http_middleware

import (
	"context"
	"net/http"
	"strings"

	"github.com/0ScPro0/affiliate-system/internal/core/config"
	core_errors "github.com/0ScPro0/affiliate-system/internal/core/errors"
	logger "github.com/0ScPro0/affiliate-system/internal/core/logger"
	core_security "github.com/0ScPro0/affiliate-system/internal/core/security"
	core_http_response "github.com/0ScPro0/affiliate-system/internal/core/transport/http/response"
)

// contextKey for storing user claims in request context.
const userClaimsKey = "user_claims"

// Auth returns a middleware that validates Bearer token from Authorization header
// and stores the token claims in the request context.
func Auth(cfg *config.Config) Middleware {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			log := logger.FromContext(r.Context())
			responseHandler := core_http_response.NewHTTPResponseHandler(log, w)

			// Skip auth for swagger docs (any path starting with /docs)
			if strings.HasPrefix(r.URL.Path, "/docs") {
				next.ServeHTTP(w, r)
				return
			}

			authHeader := r.Header.Get("Authorization")
			if authHeader == "" {
				responseHandler.ErrorResponse(
					core_errors.ErrUnauthorized,
					"missing Authorization header",
				)
				return
			}

			parts := strings.SplitN(authHeader, " ", 2)
			if len(parts) != 2 || !strings.EqualFold(parts[0], "Bearer") {
				responseHandler.ErrorResponse(
					core_errors.ErrUnauthorized,
					"invalid Authorization header format, expected 'Bearer <token>'",
				)
				return
			}

			tokenString := parts[1]

			claims, err := core_security.VerifyToken(cfg, tokenString)
			if err != nil {
				responseHandler.ErrorResponse(
					core_errors.ErrUnauthorized,
					"invalid or expired token",
				)
				return
			}

			// The sub fields (id, email, role) are stored at the root level of claims
			// by createToken, so we store the entire claims map in context.
			ctx := context.WithValue(r.Context(), userClaimsKey, map[string]any(claims))
			next.ServeHTTP(w, r.WithContext(ctx))
		})
	}
}

// GetUserClaims extracts user claims from context.
func GetUserClaims(ctx context.Context) map[string]any {
	claims, ok := ctx.Value(userClaimsKey).(map[string]any)
	if !ok {
		return nil
	}
	return claims
}