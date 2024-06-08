package main

import (
	"fmt"
	"log"
	"net/http"
	"strings"
	"html"
)

func main() {
	// Initialize the server
	server := &http.Server{
		Addr:    ":9002",
		Handler: nil, // Set the handler here
	}
	http.HandleFunc("/", greet)
	// Start the server
	if err := server.ListenAndServe(); err != nil {
		log.Fatal(err)
	} else {
		fmt.Println("Server started")
	}
}

func greet(w http.ResponseWriter, r *http.Request) {
	name := strings.Trim(r.URL.Path, "/")
	if name == "" {
		name = "Gopher adsfd"
	}

	fmt.Fprintf(w, "<!DOCTYPE html>\n")
	fmt.Fprintf(w, "%s, %s!\n", "Hello", html.EscapeString(name))
}