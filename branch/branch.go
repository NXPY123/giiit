package branch

import (
	"fmt"
	"os"
	"regexp"
)

func CreateBranch(project string, branch string) bool {

	// Create a new branch
	fmt.Println("Creating branch", branch)

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

	//Check if branch already exists in project/.giiit/branches/branches.txt
	branchesFile, err := os.ReadFile(project + "/.giiit/branches/branches.txt")
	if err != nil {
		fmt.Println("Error reading branches file")
		fmt.Println(err)
		return false
	}
	branches := regexp.MustCompile("\r\n|\n\r|\n|\r").Split(string(branchesFile), -1)
	for _, b := range branches {
		if b == branch {
			fmt.Println("fatal: branch already exists")
			return false
		}
	}

	// Add branch to project/.giiit/branches/branches.txt
	// append to branchesFile
	branchesFile = append(branchesFile, []byte("\n"+branch)...)
	err = os.WriteFile(project+"/.giiit/branches/branches.txt", branchesFile, 0777)
	if err != nil {
		fmt.Println("Error writing to branches file")
		fmt.Println(err)
		return false
	}

	//find current branch from project/.giiit/branches/current_branch.txt
	currentBranchFile, err := os.ReadFile(project + "/.giiit/branches/current_branch.txt")
	if err != nil {
		fmt.Println("Error reading current branch file")
		fmt.Println(err)
		return false
	}

	// Copy contents of "/.giiit/refs/heads/ + currentBranchFile" to "/.giiit/refs/heads/ + branch"
	currentBranch, err := os.ReadFile(project + "/.giiit/refs/heads/" + string(currentBranchFile))
	if err != nil {
		fmt.Println("Error reading current branch file")
		fmt.Println(err)
		return false
	}

	// Put the contents of current branch head file into the new branch head file
	err = os.WriteFile(project+"/.giiit/refs/heads/"+branch, currentBranch, 0777)
	if err != nil {
		fmt.Println("Error writing to branch head file")
		fmt.Println(err)
		return false
	}

	return true

}

func SwitchBranch(project string, branch string) bool {

	// Switch to a branch
	fmt.Println("Switching to branch", branch)

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

	//Check if branch exists in project/.giiit/branches/branches.txt
	branchesFile, err := os.ReadFile(project + "/.giiit/branches/branches.txt")
	if err != nil {
		fmt.Println("Error reading branches file")
		fmt.Println(err)
		return false
	}
	branches := regexp.MustCompile("\r\n|\n\r|\n|\r").Split(string(branchesFile), -1)
	found := false
	for _, b := range branches {
		if b == branch {
			found = true
			break
		}
	}
	if !found {
		fmt.Println("fatal: branch does not exist")
		return false
	}

	// Write branch to project/.giiit/branches/current_branch.txt
	// delete current_branch.txt
	err = os.Remove(project + "/.giiit/branches/current_branch.txt")
	if err != nil {
		fmt.Println("Error deleting current branch file")
		fmt.Println(err)
		return false
	}
	// create current_branch.txt
	err = os.WriteFile(project+"/.giiit/branches/current_branch.txt", []byte(branch), 0777)
	if err != nil {
		fmt.Println("Error writing to current branch file")
		fmt.Println(err)
		return false
	}

	return true

}

func Branch(project string, flag string, branch string) bool {

	// If flag is -c or --create, create a new branch
	if flag == "-c" || flag == "--create" {
		return CreateBranch(project, branch)
	}

	if flag == "-s" || flag == "--switch" {
		return SwitchBranch(project, branch)
	}

	fmt.Println("fatal: unknown flag")
	return false

}
