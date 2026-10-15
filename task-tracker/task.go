package main

import "time"

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

func (t *Task) touch() {
	t.UpdatedAt = time.Now()
}

func (t *Task) updateDescription(description string) {
	t.Description = description
	t.touch()
}

func (t *Task) updateStatus(status Status) {
	t.Status = status
	t.touch()
}

func newTask(id int, description string) Task {
	now := time.Now()
	return Task{
		ID:          int(now.UnixMilli()),
		Description: description,
		Status:      "todo",
		CreatedAt:   now,
		UpdatedAt:   now,
	}
}
