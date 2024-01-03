package log

import (
	"fmt"
	"os"
	"testing"

	"github.com/NXPY123/giiit/commit"
	"github.com/NXPY123/giiit/giiit_init"
)

func TestLog(t *testing.T) {
	project := "../test"
	// Create project directory
	err := os.Mkdir(project, 0777)
	if err != nil {
		fmt.Println("Error creating project directory")
		fmt.Println(err)
		err = os.RemoveAll(project)
		return
	}
	// Initialize giiit
	success := giiit_init.Giiit_init(project)
	if !success {
		fmt.Println("Error initializing giiit")
		err = os.RemoveAll(project)
		return
	}

	// Create a file test
	file, err := os.Create(project + "/test")
	if err != nil {
		fmt.Println("Error creating test file")
		fmt.Println(err)
		err = os.RemoveAll(project)
		return
	}
	_, err = file.WriteString("test")
	if err != nil {
		fmt.Println("Error writing to test file")
		fmt.Println(err)
		err = os.RemoveAll(project)
		return
	}
	err = file.Close()
	if err != nil {
		fmt.Println("Error closing test file")
		fmt.Println(err)
		err = os.RemoveAll(project)
		return
	}

	// Commit
	success = commit.Commit(project, "test")
	if !success {
		fmt.Println("Error committing")
		// Remove project directory
		err = os.RemoveAll(project)
		return
	}

	// Log
	success = LogView(project)
	if !success {
		fmt.Println("Error logging")
		return
	}

	// Remove project directory
	err = os.RemoveAll(project)
	if err != nil {
		fmt.Println("Error removing project directory")
		fmt.Println(err)
		return
	}

}
