package web

import (
	"fmt"
	"log"
	"net/http"
	"time"
)

func helloHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello world %v", time.Now().Format(time.RFC1123Z))
}

// Main is the http listener main function
func Main(port string) {
	http.HandleFunc("/", helloHandler)
	log.Printf("Listening on port %s", port)
	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%s", port), nil))
}
