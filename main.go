package main

import (
	"log"
	"net/http"
	"os"
)

func main() {

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

	log.Println("\nListening... on 127.0.0.1:" + port)
	http.ListenAndServe("127.0.0.1:"+port, nil)
}
