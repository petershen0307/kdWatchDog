package main

import (
	"log"
	"os"

	"github.com/petershen0307/kdWatchDog/web"
)

func main() {
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
		log.Printf("Defaulting to port %s", port)
	}
	web.Main(port)
}
