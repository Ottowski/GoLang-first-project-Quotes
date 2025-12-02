package api

import (
	"encoding/json"
	"math/rand"
	"net/http"
	"time"
)

// Quote struct
type Quote struct {
	Text   string `json:"text"`
	Author string `json:"author"`
}

// Initial quotes
var quotes = []Quote{
	{Text: "Be yourself; everyone else is already taken.", Author: "Oscar Wilde"},
	{Text: "Do what you can, with what you have, where you are.", Author: "Theodore Roosevelt"},
	{Text: "Life is what happens when you're busy making other plans.", Author: "John Lennon"},
	{Text: "In three words I can sum up everything I've learned about life: it goes on.", Author: "Robert Frost"},
}

// GetRandomQuote returns a random quote
func GetRandomQuote(w http.ResponseWriter, r *http.Request) {
	rand.Seed(time.Now().UnixNano())
	quote := quotes[rand.Intn(len(quotes))]

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(quote)
}

// GetAllQuotes returns all quotes
func GetAllQuotes(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(quotes)
}

// AddQuote adds a new quote via POST
func AddQuote(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	var newQuote Quote
	err := json.NewDecoder(r.Body).Decode(&newQuote)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	quotes = append(quotes, newQuote)

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(newQuote)
}
