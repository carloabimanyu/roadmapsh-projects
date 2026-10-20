package main

import (
	"errors"
	"fmt"
	"os"
	"strconv"
)

const tasksFile = "tasks.json"

func run(args []string) error {
	if len(args) == 0 {
		return errors.New("usage: task-tracker <command> [arguments]")
	}

	s := NewFileStore(tasksFile)

	switch args[0] {
	case "add":
		if len(args) != 2 {
			return errors.New("usage: task-tracker add <description>")
		}

		var id int

		err := s.Update(func(ts *Tasks) error {
			var err error
			id, err = ts.add(args[1])
			return err
		})
		if err != nil {
			return err
		}

		fmt.Printf("Task added successfully (ID: %d)\n", id)

	case "list":
		if len(args) > 2 {
			return errors.New("usage: task-tracker list [todo|in-progress|done]")
		}

		var status Status
		var err error

		if len(args) == 2 {
			status, err = parseStatus(args[1])
			if err != nil {
				return err
			}
		}

		ts, err := s.Load()
		if err != nil {
			return err
		}

		if status != "" {
			ts = ts.filterByStatus(status)
		}

		if len(ts) == 0 {
			if status != "" {
				fmt.Printf("No tasks with status %q.\n", status)
				return nil
			}

			fmt.Println("No tasks found.")
			return nil
		}

		fmt.Printf("%3s %-12s%s\n", "ID", "STATUS", "DESCRIPTION")
		for i := range ts {
			fmt.Printf("%3d %-12s%s\n", ts[i].ID, ts[i].Status, ts[i].Description)
		}

	case "delete":
		if len(args) != 2 {
			return errors.New("usage: task-tracker delete <id>")
		}

		id, err := strconv.Atoi(args[1])
		if err != nil {
			return fmt.Errorf("cannot convert %q to number, id must be a number", args[1])
		}

		err = s.Update(func(ts *Tasks) error {
			return ts.delete(id)
		})
		if err != nil {
			return err
		}

		fmt.Printf("Task deleted successfully (ID: %d)\n", id)

	default:
		return fmt.Errorf("unknown command %q", args[0])
	}

	return nil
}

func main() {
	if err := run(os.Args[1:]); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}
