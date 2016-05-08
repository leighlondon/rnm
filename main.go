package main

import (
	"flag"
	"fmt"
	"os"
	"strings"
)

/*
 * Usage: ./rnm [files] [old] [new]
 *
 * Attempts to rename all files and replace [old] with [new].
 */

func main() {
	// Flags.
	dryRunFlag := flag.Bool("d", false, "Dry run, don't make any changes.")
	versionFlag := flag.Bool("v", false, "Show the version number.")
	flag.Parse()
	args := flag.Args()

	version := "rnm 0.0.0"
	if *versionFlag {
		fmt.Println(version)
		return
	}

	// Fail early if there's not enough arguments.
	if len(args) < 3 {
		return
	}

	// Grab the patterns.
	o := args[len(args)-2]
	n := args[len(args)-1]
	fmt.Printf("old: %s\nnew: %s\n\n", o, n)

	fmt.Println("files:")
	for i := 0; i < len(args)-2; i++ {
		s := strings.Replace(args[i], o, n, -1)
		if s != args[i] {
			fmt.Printf("%s => %s\n", args[i], s)
		} else {
			fmt.Println(args[i])
		}

		if !*dryRunFlag {
			// Don't make the changes in a dry run.
			if err := os.Rename(args[i], s); err != nil {
				fmt.Println(err)
			}
		}
	}
}
