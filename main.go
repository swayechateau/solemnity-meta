package main

import (
	"fmt"
	"net/http"
)

func main() {
	// Serve static files from the "public" directory
	fs := http.FileServer(http.Dir("public"))
	http.HandleFunc("/", metaHandler)
	http.Handle("/docs/", http.StripPrefix("/docs/", fs))

	port := "8080"
	fmt.Printf("Server is running on port %s...\n", port)
	err := http.ListenAndServe(":"+port, nil)
	if err != nil {
		fmt.Printf("Error: %v\n", err)
	}
}

func metaHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello, World!")
}
