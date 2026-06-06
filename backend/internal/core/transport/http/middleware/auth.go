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

// protectedEndpoints defines which routes require authentication.
// The key is "METHOD path" format (path includes /api/v1/ prefix).
var protectedEndpoints = map[string]bool{
	// City
	"POST /api/v1/cities":       true,
	"PATCH /api/v1/cities/{id}": true,
	"DELETE /api/v1/cities/{id}": true,

	// Partner
	"POST /api/v1/partners":       true,
	"PATCH /api/v1/partners/{id}": true,
	"DELETE /api/v1/partners/{id}": true,

	// Category
	"POST /api/v1/categories":       true,
	"PATCH /api/v1/categories/{id}": true,
	"DELETE /api/v1/categories/{id}": true,

	// User (all protected)
	"POST /api/v1/users":       true,
	"GET /api/v1/users/{id}":   true,
	"PATCH /api/v1/users/{id}": true,
	"DELETE /api/v1/users/{id}": true,

	// Auth
	"POST /api/v1/logout": true,
}

// isProtected checks if the given method and path match a protected endpoint.
// It performs an exact match first, then falls back to pattern matching
// for paths with path parameters (e.g., /api/v1/cities/{id}).
func isProtected(method, path string) bool {
	key := method + " " + path
	if protectedEndpoints[key] {
		return true
	}

	// Try pattern matching for paths with {id} parameter.
	// Split the path into segments and compare against known patterns.
	segments := strings.Split(strings.Trim(path, "/"), "/")
	if len(segments) < 4 {
		return false
	}

	// Build a pattern key by replacing the last segment with {id}
	// if it looks like a numeric ID.
	key = method + " /" + segments[0] + "/" + segments[1] + "/" + segments[2] + "/{id}"
	return protectedEndpoints[key]
} 

// Auth returns a middleware that validates Bearer token from Authorization header
// for protected endpoints only. Public endpoints are passed through without authentication.
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

			// Check if the endpoint requires authentication
			if !isProtected(r.Method, r.URL.Path) {
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

			var tokenString string
			parts := strings.SplitN(authHeader, " ", 2)
			if len(parts) == 2 && strings.EqualFold(parts[0], "Bearer") {
				tokenString = parts[1]
			} else {
				// If no "Bearer" prefix, treat the entire header value as the token
				tokenString = authHeader
			}

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