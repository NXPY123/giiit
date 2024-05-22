package giiit_init

import (
	"os"
	"testing"
)

// Test the Giiit_init function

func TestGiiit(t *testing.T) {
	// Test with one argument
	// Set the project directory as test in the parent directory
	project := "../test"
	// Create the project directory
	err := os.Mkdir(project, 0777)
	if err != nil {
		t.Fatalf("Could not create project directory, got %v", err)
	}
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

	//Check if project/.giiit/refs/branches.txt exists
	if _, err := os.Stat(project + "/.giiit/branches.txt"); os.IsNotExist(err) {
		t.Fatalf("Expected project/.giiit/branches.txt to exist, got %v", err)
	}

	// Remove the project directory
	err = os.RemoveAll(project)
	if err != nil {
		t.Fatalf("Could not remove project directory, got %v", err)
	}

}
