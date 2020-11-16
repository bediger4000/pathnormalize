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

	// Path components: non-zero-length strings between '/' characters
	// The non-zero-length part gets taken care of when examining
	// path components inside the for-loop below.
	components := strings.Split(os.Args[1], "/")

	// Path components after normalizing
	var normalizedPath []string

	for _, component := range components {
		if component == "" {
			// Ignore zero-length components
			continue
		}
		if component == "." {
			// Path stays the same
			continue
		}
		if component == ".." {
			// pop a component off the normalizedPath stack,
			// but don't go "before" the root of the path.
			l := len(normalizedPath)
			if l > 0 {
				normalizedPath = normalizedPath[:l-1]
			}
			continue
		}
		normalizedPath = append(normalizedPath, component)
	}

	// The above algorithm doesn't do leading and trailing
	// slashes. Trailing slashes don't appear on a zero-length
	// normalized path.
	path := strings.Join(normalizedPath, "/")
	if len(path) > 0 {
		fmt.Printf("/%s/\n", path)
	} else {
		fmt.Printf("/%s\n", path)
	}
}
