//go:build linux

package main

import (
	"fmt"
	"log"
	"os/exec"
	"runtime"
)

func deleteSelf(binPath string) error {
	cmd := exec.Command("sh", "-c", fmt.Sprintf("rm -f %s", binPath))
	if err := cmd.Start(); err != nil {
		return err
	}
	log.Println("Deleting... OS:", runtime.GOOS)
	return nil
}
