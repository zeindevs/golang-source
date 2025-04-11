//go:build windows

package main

import (
	"log"
	"os/exec"
	"runtime"
	"syscall"
)

func deleteSelf(binPath string) error {
	cmd := exec.Command("cmd", "/C", "timeout 2 && del", binPath)
	cmd.SysProcAttr = &syscall.SysProcAttr{HideWindow: true}
	if err := cmd.Start(); err != nil {
		return err
	}
	log.Println("Deleting... OS:", runtime.GOOS)
	return nil
}
