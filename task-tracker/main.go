package main

import (
	"errors"
	"fmt"
	"os"
)

func run(args []string) error {
	if len(args) == 0 {
		return errors.New("usage: task-tracker <command> [arguments]")
	}

	switch args[0] {
	case "add":
		if len(args) != 2 {
			return errors.New("usage: task-tracker add <description>")
		}

		s := NewFileStore("tasks.json")
		ts, err := s.Load()
		if err != nil {
			return err
		}

		id, err := ts.add(args[1])
		if err != nil {
			return err
		}

		err = s.Save(ts)
		if err != nil {
			return err
		}
		fmt.Printf("Task added successfully (ID: %d)\n", id)

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
