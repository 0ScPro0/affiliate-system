package user

import (
	eventpkg "github.com/0ScPro0/affiliate-system/internal/core/domain/event"
)

const (
	EventCreated eventpkg.EventType = "user.created"
	EventUpdated eventpkg.EventType = "user.updated"
	EventDeleted eventpkg.EventType = "user.deleted"
)