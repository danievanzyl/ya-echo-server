package server

import (
	"encoding/json"
	"io"
	"net/http"
)

type EchoResponse struct {
	Method  string              `json:"method"`
	Headers map[string][]string `json:"headers"`
	Body    string              `json:"body"`
}

func echoHandler(w http.ResponseWriter, r *http.Request) {
	// Read the body
	bodyBytes, err := io.ReadAll(r.Body)
	if err != nil {
		http.Error(w, "Failed to read body", http.StatusInternalServerError)
		return
	}
	defer r.Body.Close()

	// Construct the response
	response := EchoResponse{
		Method:  r.Method,
		Headers: r.Header,
		Body:    string(bodyBytes),
	}

	// Set content type and write JSON response
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(response)
}
