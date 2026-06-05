package partner_transport_http

import (
	"net/http"
	"strconv"

	core_errors "github.com/0ScPro0/affiliate-system/internal/core/errors"
	"github.com/0ScPro0/affiliate-system/internal/core/logger"
	core_transport_dto "github.com/0ScPro0/affiliate-system/internal/core/transport/dto"
	core_http_response "github.com/0ScPro0/affiliate-system/internal/core/transport/http/response"
)

// GetPartnerByID godoc
// @Summary Get partner by ID
// @Description Get partner by its unique identifier
// @Tags partners
// @Accept json
// @Produce json
// @Param id path int true "Partner ID"
// @Success 200 {object} core_transport_dto.PartnerResponse "Partner found"
// @Failure 400 {object} core_http_response.ErrorResponse "Bad request"
// @Failure 404 {object} core_http_response.ErrorResponse "Not found"
// @Failure 500 {object} core_http_response.ErrorResponse "Internal server error"
// @Router /partners/{id} [get]
func (h *PartnerHTTPHandler) GetPartnerByID(rw http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	log := logger.FromContext(ctx)
	responseHandler := core_http_response.NewHTTPResponseHandler(log, rw)

	// Dont need to auth, public method

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

	// Get partner
	partner, err := h.partnerService.GetPartnerByID(ctx, partnerID)
	if err != nil {
		responseHandler.ErrorResponse(err, "failed to get partner")
		return
	}

	// Response
	response := core_transport_dto.PartnerResponse{
		ID:          partner.ID,
		Name:        partner.Name,
		Description: partner.Description,
		CreatedAt:   partner.CreatedAt,
	}
	responseHandler.JSONResponse(response, http.StatusOK)
}