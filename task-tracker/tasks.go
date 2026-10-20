package main

import (
	"errors"
	"fmt"
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
		return i, fmt.Errorf("%w for id = %d", ErrTaskNotFound, id)
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

func (ts *Tasks) add(description string) (int, error) {
	id := ts.nextID()
	t, err := newTask(id, description)
	if err != nil {
		return 0, err
	}
	*ts = append(*ts, t)
	return id, nil
}

func (ts *Tasks) updateDescription(id int, description string) error {
	i, err := ts.indexOf(id)
	if err != nil {
		return err
	}

	return (*ts)[i].updateDescription(description)
}

func (ts *Tasks) updateStatus(id int, status Status) error {
	i, err := ts.indexOf(id)
	if err != nil {
		return err
	}

	(*ts)[i].updateStatus(status)
	return nil
}

func (ts Tasks) filterByStatus(status Status) Tasks {
	result := Tasks{}
	for _, t := range ts {
		if t.Status == status {
			result = append(result, t)
		}
	}

	return result
}
