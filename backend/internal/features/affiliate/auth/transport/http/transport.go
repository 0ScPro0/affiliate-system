package auth_transport_http

import (
	"context"
	"net/http"

	core_transport_dto "github.com/0ScPro0/affiliate-system/internal/core/transport/dto"
	core_http_server "github.com/0ScPro0/affiliate-system/internal/core/transport/http/server"
)

type AuthHTTPHandler struct {
	authService AuthService
}

type AuthService interface {
	Register(
		ctx context.Context,
		req core_transport_dto.RegisterRequest,
	) (core_transport_dto.RegisterResponse, error)

	Login(
		ctx context.Context,
		req core_transport_dto.LoginRequest,
	) (core_transport_dto.LoginResponse, error)

	Logout(
		ctx context.Context,
		userID int,
	) error

	RefreshToken(
		ctx context.Context,
		req core_transport_dto.RefreshTokenRequest,
	) (core_transport_dto.RefreshTokenResponse, error)
}

func NewAuthHTTPHandler(
	authService AuthService,
) *AuthHTTPHandler {
	return &AuthHTTPHandler{
		authService: authService,
	}
}

func (h *AuthHTTPHandler) Routes() []core_http_server.Route {
	return []core_http_server.Route{
		{
			Method: http.MethodPost,
			Path: "/register",
			Handler: h.Register,
		},
		{
			Method: http.MethodPost,
			Path: "/login",
			Handler: h.Login,
		},
		{
			Method: http.MethodPost,
			Path: "/logout",
			Handler: h.Logout,
		},
		{
			Method: http.MethodPost,
			Path: "/refresh_token",
			Handler: h.RefreshToken,
		},

	}
}