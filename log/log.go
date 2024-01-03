package log

import (
	"fmt"
	"os"
)

func LogView(project string) bool {
	// For all the snapshots, print the contents of the commit message file
	// Check if project/.giiit/snapshots exists
	if _, err := os.Stat(project + ".giiit/snapshots"); os.IsNotExist(err) {
		fmt.Println("fatal: no ref to commit on: snapshots")
		return false
	}
	// Get the next snapshot number
	nextSnapshotNumber := 0
	// While the next snapshot number exists, increment it
	for {
		if _, err := os.Stat(project + ".giiit/snapshots/" + fmt.Sprint(nextSnapshotNumber)); os.IsNotExist(err) {
			break
		}

		// Print the commit message
		commitMessageFile, err := os.ReadFile(".giiit/snapshots/" + fmt.Sprint(nextSnapshotNumber) + "/.commit")
		if err != nil {
			fmt.Println("Error reading commit message file")
			fmt.Println(err)
			return false
		}
		fmt.Println(string(commitMessageFile))
		nextSnapshotNumber++
	}
	return true
}
