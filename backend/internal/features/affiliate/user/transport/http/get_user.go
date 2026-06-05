package user_transport_http

import (
	"net/http"
	"strconv"

	core_errors "github.com/0ScPro0/affiliate-system/internal/core/errors"
	"github.com/0ScPro0/affiliate-system/internal/core/logger"
	core_transport_dto "github.com/0ScPro0/affiliate-system/internal/core/transport/dto"
	core_http_middleware "github.com/0ScPro0/affiliate-system/internal/core/transport/http/middleware"
	core_http_response "github.com/0ScPro0/affiliate-system/internal/core/transport/http/response"
)

// GetUserByID godoc
// @Summary Get user by ID
// @Description Get user by its unique identifier
// @Tags users
// @Accept json
// @Produce json
// @Param id path int true "User ID"
// @Success 200 {object} core_transport_dto.UserResponse "User found"
// @Failure 400 {object} core_http_response.ErrorResponse "Bad request"
// @Failure 401 {object} core_http_response.ErrorResponse "Unauthorized"
// @Failure 403 {object} core_http_response.ErrorResponse "Forbidden"
// @Failure 404 {object} core_http_response.ErrorResponse "Not found"
// @Failure 500 {object} core_http_response.ErrorResponse "Internal server error"
// @Router /users/{id} [get]
func (h *UserHTTPHandler) GetUserByID(rw http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	log := logger.FromContext(ctx)
	responseHandler := core_http_response.NewHTTPResponseHandler(log, rw)

	// Check user auth
	user := core_http_middleware.GetUserClaims(ctx)
	if user == nil {
		responseHandler.ErrorResponse(core_errors.ErrUnauthorized, "unauthorized")
		return
	}

	// Get id from path
	id := r.PathValue("id")
	if id == "" {
		responseHandler.ErrorResponse(core_errors.ErrInvalidArgument, "user id is required")
		return
	}
	userID, err := strconv.Atoi(id)
	if err != nil {
		responseHandler.ErrorResponse(core_errors.ErrInvalidArgument, "invalid user ID format")
		return
	}

	// Get user
	foundUser, err := h.userService.GetUserByID(ctx, userID)
	if err != nil {
		responseHandler.ErrorResponse(err, "failed to get user")
		return
	}

	// Response
	response := core_transport_dto.UserResponse{
		ID:        foundUser.ID,
		UserName:  foundUser.UserName,
		Email:     foundUser.Email,
		IsAdmin:   foundUser.IsAdmin,
		CreatedAt: foundUser.CreatedAt,
	}
	responseHandler.JSONResponse(response, http.StatusOK)
}