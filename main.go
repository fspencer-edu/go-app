package main

import (
	"encoding/json"
	"log"
	"net/http"
)

type MessageRequest struct {
	Name string `json:"name"`
}

type MessageResponse struct {
	Message string `json:"message"`
}

func enableCORS(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Access-Control-Allow-Origin", "*")
		w.Header().Set("Access-Control-Allow-Methods", "GET, POST, OPTIONS")
		w.Header().Set("Access-Control-Allow-Headers", "Content-Type")

		if r.Method == http.MethodOptions {
			w.WriteHeader(http.StatusOK)
			return
		}

		next.ServeHTTP(w, r)
	})
}

func greetHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "method not allowed", http.StatusMethodNotAllowed)
		return
	}

	var req MessageRequest
	err := json.NewDecoder(r.Body).Decode(&req)
	if err != nil {
		http.Error(w, "invalid JSON body", http.StatusBadRequest)
		return
	}

	if req.Name == "" {
		req.Name = "Guest"
	}

	resp := MessageResponse{
		Message: "Hello, " + req.Name + "! Your Go backend is working.",
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(resp)
}

func main() {
	mux := http.NewServeMux()

	// API route
	mux.HandleFunc("/api/greet", greetHandler)

	// Serve frontend
	fileServer := http.FileServer(http.Dir("./static"))
	mux.Handle("/", fileServer)

	log.Println("Server running at http://localhost:8080")

	err := http.ListenAndServe(":8080", enableCORS(mux))
	if err != nil {
		log.Fatal(err)
	}
}