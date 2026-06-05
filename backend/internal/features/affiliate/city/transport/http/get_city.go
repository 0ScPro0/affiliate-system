package city_transport_http

import (
	"net/http"
	"strconv"

	core_errors "github.com/0ScPro0/affiliate-system/internal/core/errors"
	"github.com/0ScPro0/affiliate-system/internal/core/logger"
	core_transport_dto "github.com/0ScPro0/affiliate-system/internal/core/transport/dto"
	core_http_response "github.com/0ScPro0/affiliate-system/internal/core/transport/http/response"
)

func (h *CityHTTPHandler) GetCityByID(rw http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	log := logger.FromContext(ctx)
	responseHandler := core_http_response.NewHTTPResponseHandler(log, rw)

	// Dont need to auth, public method

	// Get id from path
	id := r.PathValue("id")
	if id == "" {
		responseHandler.ErrorResponse(core_errors.ErrInvalidArgument, "city id is required")
	}
	cityID, err := strconv.Atoi(id)
	if err != nil {
		responseHandler.ErrorResponse(core_errors.ErrInvalidArgument, "invalid city ID format")
		return
	}

	// Get city
	city, err := h.cityService.GetCityByID(ctx, cityID)
	if err != nil {
		responseHandler.ErrorResponse(err, "failed to get city")
	}

	// Response
	response := core_transport_dto.CityResponse{
		ID: city.ID,
		Name: city.Name,
		CreatedAt: city.CreatedAt,
	}
	responseHandler.JSONResponse(response, http.StatusCreated)
}