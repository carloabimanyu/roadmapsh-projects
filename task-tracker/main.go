package main

import (
	"fmt"
	"slices"
	"time"
)

type Task struct {
	ID          int64
	Description string
	Status      string
	CreatedAt   time.Time
	UpdatedAt   time.Time
}

func createTask(description string) Task {
	now := time.Now()
	return Task{now.UnixMilli(), description, "todo", now, now}
}

func (task *Task) updateUpdateAt() {
	task.UpdatedAt = time.Now()
}

func (task *Task) updateDescription(description string) {
	task.Description = description
}

func (task *Task) updateStatus(status string) {
	task.Status = status
}

type TaskSlice []Task

func (taskSlice *TaskSlice) addTask(task Task) {
	*taskSlice = append(*taskSlice, task)
}

func (taskSlice *TaskSlice) findTaskById(id int64) Task {
	i := slices.IndexFunc(*taskSlice, func(task Task) bool {
		return task.ID == id
	})
	if i == -1 {
		return nil
	}

	return taskSlice[i]
}

func (taskSlice *TaskSlice) deleteTask(id int64) bool {
	i := slices.IndexFunc(*taskSlice, func(task Task) bool {
		return task.ID == id
	})

	if i == -1 {
		return false
	}

	*taskSlice = slices.Delete(*taskSlice, i, i+1)
	return true
}

func (taskSlice *TaskSlice) updateDescription(id int64, description string) {
	task := *taskSlice.findTaskById()
}

func main() {
	// Empty slice
	tasks := TaskSlice{}

	fmt.Println(tasks)

	// Add task
	task1 := createTask("Buy groceries")
	tasks.addTask(task1)

	fmt.Println(tasks)

	// Update description

}
