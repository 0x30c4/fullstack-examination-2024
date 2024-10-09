package model

import "time"

// Todo is the model for the todo endpoint.
type Todo struct {
	ID        int `gorm:"primaryKey"`
	Task      string
	Status    Status
	Priority  Priority
	CreatedAt time.Time `gorm:"autoCreateTime"`
	UpdatedAt time.Time `gorm:"autoUpdateTime"`
}

// NewTodo returns a new instance of the todo model.
func NewTodo(task string, priority Priority) *Todo {
	return &Todo{
		Task:     task,
		Priority: priority,
		Status:   Created,
	}
}

// NewUpdateTodo returns a new instance of the todo model for updating.
func NewUpdateTodo(id int, task string, status Status) *Todo {
	return &Todo{
		ID:     id,
		Task:   task,
		Status: status,
	}
}

// Status represents the status of a task.
type Status string

// Priority represents the priority of a task.
type Priority string

const (
	// Created is the status for a created task.
	Created = Status("created")
	// Processing is the status for a processing task.
	Processing = Status("processing")
	// Done is the status for a done task.
	Done = Status("done")
)

// StatusMap is a map of task status.
var StatusMap = map[Status]bool{
	Created:    true,
	Processing: true,
	Done:       true,
}

const (
	// Low is the priority for a created task.
	Low = Priority("low")
	// Medium the priority for a created task.
	Medium = Priority("medium")
	// High the priority for a created task.
	High = Priority("high")
)

// PriorityMap is a map of task status.
var PriorityMap = map[Priority]bool{
	Low:    true,
	Medium: true,
	High:   true,
}
