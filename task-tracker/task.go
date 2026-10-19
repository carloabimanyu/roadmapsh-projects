package main

import (
	"errors"
	"strings"
	"time"
)

type Status string

const (
	StatusTodo       Status = "todo"
	StatusInProgress Status = "in-progress"
	StatusDone       Status = "done"
)

type Task struct {
	ID          int       `json:"id"`
	Description string    `json:"description"`
	Status      Status    `json:"status"`
	CreatedAt   time.Time `json:"createdAt"`
	UpdatedAt   time.Time `json:"updatedAt"`
}

var ErrEmptyDescription = errors.New("description cannot be empty")

func normalizeDescription(description string) (string, error) {
	normalizedDescription := strings.TrimSpace(description)
	if normalizedDescription == "" {
		return "", ErrEmptyDescription
	}
	return normalizedDescription, nil
}

func (t *Task) touch() {
	t.UpdatedAt = time.Now()
}

func (t *Task) updateDescription(description string) error {
	normalizedDescription, err := normalizeDescription(description)
	if err != nil {
		return err
	}

	t.Description = normalizedDescription
	t.touch()
	return nil
}

func (t *Task) updateStatus(status Status) {
	t.Status = status
	t.touch()
}

func newTask(id int, description string) (Task, error) {
	normalizedDescription, err := normalizeDescription(description)
	if err != nil {
		return Task{}, err
	}

	now := time.Now()
	return Task{
		ID:          id,
		Description: normalizedDescription,
		Status:      StatusTodo,
		CreatedAt:   now,
		UpdatedAt:   now,
	}, nil
}
