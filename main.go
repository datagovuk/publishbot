package main

import (
	"log"
	"net/http"
	"os"
)

func main() {

	setup_routes()
	open_db()
	defer close_db()

	port := os.Getenv("PORT")
	if len(port) == 0 {
		port = "3000"
	}

	loadConfigFile("./test.yml")

	log.Printf("Found %d adapters\n", len(config.Adapters))

	log.Println("Listening... on 127.0.0.1:" + port)
	http.ListenAndServe("127.0.0.1:"+port, nil)
}
