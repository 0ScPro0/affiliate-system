package core_database_models

import (
	"time"
)

type OfferModel struct {
	ID          int
	PartnerID   int
	CategoryID  int
	CityID      int
	Name        string
	Description *string
	CreatedAt   time.Time
	ExpireAt 	time.Time
}
