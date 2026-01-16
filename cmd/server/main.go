package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
)

const (
	version = "0.1.0"
	port    = "8545"
)

func main() {
	// Health check endpoint
	http.HandleFunc("/health", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		fmt.Fprintf(w, `{"status":"healthy","version":"%s","service":"go-pvm-rpc-server"}`, version)
	})

	// Root endpoint
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		if r.Method == "GET" {
			fmt.Fprintf(w, "Go-PVM RPC Server v%s", version)
			return
		}

		if r.Method == "POST" {
			// TODO: Implement RPC handler
			w.Header().Set("Content-Type", "application/json")
			fmt.Fprintf(w, `{"error":"RPC handler not yet implemented"}`)
			return
		}

		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
	})

	serverPort := os.Getenv("PORT")
	if serverPort == "" {
		serverPort = port
	}

	addr := fmt.Sprintf(":%s", serverPort)
	log.Printf("Go-PVM RPC Server v%s running on http://localhost:%s", version, serverPort)

	if err := http.ListenAndServe(addr, nil); err != nil {
		log.Fatal(err)
	}
}
