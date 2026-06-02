package partner

import (
	eventpkg "github.com/0ScPro0/affiliate-system/internal/core/domain/event"
)

const (
	EventCreated eventpkg.EventType = "partner.created"
	EventUpdated eventpkg.EventType = "partner.updated"
	EventDeleted eventpkg.EventType = "partner.deleted"
)