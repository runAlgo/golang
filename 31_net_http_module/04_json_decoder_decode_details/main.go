package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strings"
	"time"
)
// Implementing a basic HTTP server with a JSON-based REST API endpoint, 
// and it demonstrates the core concept of handling HTTP requests, decoding
// JSON from the requestt body into a GO struct, validating the input, and
// sending a JSON response back to the client. This is one of the foundational
// building blocks for creating APIs and web services in Go.

func writeJSON(w http.ResponseWriter, status int, data any) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)

	_ = json.NewEncoder(w).Encode(data)
}

type TestRequeset struct {
	Name string `json:"name"`
}

func decodeHandler(w http.ResponseWriter, r *http.Request) {

	if r.Method != http.MethodPost {
		writeJSON(w, http.StatusMethodNotAllowed, map[string]any{
			"ok":    false,
			"error": "Only post is allowed",
		})
		return
	}

	// close the request body
	defer r.Body.Close()

	var req TestRequeset

	dec := json.NewDecoder(r.Body)

	if err := dec.Decode(&req); err != nil {
		writeJSON(w, http.StatusBadRequest, map[string]any{
			"ok":    false,
			"error": "Invalid json format",
		})
		return
	}

	req.Name = strings.TrimSpace(req.Name)

	if req.Name == "" {
		writeJSON(w, http.StatusBadRequest, map[string]any{
			"ok":    false,
			"error": "All inputs are required!",
		})
		return
	}

	writeJSON(w, http.StatusOK, map[string]any{
		"ok":        true,
		"data":      req,
		"timeStamp": time.Now().UTC(),
	})
}

func main() {
	http.HandleFunc("/ok", decodeHandler)
	fmt.Println("Switching to PORT:3000")
	err := http.ListenAndServe(":3000", nil)
	if err != nil {
		log.Fatalf("Server failed to start %v", err)
	}
}
