package cmd

import (
	"bytes"
	"testing"
)

func TestNewCmdVersion(t *testing.T) {
	var buf bytes.Buffer
	cmd := newCmdVersion(&buf)
	if err := cmd.Execute(); err != nil {
		t.Errorf("Cannot execute version command: %v", err)
	}
}

func TestRunVersion(t *testing.T) {
	tests := []struct {
		name    string
		wantOut string
	}{
		{name: "valid", wantOut: "v0.0.0"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			out := &bytes.Buffer{}
			RunVersion(out)
			if gotOut := out.String(); gotOut != tt.wantOut {
				t.Errorf("RunVersion() = %v, want %v", gotOut, tt.wantOut)
			}
		})
	}
}
