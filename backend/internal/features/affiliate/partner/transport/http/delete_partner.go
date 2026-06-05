package partner_transport_http

import (
	"net/http"
	"strconv"

	core_errors "github.com/0ScPro0/affiliate-system/internal/core/errors"
	"github.com/0ScPro0/affiliate-system/internal/core/logger"
	core_http_middleware "github.com/0ScPro0/affiliate-system/internal/core/transport/http/middleware"
	core_http_response "github.com/0ScPro0/affiliate-system/internal/core/transport/http/response"
)

// DeletePartner godoc
// @Summary Delete partner
// @Description Delete partner by ID
// @Tags partners
// @Accept json
// @Produce json
// @Param id path int true "Partner ID"
// @Success 204 "Partner deleted successfully"
// @Failure 400 {object} core_http_response.ErrorResponse "Bad request"
// @Failure 401 {object} core_http_response.ErrorResponse "Unauthorized"
// @Failure 403 {object} core_http_response.ErrorResponse "Forbidden"
// @Failure 404 {object} core_http_response.ErrorResponse "Not found"
// @Failure 500 {object} core_http_response.ErrorResponse "Internal server error"
// @Router /partners/{id} [delete]
func (h *PartnerHTTPHandler) DeletePartner(rw http.ResponseWriter, r *http.Request) {
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

	// Get id from path
	id := r.PathValue("id")
	if id == "" {
		responseHandler.ErrorResponse(core_errors.ErrInvalidArgument, "partner id is required")
		return
	}
	partnerID, err := strconv.Atoi(id)
	if err != nil {
		responseHandler.ErrorResponse(core_errors.ErrInvalidArgument, "invalid partner ID format")
		return
	}

	// Delete partner
	if err := h.partnerService.DeletePartner(ctx, partnerID); err != nil {
		responseHandler.ErrorResponse(err, "failed to delete partner")
		return
	}

	// Response
	responseHandler.JSONResponse(nil, http.StatusNoContent)
}