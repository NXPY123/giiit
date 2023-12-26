package commit

import (
	"os"
	"testing"

	"github.com/NXPY123/giiit/giiit_init"
)

func TestCommit(t *testing.T) {
	// Test with one argument
	project := "../test"
	// Create the project directory
	err := os.Mkdir(project, 0777)
	if err != nil {
		t.Fatalf("Could not create project directory, got %v", err)
	}
	isCreated := giiit_init.Giiit_init(project)
	if !isCreated {
		t.Fatalf("Expected no error, got %v", isCreated)
	}
	// Create a file in the project directory
	file, err := os.Create(project + "/test.txt")
	if err != nil {
		t.Fatalf("Could not create file, got %v", err)
	}
	// Write to the file
	_, err = file.WriteString("Hello World")
	if err != nil {
		t.Fatalf("Could not write to file, got %v", err)
	}
	// Close the file
	err = file.Close()
	if err != nil {
		t.Fatalf("Could not close file, got %v", err)
	}
	// Commit the file
	isCommitted := Commit(project, "test commit")
	if !isCommitted {
		t.Fatalf("Expected no error, got %v", isCommitted)
	}
	// Check if project/.giiit/snapshots/1/test.txt exists
	if _, err := os.Stat(project + "/.giiit/snapshots/0/test.txt"); os.IsNotExist(err) {
		t.Fatalf("Expected project/.giiit/snapshots/0/test.txt to exist, got %v", err)
	}
	// Remove ./project
	err = os.RemoveAll(project)
	if err != nil {
		t.Fatalf("Could not remove project/.giiit, got %v", err)
	}

}
