package core_database_models

import (
	"time"
)

type CityModel struct {
	ID        int
	Name      string
	CreatedAt time.Time
}
