package category

import (
	eventpkg "github.com/0ScPro0/affiliate-system/internal/core/domain/event"
)

const (
	EventCreated eventpkg.EventType = "category.created"
	EventUpdated eventpkg.EventType = "category.updated"
	EventDeleted eventpkg.EventType = "category.deleted"
)