package main

import (
	"bytes"
	"encoding/json"
	"io"
	"log"
	"net/http"
	"os"
	"time"
)

// PullRequest struct remains the same
type PullRequest struct {
	Name string `json:"name"`
}

func getOllamaHost() string {
	host := os.Getenv("OLLAMA_HOST")
	if host == "" {
		host = "http://localhost:11434"
	}
	return host
}

// ---- New Handler to List Models ----
func listModelsHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		http.Error(w, "Only GET method is allowed", http.StatusMethodNotAllowed)
		return
	}

	ollamaURL := getOllamaHost() + "/api/tags"

	client := &http.Client{Timeout: 30 * time.Second}
	resp, err := client.Get(ollamaURL)
	if err != nil {
		log.Printf("Error communicating with Ollama for listing models: %v", err)
		http.Error(w, "Failed to communicate with Ollama service", http.StatusInternalServerError)
		return
	}

	// Defensive check to prevent panic on resp.Body.Close()
	if resp != nil && resp.Body != nil {
		defer resp.Body.Close()
	} else {
		log.Println("Received a nil response or body from Ollama")
		http.Error(w, "Received an invalid response from Ollama", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(resp.StatusCode)
	io.Copy(w, resp.Body)
}

// pullModelHandler with more robust checks
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

	ollamaURL := getOllamaHost() + "/api/pull"
	log.Printf("Forwarding pull request to Ollama at: %s", ollamaURL)

	resp, err := http.Post(ollamaURL, "application/json", bytes.NewBuffer(ollamaReqBody))
	if err != nil {
		log.Printf("Error communicating with Ollama service: %v", err)
		http.Error(w, "Failed to communicate with Ollama service", http.StatusInternalServerError)
		return
	}

	// Defensive check to prevent panic on resp.Body.Close()
	if resp != nil && resp.Body != nil {
		defer resp.Body.Close()
	} else {
		log.Println("Received a nil response or body from Ollama")
		http.Error(w, "Received an invalid response from Ollama", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(resp.StatusCode)
	io.Copy(w, resp.Body)
}

func main() {
	http.HandleFunc("/api/models", listModelsHandler)
	http.HandleFunc("/api/models/pull", pullModelHandler)

	log.Println("Starting server on :8080...")
	if err := http.ListenAndServe(":8080", nil); err != nil {
		log.Fatalf("Could not start server: %s\n", err)
	}
}
