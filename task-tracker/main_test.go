package main

import "testing"

func TestParseID(t *testing.T) {
	tests := []struct {
		name    string
		input   string
		want    int
		wantErr bool
	}{
		{name: "valid 1 with space", input: " 1  ", want: 1, wantErr: false},
		{name: "invalid 0", input: "0", want: 0, wantErr: true},
		{name: "invalid -1", input: "-1", want: 0, wantErr: true},
		{name: "invalid not a number", input: "abc", want: 0, wantErr: true},
		{name: "invalid empty string", input: "", want: 0, wantErr: true},
		{name: "invalid positive decimal", input: "0.1", want: 0, wantErr: true},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			id, err := parseID(tt.input)

			if (err != nil) != tt.wantErr {
				t.Errorf("parseID(%q) error = %v, wantErr %v", tt.input, err, tt.wantErr)
			}

			if id != tt.want {
				t.Errorf("parseID(%q) = %d, want %d", tt.input, id, tt.want)
			}
		})
	}
}
