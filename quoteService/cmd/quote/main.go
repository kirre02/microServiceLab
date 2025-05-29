package main

import (
	"encoding/json"
	"fmt"
	"math/rand"
	"net/http"
	"time"
)

type quote struct {
	ID     int    `json:"id"`
	Author string `json:"author"`
	Text   string `json:"text"`
}

var quotes = []quote{
	{1, "Albert Einstein", "Life is like riding a bicycle. To keep your balance you must keep moving."},
	{2, "Yoda", "Do or do not. There is no try."},
	{3, "Oscar Wilde", "Be yourself; everyone else is already taken."},
}

func quotesHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	q := quotes[rand.Intn(len(quotes))]
	err := json.NewEncoder(w).Encode(q)
	if err != nil {
		return
	}
}

func main() {
	rand.Seed(time.Now().UnixNano())

	http.HandleFunc("/quotes", quotesHandler)

	fmt.Println("Starting quote service on :8083...")
	err := http.ListenAndServe(":8083", nil)
	if err != nil {
		return
	}
}
