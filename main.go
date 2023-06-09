package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
)

func handler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hi there, I love %s!", r.URL.Path[1:])
}

func main() {
	http.HandleFunc("/", handler)
	log.Fatal(http.ListenAndServe(determinePort(), nil))
}

func determinePort() string {
	port := os.Getenv("TIPOLA_PORT")
	if port == "" {
		port = ":8999"
	}
	return port
}
