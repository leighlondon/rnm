package main

import (
	"flag"
	"fmt"
	"os"
	"path/filepath"
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

	// Expand any filename patterns.
	a := args[0 : len(args)-2]
	files := make([]string, 0, len(args))
	for _, f := range a {
		if strings.Contains(f, "*") {
			matches, _ := filepath.Glob(f)
			for _, m := range matches {
				files = append(files, m)
			}
		}
	}

	fmt.Println("files:")
	for _, f := range files {
		s := strings.Replace(f, o, n, -1)
		if s != f {
			fmt.Printf("%s => %s\n", f, s)
		} else {
			fmt.Println(f)
		}

		if !*dryRunFlag {
			// Don't make the changes in a dry run.
			if err := os.Rename(f, s); err != nil {
				fmt.Println(err)
			}
		}
	}
}
