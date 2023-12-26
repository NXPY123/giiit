package giiit_init

import (
	"os"
	"testing"
)

// Test the Giiit_init function

func TestGiiit(t *testing.T) {
	// Test with one argument
	project := "/Users/neeraj_py/Desktop/VCS/giit/test"
	isCreated := Giiit_init(project)
	if !isCreated {
		t.Fatalf("Expected no error, got %v", isCreated)
	}
	// Check if project/.giiit exists
	if _, err := os.Stat(project + "/.giiit"); os.IsNotExist(err) {
		t.Fatalf("Expected project/.giiit to exist, got %v", err)
	}
	// Check if project/.giiit/refs exists
	if _, err := os.Stat(project + "/.giiit/refs"); os.IsNotExist(err) {
		t.Fatalf("Expected project/.giiit/refs to exist, got %v", err)
	}
	// Check if project/.giiit/refs/tags exists
	if _, err := os.Stat(project + "/.giiit/refs/tags"); os.IsNotExist(err) {
		t.Fatalf("Expected project/.giiit/refs/tags to exist, got %v", err)
	}
	// Check if project/.giiit/snapshots exists
	if _, err := os.Stat(project + "/.giiit/snapshots"); os.IsNotExist(err) {
		t.Fatalf("Expected project/.giiit/snapshots to exist, got %v", err)
	}
	// Remove ./project/.giiit
	err := os.RemoveAll(project + "/.giiit")
	if err != nil {
		t.Fatalf("Could not remove project/.giiit, got %v", err)
	}

}
