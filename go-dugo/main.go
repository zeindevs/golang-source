package main

import (
	"fmt"
	"os"
	"path/filepath"
	"sort"
	"strconv"
	"strings"
	"time"
)

func formatBytes(bytes int64) string {
	sizes := []string{"B", "KB", "MB", "GB", "TB"}
	var sizeIndex int
	floatBytes := float64(bytes)

	for floatBytes >= 1024 && sizeIndex < len(sizes)-1 {
		sizeIndex++
		floatBytes /= 1024
	}

	return fmt.Sprintf("%.2f %s", floatBytes, sizes[sizeIndex])
}

type DirSize struct {
	Path  string
	Size  int64
	Depth int
}

// Function to calculate the depth of a path relative to a root directory
func getDepth(rootDir, path string) int {
	rel, err := filepath.Rel(rootDir, path)
	if err != nil {
		return -1 // Error occurred
	}
	return len(filepath.SplitList(rel))
}

func main() {
	start := time.Now()

	if len(os.Args) < 4 {
		fmt.Printf("Usage: dugo <directory> <max-depth> -d/-f")
		return
	}

	// Define a map to store directory sizes
	var dirSizes []DirSize

	// Get the starting directory
	filePath := os.Args[1] // Change this to the desired directory

	// Get operation mode
	optmode := os.Args[3]

	// Specify the maximum depth
	maxDepth, err := strconv.Atoi(os.Args[2]) // Change this to the desired maximum depth
	if err != nil {
		fmt.Println("Depth must a number not string")
		return
	}

	// Get the absolute path of the file
	rootDir, err := filepath.Abs(filePath)
	if err != nil {
		fmt.Println("Error getting absolute path:", err)
		return
	}

	if optmode == "-d" {
		// Walk through the directory tree
		err = filepath.Walk(rootDir, func(path string, info os.FileInfo, err error) error {
			if err != nil {
				fmt.Printf("Error accessing path %s: %v\n", path, err)
				return nil
			}
			depth := getDepth(rootDir, path)
			if depth <= maxDepth {
				if info.IsDir() {
					dirSizes = append(dirSizes, DirSize{Path: path, Size: 0, Depth: depth})
				} else {
					// Accumulate the file size for its parent directory
					parentDir := filepath.Dir(path)
					for i := range dirSizes {
						if dirSizes[i].Path == parentDir {
							dirSizes[i].Size += info.Size()
							break
						}
					}
				}
			}
			return nil
		})

	} else if optmode == "-f" {
		// Walk through the directory tree
		err = filepath.Walk(rootDir, func(path string, info os.FileInfo, err error) error {
			if err != nil {
				fmt.Printf("Error accessing path %s: %v\n", path, err)
				return nil
			}
			depth := getDepth(rootDir, path)
			if depth <= maxDepth {
				if !info.IsDir() {
					dirSizes = append(dirSizes, DirSize{Path: path, Size: info.Size(), Depth: depth})
				}
			}
			return nil
		})
	} else {
		fmt.Println("Operation must a following -d/-f")
		return
	}

	if err != nil {
		fmt.Printf("Error walking through directory: %v\n", err)
		return
	}
	// Sort directory sizes by size
	sort.Slice(dirSizes, func(i, j int) bool {
		return dirSizes[i].Size > dirSizes[j].Size
	})

	// Print the directory sizes
	for _, dirSize := range dirSizes {
		fmt.Printf("%s | %s\n", formatBytes(dirSize.Size), strings.ReplaceAll(dirSize.Path, rootDir, "."))
	}

	fmt.Println("took", time.Since(start).String())
}
