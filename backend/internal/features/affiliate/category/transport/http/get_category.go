package category_transport_http

import (
	"net/http"
	"strconv"

	core_errors "github.com/0ScPro0/affiliate-system/internal/core/errors"
	"github.com/0ScPro0/affiliate-system/internal/core/logger"
	core_transport_dto "github.com/0ScPro0/affiliate-system/internal/core/transport/dto"
	core_http_response "github.com/0ScPro0/affiliate-system/internal/core/transport/http/response"
)

// GetCategoryByID godoc
// @Summary Get category by ID
// @Description Get category by its unique identifier
// @Tags categories
// @Accept json
// @Produce json
// @Param id path int true "Category ID"
// @Success 200 {object} core_transport_dto.CategoryResponse "Category found"
// @Failure 400 {object} core_http_response.ErrorResponse "Bad request"
// @Failure 404 {object} core_http_response.ErrorResponse "Not found"
// @Failure 500 {object} core_http_response.ErrorResponse "Internal server error"
// @Router /categories/{id} [get]
func (h *CategoryHTTPHandler) GetCategoryByID(rw http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	log := logger.FromContext(ctx)
	responseHandler := core_http_response.NewHTTPResponseHandler(log, rw)

	// Dont need to auth, public method

	// Get id from path
	id := r.PathValue("id")
	if id == "" {
		responseHandler.ErrorResponse(core_errors.ErrInvalidArgument, "category id is required")
		return
	}
	categoryID, err := strconv.Atoi(id)
	if err != nil {
		responseHandler.ErrorResponse(core_errors.ErrInvalidArgument, "invalid category ID format")
		return
	}

	// Get category
	category, err := h.categoryService.GetCategoryByID(ctx, categoryID)
	if err != nil {
		responseHandler.ErrorResponse(err, "failed to get category")
		return
	}

	// Response
	response := core_transport_dto.CategoryResponse{
		ID:          category.ID,
		Name:        category.Name,
		Description: category.Description,
		CreatedAt:   category.CreatedAt,
	}
	responseHandler.JSONResponse(response, http.StatusOK)
}