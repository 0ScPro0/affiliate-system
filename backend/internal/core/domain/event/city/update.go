package city

import (
	"encoding/json"

	entitypkg "github.com/0ScPro0/affiliate-system/internal/core/domain/entity"
	eventpkg "github.com/0ScPro0/affiliate-system/internal/core/domain/event"
)

// UpdatedEvent is published when an existing city is updated.
type UpdatedEvent struct {
	eventpkg.BaseEvent
	Payload entitypkg.City `json:"payload"`
}

// NewUpdatedEvent creates a new UpdatedEvent.
func NewUpdatedEvent(city entitypkg.City) UpdatedEvent {
	return UpdatedEvent{
		BaseEvent: eventpkg.BaseEvent{Type: EventUpdated},
		Payload:   city,
	}
}

// MarshalJSON serializes the event to JSON.
func (e UpdatedEvent) MarshalJSON() ([]byte, error) {
	type Alias UpdatedEvent
	return json.Marshal(struct {
		Alias
		Type eventpkg.EventType `json:"type"`
	}{
		Alias: Alias(e),
		Type:  e.Type,
	})
}