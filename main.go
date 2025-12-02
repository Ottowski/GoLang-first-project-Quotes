package main

import (
	"fmt"
	"log"
	"net/http"

	"go-project/api"
)

func main() {
	// Serve index.html at "/"
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		http.ServeFile(w, r, "index.html")
	})

	// API routes
	http.HandleFunc("/quote/random", api.GetRandomQuote)
	http.HandleFunc("/quotes", api.GetAllQuotes)
	http.HandleFunc("/quote/add", api.AddQuote)

	port := 9090
	fmt.Printf("Random Quote API running on http://localhost:%d\n", port)
	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%d", port), nil))
}
