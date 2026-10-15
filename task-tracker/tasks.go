package main

import (
	"errors"
	"slices"
)

var ErrTaskNotFound = errors.New("task not found")

type Tasks []Task

func (ts Tasks) nextID() int {
	maxID := 0
	for _, t := range ts {
		maxID = max(maxID, t.ID)
	}

	return maxID + 1
}

func (ts Tasks) indexOf(id int) (int, error) {
	i := slices.IndexFunc(ts, func(t Task) bool {
		return t.ID == id
	})

	if i == -1 {
		return i, ErrTaskNotFound
	}
	return i, nil
}

func (ts *Tasks) delete(id int) error {
	i, err := ts.indexOf(id)
	if err != nil {
		return err
	}

	*ts = slices.Delete(*ts, i, i+1)
	return nil
}

func (ts *Tasks) add(description string) int {
	id := ts.nextID()
	t := newTask(id, description)
	*ts = append(*ts, t)
	return id
}

func (ts *Tasks) updateDescription(id int, description string) error {
	i, err := ts.indexOf(id)
	if err != nil {
		return err
	}

	(*ts)[i].updateDescription(description)
	return nil
}

func (ts *Tasks) updateStatus(id int, status Status) error {
	i, err := ts.indexOf(id)
	if err != nil {
		return err
	}

	(*ts)[i].updateStatus(status)
	return nil
}
