package city

import (
	eventpkg "github.com/0ScPro0/affiliate-system/internal/core/domain/event"
)

const (
	EventCreated eventpkg.EventType = "city.created"
	EventUpdated eventpkg.EventType = "city.updated"
	EventDeleted eventpkg.EventType = "city.deleted"
)