package main

import (
	"errors"
	"slices"
)

var ErrTaskNotFound = errors.New("task not found")

type Tasks []Task

func (ts Tasks) nextID() int {
	if len(ts) == 0 {
		return 1
	}

	maxID := ts[0].ID
	for _, t := range ts {
		maxID = max(maxID, t.ID)
	}

	return maxID
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

func (ts *Tasks) add(description string) {
	t := createTask(description)
	*ts = append(*ts, t)
}
