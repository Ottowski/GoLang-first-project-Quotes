package main

import (
	"fmt"
	"log"
	"net/http"

	"go-project/api"
)

func main() {
	// Route
	http.HandleFunc("/quote", api.GetRandomQuote)

	port := 9090
	fmt.Printf("Random Quote API running on http://localhost:%d/quote\n", port)
	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%d", port), nil))

}

// http://localhost:8080/quote
