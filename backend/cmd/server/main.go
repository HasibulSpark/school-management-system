package main

import (
	"log"
	"net/http"
)

func main() {
	mux := http.NewServeMux()

	mux.HandleFunc("GET /api/health", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json; charset=utf-8")
		_, err := w.Write([]byte(`{"status":"ok","message":"School management API is running"}`))
		if err != nil {
			log.Printf("Could not write health response: %v", err)
		}
	})

	log.Println("Go API listening at http://localhost:8080")
	log.Fatal(http.ListenAndServe(":8080", mux))
}