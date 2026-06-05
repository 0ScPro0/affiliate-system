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

func (h *CityHTTPHandler) UpdateCity(rw http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	log := logger.FromContext(ctx)
	responseHandler := core_http_response.NewHTTPResponseHandler(log, rw)

	// Check user auth and permissions
	user := core_http_middleware.GetUserClaims(ctx)
	if user == nil {
		responseHandler.ErrorResponse(core_errors.ErrUnauthorized, "Unauthorized")
	}
	if isAdmin, ok := user["is_admin"].(bool); !ok || !isAdmin {
		responseHandler.ErrorResponse(core_errors.ErrForbidden, "Not enough permissions")
	}

	// Validate request
	var request core_transport_dto.UpdateCityRequest
	if err := core_http_request.DecodeAndValidateRequest(r, request); err != nil {
		responseHandler.ErrorResponse(err, "failed to decode and validate HTTP request")
	}

	// Update city
	city, err := h.cityService.UpdateCity(ctx, request)
	if err != nil {
		responseHandler.ErrorResponse(err, "failed to update city")
	}

	// Response
	response := core_transport_dto.CityResponse{
		ID: city.ID,
		Name: city.Name,
		CreatedAt: city.CreatedAt,
	}
	responseHandler.JSONResponse(response, http.StatusCreated)
}