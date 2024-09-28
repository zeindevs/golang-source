package main

import "fmt"

var (
	buildDate string = "N/A"
	commitHash string = "N/A"
)

func main() {
	fmt.Println("This binary is built at", buildDate)
	fmt.Println("Commit hash", commitHash)
}
