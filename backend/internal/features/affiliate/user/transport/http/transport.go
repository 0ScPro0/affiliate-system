package user_transport_http

import (
	"context"
	"net/http"

	domain "github.com/0ScPro0/affiliate-system/internal/core/domain/entity"
	core_transport_dto "github.com/0ScPro0/affiliate-system/internal/core/transport/dto"
	core_http_server "github.com/0ScPro0/affiliate-system/internal/core/transport/http/server"
)

type UserHTTPHandler struct {
	userService UserService
}

type UserService interface {
	CreateUser(
		ctx context.Context,
		user core_transport_dto.CreateUserRequest,
	) (domain.User, error)

	GetUserByID(
		ctx context.Context,
		id int,
	) (domain.User, error)

	UpdateUser(
		ctx context.Context,
		id int,
		user core_transport_dto.UpdateUserRequest,
	) (domain.User, error)

	DeleteUser(
		ctx context.Context,
		id int,
	) error
}

func NewUserHTTPHandler(
	userService UserService,
) *UserHTTPHandler {
	return &UserHTTPHandler{
		userService: userService,
	}
}

func (h *UserHTTPHandler) Routes() []core_http_server.Route {
	return []core_http_server.Route{
		{
			Method:  http.MethodPost,
			Path:    "/users",
			Handler: h.CreateUser,
		},
		{
			Method:  http.MethodGet,
			Path:    "/users/{id}",
			Handler: h.GetUserByID,
		},
		{
			Method:  http.MethodPatch,
			Path:    "/users/{id}",
			Handler: h.UpdateUser,
		},
		{
			Method:  http.MethodDelete,
			Path:    "/users/{id}",
			Handler: h.DeleteUser,
		},
	}
}