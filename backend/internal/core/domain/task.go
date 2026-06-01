package domain

import (
	"fmt"
	"time"

	core_errors "github.com/0ScPro0/affiliate-system/internal/core/errors"
)

// TaskStatus type
type TaskStatus int

const (
	NEW TaskStatus = iota
	IN_PROGRESS
	DONE
)

func (s TaskStatus) String() string {
	return [...]string{"new", "in_progress", "done"}[s]
}

// Domain entity Task
// Task has required fields Name and Status
// and optional field Description
type Task struct {
	ID          int        `json:"id"`
	Version     int        `json:"version"`
	CreatedAt   time.Time  `json:"created_at"`

	Name        string     `json:"name"`
	Description *string    `json:"description"`
	Status      TaskStatus `json:"status"`
}

func NewTask(
	id int,
	version int,
	created_at time.Time,
	name string,
	description *string,
	status TaskStatus,
) Task {
	return Task{
		ID: id,
		Version: version,
		CreatedAt: created_at,
		Name: name,
		Description: description,
		Status: status,
	}
}

func NewTaskUnitialized(
	name string,
	description *string,
	status TaskStatus,
) Task {
	return Task{
		ID: UnitializedID,
		Version: UnitializedVersion,
		CreatedAt: UnitializedCreatedAt,
		Name: name,
		Description: description,
		Status: status,
	}
}

func (t *Task) Validate() error {
	nameLength := len([]rune(t.Name))
	if nameLength < 1 || nameLength > 100 {
		return fmt.Errorf(
			"Invalid `Name` len: %d: %w", 
			nameLength, 
			core_errors.ErrInvalidArgument,
		)
	}

	if t.Description != nil {
		descriptionLength := len([]rune(*t.Description))
		if descriptionLength < 1 || descriptionLength > 1000 {
			return fmt.Errorf(
				"Invalid `Description` len: %d: %w", 
				descriptionLength, 
				core_errors.ErrInvalidArgument,
			)
		}
	}

	if t.Status < 0 || t.Status > 2 {
		return fmt.Errorf(
			"Invalid `Status`: %d: %w", 
			t.Status, 
			core_errors.ErrInvalidArgument,
		)
	}
	
	return nil
}
