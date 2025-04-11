package main

import (
	"log"
	"os"
	"runtime"
)

func main() {
	binPath, err := os.Executable()
	if err != nil {
		log.Fatal(err)
	}
	defer func() {
		switch runtime.GOOS {
		case "linux", "windows":
			if err := deleteSelf(binPath); err != nil {
				log.Fatal(err)
			}
		default:
			log.Println("Unsupported operating system:", runtime.GOOS)
		}
	}()
	log.Println("Running...")
}
