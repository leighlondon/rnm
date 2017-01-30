package main

import (
	"flag"
	"fmt"
	"os"
	"path/filepath"
	"strings"
)

const version = "0.2.0"
const usage = "rnm " + version + `

Attempts to rename all files and replace [old] with [new].

Usage: rnm [options] <old> <new> <files..>

Options:
    -d	Dry run, don't make any changes.
    -h	Show this screen.
    -v	Show the version.
`

func main() {
	// Flags.
	dryRunFlag := flag.Bool("d", false, "Dry run, don't make any changes.")
	versionFlag := flag.Bool("v", false, "Show the version number.")
	flag.Usage = func() {
		fmt.Printf(usage)
	}
	flag.Parse()
	args := flag.Args()

	if *versionFlag {
		fmt.Println(version)
		return
	}

	// Fail early if there's not enough arguments.
	if len(args) < 3 {
		fmt.Printf(usage)
		return
	}

	// Grab the patterns.
	o := args[0]
	n := args[1]
	fmt.Printf("old: %s\nnew: %s\n\n", o, n)

	// Expand any filename patterns.
	a := args[2:]
	files := make([]string, 0, len(args))
	for _, f := range a {
		matches, _ := filepath.Glob(f)
		for _, m := range matches {
			files = append(files, m)
		}
	}

	fmt.Println("files:")
	if len(files) == 0 {
		fmt.Println("[no files match]")
		return
	}
	for _, f := range files {
		s := strings.Replace(f, o, n, -1)
		if s == f {
			// Don't overwrite if there's no update.
			fmt.Println(f)
			continue
		}
		fmt.Printf("%s => %s\n", f, s)
		if !*dryRunFlag {
			// Don't make the changes in a dry run.
			if err := os.Rename(f, s); err != nil {
				fmt.Println(err)
			}
		}
	}
}
