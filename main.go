package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
)

func handler(w http.ResponseWriter, r *http.Request) { // This is what is displayed on the page
	fmt.Fprintf(w, "Hi there, I love %s!", r.URL.Path[1:])
}

func main() {
	CheckEnvironment()                                         // Check the environment variables are set correctly.
	SetEnvironment()                                           // Apply the environment variables.
	println("Server started on:" + os.Getenv("port"))          // Output
	http.HandleFunc("/", handler)                              // Serve requests from root
	log.Fatal(http.ListenAndServe(":"+os.Getenv("port"), nil)) // Begin serving requests
}
