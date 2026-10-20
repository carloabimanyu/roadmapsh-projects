package main

import (
	"errors"
	"testing"
)

func TestNormalizeDescription(t *testing.T) {
	tests := []struct {
		name        string
		description string
		wantErr     error
		want        string
	}{
		{name: "valid", description: "Beli sayur", wantErr: nil, want: "Beli sayur"},
		{name: "empty string", description: "", wantErr: ErrEmptyDescription, want: ""},
		{name: "whitespace only", description: "    ", wantErr: ErrEmptyDescription, want: ""},
		{name: "newline only", description: "\n", wantErr: ErrEmptyDescription, want: ""},
		{name: "tab only", description: "\t", wantErr: ErrEmptyDescription, want: ""},
		{name: "valid with space", description: "  Beli sayur ", wantErr: nil, want: "Beli sayur"},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			normalizedDescription, err := normalizeDescription(tt.description)
			if !errors.Is(err, tt.wantErr) {
				t.Errorf("normalizeDescription() error = %v, want %v", err, tt.wantErr)
			}
			if normalizedDescription != tt.want {
				t.Errorf("normalizeDescription() got = %q, want %q", normalizedDescription, tt.want)
			}
		})
	}
}

func TestParseStatus(t *testing.T) {
	tests := []struct {
		name    string
		status  string
		want    Status
		wantErr error
	}{
		{name: "valid todo", status: "todo", want: StatusTodo, wantErr: nil},
		{name: "valid in-progress", status: "in-progress", want: StatusInProgress, wantErr: nil},
		{name: "valid done", status: "done", want: StatusDone, wantErr: nil},
		{name: "invalid", status: "completed", want: "", wantErr: ErrInvalidStatus},
		{name: "valid uppercase", status: "TODO", want: StatusTodo, wantErr: nil},
		{name: "valid with space", status: "  done ", want: StatusDone, wantErr: nil},
		{name: "invalid empty string", status: "", want: "", wantErr: ErrInvalidStatus},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			status, err := parseStatus(tt.status)
			if !errors.Is(err, tt.wantErr) {
				t.Errorf("parseStatus() error = %v, want %v", err, tt.wantErr)
			}
			if status != tt.want {
				t.Errorf("parseStatus() got = %q, want %q", status, tt.want)
			}
		})
	}
}
