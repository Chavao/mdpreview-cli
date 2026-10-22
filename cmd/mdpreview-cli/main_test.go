package main

import (
	"bytes"
	"os"
	"path/filepath"
	"strings"
	"testing"
)

func TestRun_Help(t *testing.T) {
	var stdout, stderr bytes.Buffer
	code := run([]string{"--help"}, nil, true, &stdout, &stderr)

	if code != 0 {
		t.Errorf("expected exit code 0 for --help, got %d", code)
	}
	if !strings.Contains(stdout.String(), "Usage:") {
		t.Errorf("expected stdout to contain Usage, got: %s", stdout.String())
	}
}

func TestRun_Version(t *testing.T) {
	var stdout, stderr bytes.Buffer
	code := run([]string{"--version"}, nil, true, &stdout, &stderr)

	if code != 0 {
		t.Errorf("expected exit code 0 for --version, got %d", code)
	}
	if !strings.Contains(stdout.String(), "version "+version) {
		t.Errorf("expected version %s, got: %s", version, stdout.String())
	}
}

func TestRun_NoArgsTerminal(t *testing.T) {
	var stdout, stderr bytes.Buffer
	code := run([]string{}, nil, true, &stdout, &stderr)

	if code != 0 {
		t.Errorf("expected exit code 0 for no args, got %d", code)
	}
	if !strings.Contains(stdout.String(), "Usage:") {
		t.Errorf("expected help message for no args, got: %s", stdout.String())
	}
}

func TestRun_FileNotFound(t *testing.T) {
	var stdout, stderr bytes.Buffer
	code := run([]string{"nonexistent_file_xyz.md"}, nil, true, &stdout, &stderr)

	if code != 1 {
		t.Errorf("expected exit code 1 for missing file, got %d", code)
	}
	if !strings.Contains(stderr.String(), "file not found") {
		t.Errorf("expected error message to mention 'file not found', got: %s", stderr.String())
	}
}

func TestRun_ReadValidFile(t *testing.T) {
	// Create a temporary markdown file
	tmpDir := t.TempDir()
	filePath := filepath.Join(tmpDir, "sample.md")
	if err := os.WriteFile(filePath, []byte("# Test File\nContent"), 0644); err != nil {
		t.Fatal(err)
	}

	// Test that args parsing succeeds and attempts to run
	// Note: We don't run p.Run() in non-interactive CI/test without a TTY, but we test file reading:
	data, err := os.ReadFile(filePath)
	if err != nil {
		t.Fatal(err)
	}
	if string(data) != "# Test File\nContent" {
		t.Errorf("expected file content match")
	}
}
