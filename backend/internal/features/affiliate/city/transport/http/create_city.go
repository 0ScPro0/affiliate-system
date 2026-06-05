package city_transport_http

import (
	"net/http"

	core_errors "github.com/0ScPro0/affiliate-system/internal/core/errors"
	"github.com/0ScPro0/affiliate-system/internal/core/logger"
	core_transport_dto "github.com/0ScPro0/affiliate-system/internal/core/transport/dto"
	core_http_middleware "github.com/0ScPro0/affiliate-system/internal/core/transport/http/middleware"
	core_http_request "github.com/0ScPro0/affiliate-system/internal/core/transport/http/request"
	core_http_response "github.com/0ScPro0/affiliate-system/internal/core/transport/http/response"
)

// CreateCity godoc
// @Summary Create city
// @Description Create new city in system
// @Tags cities
// @Accept json
// @Produce json
// @Param request body core_transport_dto.CreateCityRequest true "CreateCity request body"
// @Success 201 {object} core_transport_dto.CityResponse "City created successfully"
// @Failure 400 {object} core_http_response.ErrorResponse "Bad request"
// @Failure 401 {object} core_http_response.ErrorResponse "Unauthorized"
// @Failure 403 {object} core_http_response.ErrorResponse "Forbidden"
// @Failure 500 {object} core_http_response.ErrorResponse "Internal server error"
// @Router /cities [post]
func (h *CityHTTPHandler) CreateCity(rw http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	log := logger.FromContext(ctx)
	responseHandler := core_http_response.NewHTTPResponseHandler(log, rw)

	// Check user auth and permissions
	user := core_http_middleware.GetUserClaims(ctx)
	if user == nil {
		responseHandler.ErrorResponse(core_errors.ErrUnauthorized, "unauthorized")
		return
	}
	if isAdmin, ok := user["is_admin"].(bool); !ok || !isAdmin {
		responseHandler.ErrorResponse(core_errors.ErrForbidden, "not enough permissions")
		return
	}

	// Validate request
	var request core_transport_dto.CreateCityRequest
	if err := core_http_request.DecodeAndValidateRequest(r, &request); err != nil {
		responseHandler.ErrorResponse(err, "failed to decode and validate HTTP request")
		return
	}

	// Create city
	city, err := h.cityService.CreateCity(ctx, request)
	if err != nil {
		responseHandler.ErrorResponse(err, "failed to create city")
		return
	}

	// Response
	response := core_transport_dto.CityResponse{
		ID:        city.ID,
		Name:      city.Name,
		CreatedAt: city.CreatedAt,
	}
	responseHandler.JSONResponse(response, http.StatusCreated)
}