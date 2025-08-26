package main

import (
	"bytes"
	"encoding/json"
	"io"
	"log"
	"net/http"
	"os" // <-- Make sure "os" is imported
)

type PullRequest struct {
	Name string `json:"name"`
}

func getOllamaURL() string {
	// Get the Ollama API host from an environment variable.
	host := os.Getenv("OLLAMA_HOST")
	if host == "" {
		// Default for local, non-docker development
		host = "http://localhost:11434"
	}
	return host + "/api/pull"
}

func pullModelHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Only POST method is allowed", http.StatusMethodNotAllowed)
		return
	}
	var req PullRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, "Invalid request body", http.StatusBadRequest)
		return
	}
	ollamaReqBody, err := json.Marshal(req)
	if err != nil {
		http.Error(w, "Failed to create Ollama request", http.StatusInternalServerError)
		return
	}

	// Use the function to get the correct URL from the environment variable
	ollamaURL := getOllamaURL()
	log.Printf("Forwarding request to Ollama at: %s", ollamaURL)

	resp, err := http.Post(ollamaURL, "application/json", bytes.NewBuffer(ollamaReqBody))
	if err != nil {
		log.Printf("Error communicating with Ollama service: %v", err)
		http.Error(w, "Failed to communicate with Ollama service", http.StatusInternalServerError)
		return
	}
	defer resp.Body.Close()

	w.WriteHeader(resp.StatusCode)
	io.Copy(w, resp.Body)
}


func main() {
	http.HandleFunc("/api/models/pull", pullModelHandler)
	log.Println("Starting server on :8080...")
	if err := http.ListenAndServe(":8080", nil); err != nil {
		log.Fatalf("Could not start server: %s\n", err)
	}
}
