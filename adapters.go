package main

import "fmt"

func List() []Adapter {
	return adapters
}

func RunAdapter(config AdapterConfig) {
	fmt.Println("Running", config.Title)

}
