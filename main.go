package main

import (
	"fmt"
	"os"
	"strings"
)

/*
Given an absolute pathname that may have . or .. as part of it,
return the shortest standardized path.

For example,
given "/usr/bin/../bin/./scripts/../",
return "/usr/bin/".
*/

func main() {
	if len(os.Args) < 2 {
		fmt.Fprintf(os.Stderr, "Need a path on command line\n")
		os.Exit(1)
	}

	components := strings.Split(os.Args[1], "/")

	var normalizedPath []string

	for _, component := range components {
		if component == "." {
			continue
		}
		if component == ".." {
			l := len(normalizedPath)
			if l > 0 {
				if len(normalizedPath[l-1]) > 0 {
					normalizedPath = normalizedPath[:l-1]
				}
			}
			continue
		}

		if len(component) > 0 {
			normalizedPath = append(normalizedPath, component)
		}
	}

	path := strings.Join(normalizedPath, "/")
	if len(path) > 0 {
		fmt.Printf("/%s/\n", path)
	} else {
		fmt.Printf("/%s\n", path)
	}
}
