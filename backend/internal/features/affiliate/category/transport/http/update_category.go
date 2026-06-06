package category_transport_http

import (
	"net/http"
	"strconv"

	core_errors "github.com/0ScPro0/affiliate-system/internal/core/errors"
	"github.com/0ScPro0/affiliate-system/internal/core/logger"
	core_transport_dto "github.com/0ScPro0/affiliate-system/internal/core/transport/dto"
	core_http_middleware "github.com/0ScPro0/affiliate-system/internal/core/transport/http/middleware"
	core_http_request "github.com/0ScPro0/affiliate-system/internal/core/transport/http/request"
	core_http_response "github.com/0ScPro0/affiliate-system/internal/core/transport/http/response"
)

// UpdateCategory godoc
// @Summary Update category
// @Description Update existing category by ID
// @Tags categories
// @Accept json
// @Produce json
// @Param id path int true "Category ID"
// @Param request body core_transport_dto.UpdateCategoryRequest true "UpdateCategory request body"
// @Success 200 {object} core_transport_dto.CategoryResponse "Category updated successfully"
// @Failure 400 {object} core_http_response.ErrorResponse "Bad request"
// @Failure 401 {object} core_http_response.ErrorResponse "Unauthorized"
// @Failure 403 {object} core_http_response.ErrorResponse "Forbidden"
// @Failure 404 {object} core_http_response.ErrorResponse "Not found"
// @Failure 500 {object} core_http_response.ErrorResponse "Internal server error"
// @Security BearerAuth
// @Router /categories/{id} [patch]
func (h *CategoryHTTPHandler) UpdateCategory(rw http.ResponseWriter, r *http.Request) {
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
		responseHandler.ErrorResponse(core_errors.ErrInvalidArgument, "category id is required")
		return
	}
	categoryID, err := strconv.Atoi(id)
	if err != nil {
		responseHandler.ErrorResponse(core_errors.ErrInvalidArgument, "invalid category ID format")
		return
	}

	// Validate request
	var request core_transport_dto.UpdateCategoryRequest
	if err := core_http_request.DecodeAndValidateRequest(r, &request); err != nil {
		responseHandler.ErrorResponse(err, "failed to decode and validate HTTP request")
		return
	}
	request.ID = categoryID

	// Update category
	category, err := h.categoryService.UpdateCategory(ctx, request)
	if err != nil {
		responseHandler.ErrorResponse(err, "failed to update category")
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