package tag

import (
	"fmt"
	"os"
	"testing"

	"github.com/NXPY123/giiit/commit"
	"github.com/NXPY123/giiit/giiit_init"
)

func TestTag(t *testing.T) {

	project := "../test"
	tag := "v1.0.0"
	commit_no := "0"

	// Create project directory
	err := os.Mkdir(project, 0777)
	if err != nil {
		fmt.Println("Error creating project directory")
		fmt.Println(err)
		return
	}
	// Initialize giiit
	success := giiit_init.Giiit_init(project)
	if !success {
		fmt.Println("Error initializing giiit")
		return
	}

	// Create a file test
	file, err := os.Create(project + "/test")
	if err != nil {
		fmt.Println("Error creating test file")
		fmt.Println(err)
		return
	}
	_, err = file.WriteString("test")
	if err != nil {
		fmt.Println("Error writing to test file")
		fmt.Println(err)
		return
	}
	err = file.Close()
	if err != nil {
		fmt.Println("Error closing test file")
		fmt.Println(err)
		return
	}

	// Commit
	success = commit.Commit(project, "test")
	if !success {
		fmt.Println("Error committing")
		return
	}

	// Tag
	success = Tag(project, tag, commit_no)
	if !success {
		fmt.Println("Error tagging")
		return
	}

	// Check if project/.giiit/refs/tags/<tag> exists
	if _, err := os.Stat(project + "/.giiit/refs/tags/" + tag); os.IsNotExist(err) {
		fmt.Println("Error tagging")
		return
	}

	// Check if project/.giiit/refs/tags/<tag> contains <commit>
	tagFile, err := os.Open(project + "/.giiit/refs/tags/" + tag)
	if err != nil {
		fmt.Println("Error opening tag file")
		fmt.Println(err)
		return
	}
	var commit_no_read string
	_, err = fmt.Fscan(tagFile, &commit_no_read)
	if err != nil {
		fmt.Println("Error reading tag file")
		fmt.Println(err)
		return
	}
	err = tagFile.Close()
	if err != nil {
		fmt.Println("Error closing tag file")
		fmt.Println(err)
		return
	}
	if commit_no_read != commit_no {
		fmt.Println("Error tagging")
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
