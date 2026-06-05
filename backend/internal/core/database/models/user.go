package core_database_models

import "time"

type UserModel struct {
	ID                   int
	UserName             *string
	Email                string
	PasswordHash         string
	IsAdmin              bool
	CreatedAt            time.Time
	RefreshToken         *string
	RefreshTokenExpiresAt *time.Time
}