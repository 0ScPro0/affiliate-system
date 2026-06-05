package city_transport_http

import (
	"context"
	"net/http"

	domain "github.com/0ScPro0/affiliate-system/internal/core/domain/entity"
	core_transport_dto "github.com/0ScPro0/affiliate-system/internal/core/transport/dto"
	core_http_server "github.com/0ScPro0/affiliate-system/internal/core/transport/http/server"
)

type CityHTTPHandler struct {
	cityService CityService
}

type CityService interface {
	CreateCity(
		ctx context.Context,
		city core_transport_dto.CreateCityRequest,
	) (domain.City, error)
	
	GetCityByID(
		ctx context.Context,
		id int,
	) (domain.City, error)

	UpdateCity(
		ctx context.Context,
		city core_transport_dto.UpdateCityRequest,
	) (domain.City, error)

	DeleteCity(
		ctx context.Context,
		id int,
	) error
}

func NewCityHTTPHandler(
	cityService CityService,
) *CityHTTPHandler {
	return &CityHTTPHandler{
		cityService: cityService,
	}
}

func (h *CityHTTPHandler) Routes() []core_http_server.Route {
	return []core_http_server.Route{
		{
			Method: http.MethodPost,
			Path: "/cities",
			Handler: h.CreateCity,
		},
		{
			Method: http.MethodGet,
			Path: "/cities/{id}",
			Handler: h.GetCityByID,
		},
		{
			Method: http.MethodPatch,
			Path: "/cities/{id}",
			Handler: h.UpdateCity,
		},
		{
			Method: http.MethodDelete,
			Path: "/cities/{id}",
			Handler: h.DeleteCity,
		},
	}
}