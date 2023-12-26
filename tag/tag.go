package tag

import (
	"fmt"
	"os"
)

func Tag(project string, tag string, commit string) bool {

	// Check if project/.giiit exists
	if _, err := os.Stat(project + "/.giiit"); os.IsNotExist(err) {
		fmt.Println("fatal: not a giiit repository (or any of the parent directories): .giiit")
		return false
	}

	// Check if project/.giiit/refs exists
	if _, err := os.Stat(project + "/.giiit/refs"); os.IsNotExist(err) {
		fmt.Println("fatal: not a giiit repository (or any of the parent directories): .giiit")
		return false
	}

	// Check if project/.giiit/refs/heads exists
	if _, err := os.Stat(project + "/.giiit/refs/heads"); os.IsNotExist(err) {
		fmt.Println("fatal: no ref to commit on: refs/heads")
		return false
	}

	// Check if project/.giiit/refs/tags exists
	if _, err := os.Stat(project + "/.giiit/refs/tags"); os.IsNotExist(err) {
		fmt.Println("fatal: no ref to commit on: refs/tags")
		return false
	}

	// Check if project/.giiit/snapshots exists
	if _, err := os.Stat(project + "/.giiit/snapshots"); os.IsNotExist(err) {
		fmt.Println("fatal: no ref to commit on: snapshots")
		return false
	}

	// Create a file in project/.giiit/refs/tags/<tag> with the contents <commit>
	tagFile, err := os.Create(project + "/.giiit/refs/tags/" + tag)
	if err != nil {
		fmt.Println("Error creating tag file")
		fmt.Println(err)
		return false
	}
	_, err = tagFile.WriteString(commit)
	if err != nil {
		fmt.Println("Error writing to tag file")
		fmt.Println(err)
		return false
	}
	err = tagFile.Close()
	if err != nil {
		fmt.Println("Error closing tag file")
		fmt.Println(err)
		return false
	}

	return true
}
