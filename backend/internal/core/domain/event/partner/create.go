package partner

import (
	"encoding/json"

	entitypkg "github.com/0ScPro0/affiliate-system/internal/core/domain/entity"
	eventpkg "github.com/0ScPro0/affiliate-system/internal/core/domain/event"
)

// CreatedEvent is published when a new partner is created.
type CreatedEvent struct {
	eventpkg.BaseEvent
	Payload entitypkg.Partner `json:"payload"`
}

// NewCreatedEvent creates a new CreatedEvent.
func NewCreatedEvent(partner entitypkg.Partner) CreatedEvent {
	return CreatedEvent{
		BaseEvent: eventpkg.BaseEvent{Type: EventCreated},
		Payload:   partner,
	}
}

// MarshalJSON serializes the event to JSON.
func (e CreatedEvent) MarshalJSON() ([]byte, error) {
	type Alias CreatedEvent
	return json.Marshal(struct {
		Alias
		Type eventpkg.EventType `json:"type"`
	}{
		Alias: Alias(e),
		Type:  e.Type,
	})
}