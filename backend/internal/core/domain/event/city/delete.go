package city

import (
	"encoding/json"

	entitypkg "github.com/0ScPro0/affiliate-system/internal/core/domain/entity"
	eventpkg "github.com/0ScPro0/affiliate-system/internal/core/domain/event"
)

// DeletedEvent is published when a city is deleted.
type DeletedEvent struct {
	eventpkg.BaseEvent
	Payload entitypkg.City `json:"payload"`
}

// NewDeletedEvent creates a new DeletedEvent.
func NewDeletedEvent(city entitypkg.City) DeletedEvent {
	return DeletedEvent{
		BaseEvent: eventpkg.BaseEvent{Type: EventDeleted},
		Payload:   city,
	}
}

// MarshalJSON serializes the event to JSON.
func (e DeletedEvent) MarshalJSON() ([]byte, error) {
	type Alias DeletedEvent
	return json.Marshal(struct {
		Alias
		Type eventpkg.EventType `json:"type"`
	}{
		Alias: Alias(e),
		Type:  e.Type,
	})
}