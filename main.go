package main

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"
)

func cleanFilename(name string) string {
	// Convert to lowercase
	lower := strings.ToLower(name)
	// Replace spaces with dashes
	cleaned := strings.ReplaceAll(lower, " ", "-")
	return cleaned
}

func main() {
	if len(os.Args) < 2 {
		fmt.Fprintf(os.Stderr, "Usage: %s <directory>\n", os.Args[0])
		os.Exit(1)
	}

	dir := os.Args[1]

	// Check if the path exists and is a directory
	info, err := os.Stat(dir)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error: cannot access '%s': %v\n", dir, err)
		os.Exit(1)
	}
	if !info.IsDir() {
		fmt.Fprintf(os.Stderr, "Error: '%s' is not a directory\n", dir)
		os.Exit(1)
	}

	// Read the directory entries
	entries, err := os.ReadDir(dir)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error reading directory '%s': %v\n", dir, err)
		os.Exit(1)
	}

	for _, entry := range entries {
		// Skip directories and special files (like . and .. which are not returned by ReadDir anyway)
		if entry.IsDir() {
			continue
		}

		oldName := entry.Name()
		newName := cleanFilename(oldName)

		// Only rename if the name actually changes
		if oldName != newName {
			oldPath := filepath.Join(dir, oldName)
			newPath := filepath.Join(dir, newName)

			err = os.Rename(oldPath, newPath)
			if err != nil {
				fmt.Fprintf(os.Stderr, "Error renaming '%s' to '%s': %v\n", oldPath, newPath, err)
			} else {
				fmt.Printf("Renamed: '%s' -> '%s'\n", oldName, newName)
			}
		}
	}
}

