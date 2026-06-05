package city_transport_http

import (
	"net/http"
	"strconv"

	core_errors "github.com/0ScPro0/affiliate-system/internal/core/errors"
	"github.com/0ScPro0/affiliate-system/internal/core/logger"
	core_http_middleware "github.com/0ScPro0/affiliate-system/internal/core/transport/http/middleware"
	core_http_response "github.com/0ScPro0/affiliate-system/internal/core/transport/http/response"
)

func (h *CityHTTPHandler) DeleteCity(rw http.ResponseWriter, r *http.Request) {
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

	// Delete city
	if err := h.cityService.DeleteCity(ctx, cityID); err != nil{
		responseHandler.ErrorResponse(err, "failed to delete city")
	}

	// Response
	responseHandler.JSONResponse(nil, http.StatusCreated)
}