package core_database_models

import (
	"time"
)

type CategoryModel struct {
	ID          int
	Name        string
	Description *string
	CreatedAt   time.Time
}
