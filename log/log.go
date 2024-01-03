package log

import (
	"fmt"
	"os"
)

func LogView(project string) bool {
	// For all the snapshots, print the contents of the commit message file
	// Check if project/.giiit/snapshots exists
	if _, err := os.Stat(project + "/.giiit/snapshots"); os.IsNotExist(err) {
		fmt.Println("fatal: no ref to commit on: snapshots")
		return false
	}
	// Get the snapshots in the snapshots directory
	snapshots, err := os.ReadDir(project + "/.giiit/snapshots")
	if err != nil {
		fmt.Println("Error reading snapshots directory")
		fmt.Println(err)
		return false
	}
	// For each snapshot
	for _, snapshot := range snapshots {
		// Print the commit message
		commitMessageFile, err := os.ReadFile(project + "/.giiit/snapshots/" + snapshot.Name() + "/.commit")
		if err != nil {
			fmt.Println("Error reading commit message file")
			fmt.Println(err)
			return false
		}
		fmt.Println(string(commitMessageFile))
	}
	return true
}
