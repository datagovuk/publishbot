package main

import (
	"fmt"
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

func findAdapter(name string) AdapterConfig {
	for _, adapter := range config.Adapters {
		if adapter.Name == name {
			return adapter
		}
	}
	return AdapterConfig{}
}

func directoryAdapter(config AdapterConfig) {
	log.Println("Watching for changes in ", config.Arguments["folder"])
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
				if event.Op == fsnotify.Create {
					fmt.Println("Created new file")
					// Notify user
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
