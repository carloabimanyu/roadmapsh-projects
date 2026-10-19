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

func TestSaveLoad_RoundTrip(t *testing.T) {
	ts := Tasks{}
	ts.add("Learn Go")
	ts.add("Create unit test")

	dir := t.TempDir()
	path := filepath.Join(dir, "tasks.json")
	s := NewFileStore(path)

	if err := s.Save(ts); err != nil {
		t.Fatalf("Save() error = %v, want nil", err)
	}

	tsLoaded, err := s.Load()
	if err != nil {
		t.Fatalf("Load() error = %v, want nil", err)
	}

	if len(ts) != len(tsLoaded) {
		t.Fatalf("len(tsLoaded) = %d, want %d", len(tsLoaded), len(ts))
	}

	for i := range ts {
		if ts[i].ID != tsLoaded[i].ID {
			t.Errorf("tsLoaded[%d].ID = %d, want %d", i, tsLoaded[i].ID, ts[i].ID)
		}
		if ts[i].Description != tsLoaded[i].Description {
			t.Errorf("tsLoaded[%d].Description = %q, want %q", i, tsLoaded[i].Description, ts[i].Description)
		}
		if ts[i].Status != tsLoaded[i].Status {
			t.Errorf("tsLoaded[%d].Status = %q, want %q", i, tsLoaded[i].Status, ts[i].Status)
		}
		if !ts[i].CreatedAt.Equal(tsLoaded[i].CreatedAt) {
			t.Errorf("tsLoaded[%d].CreatedAt = %v, want %v", i, tsLoaded[i].CreatedAt, ts[i].CreatedAt)
		}
		if !ts[i].UpdatedAt.Equal(tsLoaded[i].UpdatedAt) {
			t.Errorf("tsLoaded[%d].UpdatedAt = %v, want %v", i, tsLoaded[i].UpdatedAt, ts[i].UpdatedAt)
		}
	}
}
