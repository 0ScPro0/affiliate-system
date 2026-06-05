package category_transport_http

import (
	"context"
	"net/http"

	domain "github.com/0ScPro0/affiliate-system/internal/core/domain/entity"
	core_transport_dto "github.com/0ScPro0/affiliate-system/internal/core/transport/dto"
	core_http_server "github.com/0ScPro0/affiliate-system/internal/core/transport/http/server"
)

type CategoryHTTPHandler struct {
	categoryService CategoryService
}

type CategoryService interface {
	CreateCategory(
		ctx context.Context,
		category core_transport_dto.CreateCategoryRequest,
	) (domain.Category, error)

	GetCategoryByID(
		ctx context.Context,
		id int,
	) (domain.Category, error)

	UpdateCategory(
		ctx context.Context,
		category core_transport_dto.UpdateCategoryRequest,
	) (domain.Category, error)

	DeleteCategory(
		ctx context.Context,
		id int,
	) error
}

func NewCategoryHTTPHandler(
	categoryService CategoryService,
) *CategoryHTTPHandler {
	return &CategoryHTTPHandler{
		categoryService: categoryService,
	}
}

func (h *CategoryHTTPHandler) Routes() []core_http_server.Route {
	return []core_http_server.Route{
		{
			Method:  http.MethodPost,
			Path:    "/categories",
			Handler: h.CreateCategory,
		},
		{
			Method:  http.MethodGet,
			Path:    "/categories/{id}",
			Handler: h.GetCategoryByID,
		},
		{
			Method:  http.MethodPatch,
			Path:    "/categories/{id}",
			Handler: h.UpdateCategory,
		},
		{
			Method:  http.MethodDelete,
			Path:    "/categories/{id}",
			Handler: h.DeleteCategory,
		},
	}
}