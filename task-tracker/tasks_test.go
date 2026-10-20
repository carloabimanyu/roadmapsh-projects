package main

import (
	"slices"
	"testing"
)

func taskIDs(ts Tasks) []int {
	result := []int{}
	for _, t := range ts {
		result = append(result, t.ID)
	}

	return result
}

func TestFilterByStatus(t *testing.T) {
	data := []struct {
		description string
		status      Status
	}{
		{description: "Buy groceries", status: StatusInProgress},
		{description: "Prepare for lunch", status: StatusDone},
		{description: "Play video games", status: StatusInProgress},
	}

	ts := Tasks{}
	for _, d := range data {
		id, err := ts.add(d.description)
		if err != nil {
			t.Fatalf("setup: Tasks.add() error = %v, want nil", err)
		}

		if err = ts.updateStatus(id, d.status); err != nil {
			t.Fatalf("setup: Tasks.updateStatus() error = %v, want nil", err)
		}
	}

	tests := []struct {
		name   string
		status Status
		want   []int
	}{
		{name: "todo", status: StatusTodo, want: nil},
		{name: "in-progress", status: StatusInProgress, want: []int{1, 3}},
		{name: "done", status: StatusDone, want: []int{2}},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			filteredTs := ts.filterByStatus(tt.status)
			filteredIDs := taskIDs(filteredTs)
			if !slices.Equal(filteredIDs, tt.want) {
				t.Errorf("filterByStatus(%q) IDs = %v, want %v", tt.status, filteredIDs, tt.want)
			}
		})
	}
}
