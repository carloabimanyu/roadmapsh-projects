package main

import (
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"os"
	"path/filepath"
)

type FileStore struct {
	path string
}

func NewFileStore(path string) FileStore {
	return FileStore{path: path}
}

func (s FileStore) Load() (Tasks, error) {
	ts := Tasks{}

	reader, err := os.Open(s.path)
	if errors.Is(err, os.ErrNotExist) {
		return ts, nil
	}
	if err != nil {
		return nil, err
	}
	defer reader.Close()

	decoder := json.NewDecoder(reader)
	err = decoder.Decode(&ts)
	if errors.Is(err, io.EOF) {
		return ts, nil
	}
	if err != nil {
		return nil, fmt.Errorf("decode %s: %w", s.path, err)
	}

	return ts, nil
}

func (s FileStore) Save(ts Tasks) error {
	writer, err := os.CreateTemp(filepath.Dir(s.path), "tasks-*.json")
	if err != nil {
		return err
	}

	// Clean up the temp file if anything fails before the rename.
	// On success the file has already been renamed, so this is a no-op
	// and its error is safe to ignore.
	defer os.Remove(writer.Name())

	encoder := json.NewEncoder(writer)
	encoder.SetIndent("", "  ")
	if err = encoder.Encode(ts); err != nil {
		writer.Close()
		return fmt.Errorf("encode %s: %w", s.path, err)
	}

	if err := writer.Close(); err != nil {
		return err
	}

	if err := os.Rename(writer.Name(), s.path); err != nil {
		return err
	}

	return nil
}

func (s FileStore) Update(fn func(ts *Tasks) error) error {
	ts, err := s.Load()
	if err != nil {
		return err
	}

	if err = fn(&ts); err != nil {
		return err
	}

	return s.Save(ts)
}
