package main

import (
	"fmt"
	"log"
	"os"
	"os/signal"
	"syscall"
)

func main() {
	sigs := make(chan os.Signal, 1)
	done := make(chan bool, 1)
	signal.Notify(sigs, syscall.SIGINT, syscall.SIGTERM)

	setup_routes()

	loadConfigFile("./test.yml")

	port := os.Getenv("PORT")
	if len(port) == 0 {
		port = config.Port
	}

	log.Printf("Found %d adapters\n", len(config.Adapters))
	for _, adapter := range config.Adapters {
		log.Println("  Name: ", adapter.Name)
		log.Println("  Type: ", adapter.Type)
		log.Println("  Arguments")
		for k, v := range adapter.Arguments {
			log.Println("      ", k, v)
		}

		go RunAdapter(adapter)
	}

	go func() {
		sig := <-sigs
		fmt.Println()
		fmt.Println(sig)
		done <- true
	}()

	<-done

	// log.Println("\nListening... on 127.0.0.1:" + port)
	// http.ListenAndServe("127.0.0.1:"+port, nil)
}
