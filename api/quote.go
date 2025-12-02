package api

import (
	"encoding/json"
	"math/rand"
	"net/http"
	"time"
)

type Quote struct {
	Text   string `json:"text"`
	Author string `json:"author"`
}

var quotes = []Quote{
	{Text: "Be yourself; everyone else is already taken.", Author: "Oscar Wilde"},
	{Text: "Do what you can, with what you have, where you are.", Author: "Theodore Roosevelt"},
	{Text: "Life is what happens when you're busy making other plans.", Author: "John Lennon"},
	{Text: "In three words I can sum up everything I've learned about life: it goes on.", Author: "Robert Frost"},
}

func GetRandomQuote(w http.ResponseWriter, r *http.Request) {
	rand.Seed(time.Now().UnixNano())
	quote := quotes[rand.Intn(len(quotes))]

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(quote)
}
