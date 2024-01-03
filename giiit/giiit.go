package main

import (
	"fmt"
	"os"

	"github.com/NXPY123/giiit/commit"
	"github.com/NXPY123/giiit/giiit_init"
	"github.com/NXPY123/giiit/log"
	"github.com/NXPY123/giiit/tag"
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
		os.Exit(bool2int(!success))
	case "commit":
		success := commit.Commit(os.Args[1], os.Args[2])
		os.Exit(bool2int(!success))
	case "log":
		success := log.LogView(os.Args[1])
		os.Exit(bool2int(!success))
	case "tag":
		project := os.Args[1]
		tag_name := os.Args[2]
		commit := os.Args[3]
		success := tag.Tag(project, tag_name, commit)
		os.Exit(bool2int(!success))
	default:
		fmt.Println("giiit: '" + os.Args[0] + "' is not a giiit command. See 'giiit --help'")

	}
}

func bool2int(b bool) int {
	if b {
		return 1
	}
	return 0
}
