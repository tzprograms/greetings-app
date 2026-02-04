package main

import (
	"encoding/json"
	"net/http"
)

// Define the structure of our response
type Message struct {
	Text string `json:"text"`
}

func main() {
	// w -> writer , r -> reader 
	http.HandleFunc("/api/data", func(w http.ResponseWriter, r *http.Request) {
		// 2. set header
		w.Header().Set("Content-Type", "application/json")
		
		msg := Message{Text: "Hello from the Go Backend!"}
		json.NewEncoder(w).Encode(msg)
	})
	
	http.ListenAndServe(":8080", nil)
}