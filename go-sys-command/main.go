package main

import (
	"bytes"
	"fmt"
	"io"
	"log"
	"os"
	"os/exec"
)

func withBuffer() {
	var buffer bytes.Buffer
	cmd := exec.Command("ping", "google.com")
	cmd.Stdout = &buffer
	_ = cmd.Run()

	fmt.Printf("%s\n", buffer.String())
}

func main() {
	f, err := os.OpenFile("log.log", os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)
	if err != nil {
		log.Fatalf("Error opening file: %v", err)
	}
	defer f.Close()

	mwriter := io.MultiWriter(f, os.Stdout)
	cmd := exec.Command("ping", "google.com")
	cmd.Stderr = mwriter
	cmd.Stdout = mwriter
	err = cmd.Run()
	if err != nil {
		panic(err)
	}
}
