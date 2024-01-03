package commit

import (
	"fmt"
	"io"
	"os"
	"regexp"
	"time"
)

func Commit(project string, message string) bool {

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

	// Create a sub directory in snapshots with the directory name as the next snapshot number
	// Get the next snapshot number
	nextSnapshotNumber := 0
	// While the next snapshot number exists, increment it
	for {
		if _, err := os.Stat(project + "/.giiit/snapshots/" + fmt.Sprint(nextSnapshotNumber)); os.IsNotExist(err) {
			break
		}
		nextSnapshotNumber++
	}
	// Create the next snapshot number directory
	err := os.Mkdir(project+"/.giiit/snapshots/"+fmt.Sprint(nextSnapshotNumber), 0777)
	if err != nil {
		fmt.Println("Error creating next snapshot number directory")
		fmt.Println(err)
		return false
	}

	// Copy all the files to the next snapshot number directory except .giiit
	// Get all the files in the project directory
	files, err := os.ReadDir(project)
	if err != nil {
		fmt.Println("Error reading project directory")
		fmt.Println(err)
		// Remove the next snapshot number directory
		err = os.RemoveAll(project + "/.giiit/snapshots/" + fmt.Sprint(nextSnapshotNumber))
		return false
	}

	fileregex := regexp.MustCompile("^$")
	// Check if there is a .giiitignore file
	if _, err := os.Stat(project + "/.giiitignore"); err == nil {
		// Read the .giiitignore file
		giiitignoreFile, err := os.ReadFile(project + "/.giiitignore")
		if err != nil {
			fmt.Println("Error reading .giiitignore file")
			fmt.Println(err)
			return false
		}
		// Split the .giiitignore file into lines
		giiitignoreLines := regexp.MustCompile("\r\n|\n\r|\n|\r").Split(string(giiitignoreFile), -1)
		// For each line in the .giiitignore file
		regexpstring := ""
		for _, giiitignoreLine := range giiitignoreLines {
			// If the line is not empty
			if giiitignoreLine != "" {
				// Add the line to the regexpstring
				regexpstring += giiitignoreLine + "|"
			}
		}
		// Remove the last | from the regexpstring
		regexpstring = regexpstring[:len(regexpstring)-1]
		// Compile the regexpstring
		fileregex, err = regexp.Compile(regexpstring)
		if err != nil {
			fmt.Println("Error compiling regexpstring")
			fmt.Println(err)
			return false
		}
	}

	// Copy all the files to the next snapshot number directory except .giiit
	for _, file := range files {
		if file.Name() != ".giiit" {
			// Check if the file is in the .giiitignore file
			filePath := project + "/" + file.Name()
			if fileregex.MatchString(filePath) {
				// If the file is in the .giiitignore file, do not copy it
				continue
			}
			_, err := Copy(project+"/"+file.Name(), project+"/.giiit/snapshots/"+fmt.Sprint(nextSnapshotNumber)+"/"+file.Name())
			if err != nil {
				fmt.Println("Error copying file")
				fmt.Println(err)
				// Remove the next snapshot number directory
				err = os.RemoveAll(project + "/.giiit/snapshots/" + fmt.Sprint(nextSnapshotNumber))
				return false
			}

		}
	}

	// Create a file in the next snapshot number directory with the commit message
	commitMessageFile, err := os.Create(project + "/.giiit/snapshots/" + fmt.Sprint(nextSnapshotNumber) + "/.commit")
	if err != nil {
		fmt.Println("Error creating commit message file")
		// Remove the next snapshot number directory
		err = os.RemoveAll(project + "/.giiit/snapshots/" + fmt.Sprint(nextSnapshotNumber))
		return false
	}

	// Write the commit message to the commit message file
	// Get the current time
	currentTime := time.Now()
	// Write the user name, commit message and current time to the commit message file
	commitMessageFile.WriteString("Author: " + os.Getenv("USER") + "\n")
	commitMessageFile.WriteString("Date: " + currentTime.Format("Mon Jan 2 15:04:05 2006 -0700") + "\n")
	commitMessageFile.WriteString("\n")
	commitMessageFile.WriteString(message)
	commitMessageFile.Close()

	return true

}

func Copy(src string, dst string) (int64, error) {
	src_file, err := os.Open(src)
	if err != nil {
		return 0, err
	}
	defer src_file.Close()

	src_file_stat, err := src_file.Stat()
	if err != nil {
		return 0, err
	}

	if !src_file_stat.Mode().IsRegular() {
		return 0, fmt.Errorf("%s is not a regular file", src)
	}

	dst_file, err := os.Create(dst)
	if err != nil {
		return 0, err
	}
	defer dst_file.Close()
	return io.Copy(dst_file, src_file)
}
