package main

import (
	"fmt"
	"os"

	"github.com/NXPY123/giiit/commit"
	"github.com/NXPY123/giiit/giiit_init"
)

func main() {

	os.Args = os.Args[1:]

	if len(os.Args) == 0 {
		fmt.Println("giiit version 0.1.0")
		fmt.Println("usage: giiit <command> [<args>]")
		fmt.Println("For more information about a command, run giiit help <command>")
		os.Exit(0)
	}

	switch os.Args[0] {
	case "init":
		success := giiit_init.Giiit_init(os.Args[1])
		if success {
			os.Exit(0)
		} else {
			os.Exit(1)
		}
	case "commit":
		success := commit.Commit(os.Args[1], os.Args[2])
		if success {
			os.Exit(0)
		} else {
			os.Exit(1)
		}
	default:
		fmt.Println("giiit: '" + os.Args[0] + "' is not a giiit command. See 'giiit --help'")

	}
}
