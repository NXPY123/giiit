package giiit_init

import (
	"fmt"
	"os"
)

func Giiit_init(project string) bool {

	fmt.Println("Initializing project", project)

	// Create project directory
	err := os.Mkdir(project+"/.giiit", 0777)
	if err != nil {
		fmt.Println("Error creating project directory")
		fmt.Println(err)
		return false
	}

	// Inside .giiit, create the following directories:
	// - refs
	// - refs/tags
	// - snapshots

	// Create refs directory
	err = os.Mkdir(project+"/.giiit/refs", 0777)
	if err != nil {
		fmt.Println("Error creating refs directory")
		fmt.Println(err)
		err = os.Remove(project + "/.giiit")
		return false

	}

	// Create refs/heads directory
	err = os.Mkdir(project+"/.giiit/refs/heads", 0777)
	if err != nil {
		fmt.Println("Error creating refs/heads directory")
		fmt.Println(err)
		err = os.Remove(project + "/.giiit")
		return false

	}

	// Create refs/tags directory
	err = os.Mkdir(project+"/.giiit/refs/tags", 0777)
	if err != nil {
		fmt.Println("Error creating refs/tags directory")
		fmt.Println(err)
		err = os.Remove(project + "/.giiit")
		return false
	}

	// Create snapshots directory
	err = os.Mkdir(project+"/.giiit/snapshots", 0777)
	if err != nil {
		fmt.Println("Error creating snapshots directory")
		fmt.Println(err)
		err = os.Remove(project + "/.giiit")
		return false
	}

	fmt.Println("Initialized empty Giiit repository in", project+"/.giiit")
	return true

}
