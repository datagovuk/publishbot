package main

import "log"

func RunAdapter(config AdapterConfig) {
	log.Println("Running", config.Title)
	if config.Type == "directory" {
		directoryAdapter(config)
	} else {
		log.Printf("Don't know how to run a %s adapter\n", config.Type)
	}
}

func directoryAdapter(config AdapterConfig) {

}
