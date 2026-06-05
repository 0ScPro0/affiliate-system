package auth_transport_http

import (
	"net/http"

	"github.com/0ScPro0/affiliate-system/internal/core/logger"
	core_transport_dto "github.com/0ScPro0/affiliate-system/internal/core/transport/dto"
	core_http_request "github.com/0ScPro0/affiliate-system/internal/core/transport/http/request"
	core_http_response "github.com/0ScPro0/affiliate-system/internal/core/transport/http/response"
)

// Register godoc
// @Summary Register user
// @Description Register a new user in the system
// @Tags auth
// @Accept json
// @Produce json
// @Param request body core_transport_dto.RegisterRequest true "Register request body"
// @Success 201 {object} core_transport_dto.RegisterResponse "User registered successfully"
// @Failure 400 {object} core_http_response.ErrorResponse "Bad request"
// @Failure 409 {object} core_http_response.ErrorResponse "Conflict"
// @Failure 500 {object} core_http_response.ErrorResponse "Internal server error"
// @Router /register [post]
func (h *AuthHTTPHandler) Register(rw http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	log := logger.FromContext(ctx)
	responseHandler := core_http_response.NewHTTPResponseHandler(log, rw)

	// Decode and validate request
	var request core_transport_dto.RegisterRequest
	if err := core_http_request.DecodeAndValidateRequest(r, &request); err != nil {
		responseHandler.ErrorResponse(err, "failed to decode and validate HTTP request")
		return
	}

	// Register user
	resp, err := h.authService.Register(ctx, request)
	if err != nil {
		responseHandler.ErrorResponse(err, "failed to register user")
		return
	}

	// Response
	responseHandler.JSONResponse(resp, http.StatusCreated)
}