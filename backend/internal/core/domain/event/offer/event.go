package offer

import (
	eventpkg "github.com/0ScPro0/affiliate-system/internal/core/domain/event"
)

const (
	EventCreated eventpkg.EventType = "offer.created"
	EventUpdated eventpkg.EventType = "offer.updated"
	EventDeleted eventpkg.EventType = "offer.deleted"
)