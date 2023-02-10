package host

import (
	"bufio"
	"io/ioutil"
	"os"
	"testing"
)

func TestUpdateLine(t *testing.T) {
	file, err := ioutil.TempFile("", "hosts")
	if err != nil {
		t.Fatalf("Failed to create temp file: %s", err)
	}
	defer os.Remove(file.Name())

	// Test updating an existing line
	oldLine := "127.0.0.1 localhost"
	newLine := "127.0.0.1 example.com"
	if _, err := file.WriteString(oldLine + "\n"); err != nil {
		t.Fatalf("Failed to write to file: %s", err)
	}
	if err := updateLine(file, oldLine, newLine); err != nil {
		t.Fatalf("Failed to update line: %s", err)
	}
	file.Seek(0, 0)
	scanner := bufio.NewScanner(file)
	scanner.Scan()
	if got := scanner.Text(); got != newLine {
		t.Errorf("Expected %q, got %q", newLine, got)
	}

	// Test updating a non-existing line
	oldLine = "127.0.0.2 example.org"
	newLine = "127.0.0.2 example.net"
	file.Truncate(0)
	file.Seek(0, 0)
	if err := updateLine(file, oldLine, newLine); err != nil {
		t.Fatalf("Failed to update line: %s", err)
	}
	file.Seek(0, 0)
	scanner = bufio.NewScanner(file)
	if scanner.Scan() {
		t.Error("Expected file to be empty")
	}
}
