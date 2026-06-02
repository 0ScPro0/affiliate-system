package partner

import (
	"encoding/json"

	entitypkg "github.com/0ScPro0/affiliate-system/internal/core/domain/entity"
	eventpkg "github.com/0ScPro0/affiliate-system/internal/core/domain/event"
)

// UpdatedEvent is published when an existing partner is updated.
type UpdatedEvent struct {
	eventpkg.BaseEvent
	Payload entitypkg.Partner `json:"payload"`
}

// NewUpdatedEvent creates a new UpdatedEvent.
func NewUpdatedEvent(partner entitypkg.Partner) UpdatedEvent {
	return UpdatedEvent{
		BaseEvent: eventpkg.BaseEvent{Type: EventUpdated},
		Payload:   partner,
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