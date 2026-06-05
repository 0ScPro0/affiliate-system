package auth_transport_http

import (
	"net/http"

	"github.com/0ScPro0/affiliate-system/internal/core/logger"
	core_transport_dto "github.com/0ScPro0/affiliate-system/internal/core/transport/dto"
	core_http_request "github.com/0ScPro0/affiliate-system/internal/core/transport/http/request"
	core_http_response "github.com/0ScPro0/affiliate-system/internal/core/transport/http/response"
)

// RefreshToken godoc
// @Summary Refresh access token
// @Description Refresh access token using refresh token
// @Tags auth
// @Accept json
// @Produce json
// @Param request body core_transport_dto.RefreshTokenRequest true "RefreshToken request body"
// @Success 200 {object} core_transport_dto.RefreshTokenResponse "Token refreshed successfully"
// @Failure 400 {object} core_http_response.ErrorResponse "Bad request"
// @Failure 401 {object} core_http_response.ErrorResponse "Unauthorized"
// @Failure 500 {object} core_http_response.ErrorResponse "Internal server error"
// @Router /refresh_token [post]
func (h *AuthHTTPHandler) RefreshToken(rw http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	log := logger.FromContext(ctx)
	responseHandler := core_http_response.NewHTTPResponseHandler(log, rw)

	// Decode and validate request
	var request core_transport_dto.RefreshTokenRequest
	if err := core_http_request.DecodeAndValidateRequest(r, &request); err != nil {
		responseHandler.ErrorResponse(err, "failed to decode and validate HTTP request")
		return
	}

	// Refresh token
	resp, err := h.authService.RefreshToken(ctx, request)
	if err != nil {
		responseHandler.ErrorResponse(err, "failed to refresh token")
		return
	}

	// Response
	responseHandler.JSONResponse(resp, http.StatusOK)
}