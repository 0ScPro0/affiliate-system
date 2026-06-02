package partner

import (
	"encoding/json"

	entitypkg "github.com/0ScPro0/affiliate-system/internal/core/domain/entity"
	eventpkg "github.com/0ScPro0/affiliate-system/internal/core/domain/event"
)

// DeletedEvent is published when a partner is deleted.
type DeletedEvent struct {
	eventpkg.BaseEvent
	Payload entitypkg.Partner `json:"payload"`
}

// NewDeletedEvent creates a new DeletedEvent.
func NewDeletedEvent(partner entitypkg.Partner) DeletedEvent {
	return DeletedEvent{
		BaseEvent: eventpkg.BaseEvent{Type: EventDeleted},
		Payload:   partner,
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