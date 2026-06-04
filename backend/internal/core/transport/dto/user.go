package core_transport_dto

import "time"

type CreateUserRequest struct {
	UserName     *string `json:"username" validate:"omitempty,min=1,max=50"`
	Email        string  `json:"email" validate:"required,email,max=100"`
	PasswordHash string  `json:"password_hash" validate:"required,min=1,max=255"`
	IsAdmin      bool    `json:"is_admin"`
}

type UpdateUserRequest struct {
	ID           int     `json:"id" validate:"required"`
	UserName     *string `json:"username" validate:"omitempty,min=1,max=50"`
	Email        *string `json:"email" validate:"omitempty,email,max=100"`
	PasswordHash *string `json:"password_hash" validate:"omitempty,min=1,max=255"`
	IsAdmin      *bool   `json:"is_admin"`
}

type UserResponse struct {
	ID           int       `json:"id"`
	UserName     *string   `json:"username"`
	Email        string    `json:"email"`
	IsAdmin      bool      `json:"is_admin"`
	CreatedAt    time.Time `json:"created_at"`
}