package main

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
)

func echoHandler(w http.ResponseWriter, r *http.Request) {
	// Read body
	bodyBytes, err := io.ReadAll(r.Body)
	if err != nil {
		http.Error(w, "Failed to read body", http.StatusInternalServerError)
		return
	}

	// Build response
	response := map[string]interface{}{
		"method":  r.Method,
		"path":    r.URL.Path,
		"headers": r.Header,
		"body":    string(bodyBytes),
	}

	// Return JSON
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(response)
}

func main() {
	http.HandleFunc("/echo", echoHandler)
	fmt.Println("Starting WhoAmI service on :8082...")
	log.Fatal(http.ListenAndServe(":8082", nil))
}
