package main

import (
	"encoding/json"
	"log"
	"net/http"
)

func main() {
	mux := http.NewServeMux()
	mux.HandleFunc("GET /api/dashboard/summary", dashboardSummaryHandler)

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

func dashboardSummaryHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json; charset=utf-8")

	response := map[string]any{
		"totals": map[string]int{
			"students": 850,
			"teachers": 42,
			"staff":    18,
			"classes":  12,
			"sections": 24,
		},
		"attendance": map[string]int{
			"present": 812,
			"absent":   26,
			"late":     12,
		},
		"unpaidFees": map[string]any{
			"students": 35,
			"amount":   87500,
			"currency": "BDT",
		},
		"recentActivities": []map[string]string{
			{
				"time":        "09:15 AM",
				"description": "New student admission completed",
			},
			{
				"time":        "10:30 AM",
				"description": "Class 8 attendance recorded",
			},
			{
				"time":        "11:45 AM",
				"description": "Monthly tuition fee payment received",
			},
		},
	}

	if err := json.NewEncoder(w).Encode(response); err != nil {
		http.Error(w, "Could not encode dashboard summary", http.StatusInternalServerError)
	}
}