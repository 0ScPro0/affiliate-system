package domain

import "encoding/json"

// EventType represents the type of a domain event.
type EventType string

// DomainEvent is the interface that all domain events must implement.
type DomainEvent interface {
	// EventType returns the type of the event.
	EventType() EventType
	// MarshalJSON serializes the event to JSON.
	MarshalJSON() ([]byte, error)
}

// BaseEvent contains common fields for all domain events.
type BaseEvent struct {
	Type EventType `json:"type"`
}

// EventType returns the type of the event.
func (e BaseEvent) EventType() EventType {
	return e.Type
}

// MarshalJSON serializes the base event to JSON.
func (e BaseEvent) MarshalJSON() ([]byte, error) {
	return json.Marshal(e)
}