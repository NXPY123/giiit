package commit

import (
	"os"
	"testing"

	"github.com/NXPY123/giiit/giiit_init"
)

func TestCommit(t *testing.T) {
	// Test with one argument
	project := "../test"
	commit_hash := "rDKHukRcXVc0Vb5WQM+bhaIut7E="
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
	// Create a .giiitignore file
	file, err = os.Create(project + "/.giiitignore")
	if err != nil {
		t.Fatalf("Could not create .giiitignore file, got %v", err)
	}
	// Write to the file
	_, err = file.WriteString(project + ".test2.txt")
	if err != nil {
		t.Fatalf("Could not write to .giiitignore file, got %v", err)
	}
	// Close the file
	err = file.Close()
	if err != nil {
		t.Fatalf("Could not close .giiitignore file, got %v", err)
	}
	// Create a file in the project directory
	file, err = os.Create(project + "/test2.txt")
	if err != nil {
		t.Fatalf("Could not create file, got %v", err)
	}
	// Commit the file
	isCommitted := Commit(project, "test commit")
	if !isCommitted {
		t.Fatalf("Expected no error, got %v", isCommitted)
	}
	// Check if project/.giiit/snapshots/0/test.txt exists
	if _, err := os.Stat(project + "/.giiit/snapshots/" + commit_hash + "/test.txt"); os.IsNotExist(err) {
		t.Fatalf("Expected project/.giiit/snapshots/"+commit_hash+"/test.txt to exist, got %v", err)
	}
	// Ensure test2.txt is not committed
	if _, err := os.Stat(project + "/.giiit/snapshots/" + commit_hash + "/test2.txt"); !os.IsNotExist(err) {
		t.Fatalf("Expected project/.giiit/snapshots/"+commit_hash+"/test2.txt to not exist, got %v", err)
	}
	// Remove ./project
	err = os.RemoveAll(project)
	if err != nil {
		t.Fatalf("Could not remove project/.giiit, got %v", err)
	}

}
