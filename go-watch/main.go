package main

import (
	"flag"
	"log"
	"os"
	"os/exec"
	"path"
	"strings"

	"github.com/fsnotify/fsnotify"
)

func main() {
	command := flag.String("command", "make tailwind", "Command name")
	folder := flag.String("folder", "views", "Folder for watching")

	flag.Parse()

	commands := strings.Split(*command, " ")
	folders := []string{*folder}

	entry, err := os.ReadDir(*folder)
	if err != nil {
		log.Fatal(err)
	}

	for _, f := range entry {
		if f.IsDir() {
			folders = append(folders, path.Join(*folder, f.Name()))
		}
	}

	watcher, err := fsnotify.NewWatcher()
	if err != nil {
		log.Fatal(err)
	}
	defer watcher.Close()

	go func() {
		for {
			select {
			case event, ok := <-watcher.Events:
				if !ok {
					return
				}
				if event.Has(fsnotify.Write) {
					log.Println("modified file:", event.Name)
					cmd := exec.Command(commands[0], commands[1:]...)
					stdout, err := cmd.CombinedOutput()
					if err != nil {
						log.Println(err.Error())
					}
					log.Println(string(stdout))
				}
			case err, ok := <-watcher.Errors:
				if !ok {
					return
				}
				log.Println("error:", err)
			}
		}
	}()

	for _, folder := range folders {
		log.Println("watching", folder)
		err = watcher.Add(folder)
		if err != nil {
			log.Fatal(err)
		}
	}

	log.Println("gowatch running")

	<-make(chan struct{})
}
