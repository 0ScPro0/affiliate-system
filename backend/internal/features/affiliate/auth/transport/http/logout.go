package auth_transport_http

import (
	"net/http"

	core_errors "github.com/0ScPro0/affiliate-system/internal/core/errors"
	"github.com/0ScPro0/affiliate-system/internal/core/logger"
	core_http_middleware "github.com/0ScPro0/affiliate-system/internal/core/transport/http/middleware"
	core_http_response "github.com/0ScPro0/affiliate-system/internal/core/transport/http/response"
)

// Logout godoc
// @Summary Logout user
// @Description Invalidate user refresh token
// @Tags auth
// @Accept json
// @Produce json
// @Success 200 {string} string "Logout successful"
// @Failure 401 {object} core_http_response.ErrorResponse "Unauthorized"
// @Failure 500 {object} core_http_response.ErrorResponse "Internal server error"
// @Router /logout [post]
func (h *AuthHTTPHandler) Logout(rw http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	log := logger.FromContext(ctx)
	responseHandler := core_http_response.NewHTTPResponseHandler(log, rw)

	// Check user auth 
	user := core_http_middleware.GetUserClaims(ctx)
	if user == nil {
		responseHandler.ErrorResponse(core_errors.ErrUnauthorized, "unauthorized")
		return
	}

	// Get user id from claims (stored as float64 from JWT)
	userIDFloat, ok := user["id"].(float64)
	if !ok || userIDFloat <= 0 {
		responseHandler.ErrorResponse(core_errors.ErrNotFound, "user not found")
		return
	}
	userID := int(userIDFloat)

	// Logout
	if err := h.authService.Logout(ctx, userID); err != nil {
		responseHandler.ErrorResponse(err, "failed to logout user")
		return
	}

	// Response
	responseHandler.JSONResponse("ok", http.StatusOK)
}