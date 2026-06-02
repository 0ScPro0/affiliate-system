package core_database_models

import (
	"time"
)

type PartnerModel struct {
	ID          int
	Name        string
	Description *string
	CreatedAt   time.Time
}
