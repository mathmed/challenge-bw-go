package main

import (
	"log"
	"os"
)


func main() {
	
	port := os.Getenv("PORT")
	
	if port == "" {
		port = "8080"
	}

	log.Printf("Listening on port %s", port)

}