package main

import (
	"os"
	"path/filepath"
	"testing"
)

func TestLoad_FileNotExist(t *testing.T) {
	dir := t.TempDir()
	path := filepath.Join(dir, "tasks.json")
	s := NewFileStore(path)

	ts, err := s.Load()

	if err != nil {
		t.Fatalf("Load() error = %v, want nil", err)
	}

	if len(ts) != 0 {
		t.Fatalf("len(tasks) = %d, want 0", len(ts))
	}
}

func TestLoad_EmptyFile(t *testing.T) {
	dir := t.TempDir()
	path := filepath.Join(dir, "tasks.json")
	s := NewFileStore(path)

	if err := os.WriteFile(path, nil, 0o644); err != nil {
		t.Fatalf("setup: write empty file: %v", err)
	}

	ts, err := s.Load()

	if err != nil {
		t.Fatalf("Load() error = %v, want nil", err)
	}

	if len(ts) != 0 {
		t.Fatalf("len(tasks) = %d, want 0", len(ts))
	}
}

func TestLoad_CorruptFile(t *testing.T) {
	dir := t.TempDir()
	path := filepath.Join(dir, "tasks.json")
	s := NewFileStore(path)

	data := []byte("{not json")
	if err := os.WriteFile(path, data, 0o644); err != nil {
		t.Fatalf("setup: write corrupt file: %v", err)
	}

	_, err := s.Load()

	if err == nil {
		t.Fatal("Load() error = nil, want non-nil")
	}
}
