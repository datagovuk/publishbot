package main

import (
	"log"

	"github.com/fsnotify/fsnotify"
)

func RunAdapter(config AdapterConfig) {
	log.Println("Running", config.Title)
	if config.Type == "directory" {
		directoryAdapter(config)
	} else {
		log.Printf("Don't know how to run a %s adapter\n", config.Type)
	}
}

func directoryAdapter(config AdapterConfig) {
	watcher, err := fsnotify.NewWatcher()
	if err != nil {
		log.Fatal(err)
	}
	defer watcher.Close()

	done := make(chan bool)
	go func() {
		for {
			select {
			case event := <-watcher.Events:
				log.Println("event:", event)
				if event.Op&fsnotify.Write == fsnotify.Write {
					log.Println("modified file:", event.Name)
				}
			case err := <-watcher.Errors:
				log.Println("error:", err)
			}
		}
	}()

	err = watcher.Add(config.Arguments["folder"])
	if err != nil {
		log.Fatal(err)
	}
	<-done
}
