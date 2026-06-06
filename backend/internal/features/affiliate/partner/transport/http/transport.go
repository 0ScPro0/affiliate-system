package partner_transport_http

import (
	"context"
	"net/http"

	domain "github.com/0ScPro0/affiliate-system/internal/core/domain/entity"
	core_transport_dto "github.com/0ScPro0/affiliate-system/internal/core/transport/dto"
	core_http_server "github.com/0ScPro0/affiliate-system/internal/core/transport/http/server"
)

type PartnerHTTPHandler struct {
	partnerService PartnerService
}

type PartnerService interface {
	CreatePartner(
		ctx context.Context,
		partner core_transport_dto.CreatePartnerRequest,
	) (domain.Partner, error)

	GetPartnerByID(
		ctx context.Context,
		id int,
	) (domain.Partner, error)

	UpdatePartner(
		ctx context.Context,
		id int,
		partner core_transport_dto.UpdatePartnerRequest,
	) (domain.Partner, error)

	DeletePartner(
		ctx context.Context,
		id int,
	) error
}

func NewPartnerHTTPHandler(
	partnerService PartnerService,
) *PartnerHTTPHandler {
	return &PartnerHTTPHandler{
		partnerService: partnerService,
	}
}

func (h *PartnerHTTPHandler) Routes() []core_http_server.Route {
	return []core_http_server.Route{
		{
			Method:  http.MethodPost,
			Path:    "/partners",
			Handler: h.CreatePartner,
		},
		{
			Method:  http.MethodGet,
			Path:    "/partners/{id}",
			Handler: h.GetPartnerByID,
		},
		{
			Method:  http.MethodPatch,
			Path:    "/partners/{id}",
			Handler: h.UpdatePartner,
		},
		{
			Method:  http.MethodDelete,
			Path:    "/partners/{id}",
			Handler: h.DeletePartner,
		},
	}
}