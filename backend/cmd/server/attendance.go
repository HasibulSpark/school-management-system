package main

import (
	"encoding/json"
	"net/http"
	"strings"
	"sync"
)

type AttendanceRecord struct{
	ID          int    `json:"id"`
	Date        string `json:"date"`
	StudentName string `json:"studentName"`
	Roll        string `json:"roll"`
	Class       string `json:"class"`
	Section     string `json:"section"`
	Status      string `json:"status"`
}

var(attendanceRecords = []AttendanceRecord{
		{
			ID:          1,
			Date:        "2026-11-02",
			StudentName: "Arafat Rahman",
			Roll:        "01",
			Class:       "8",
			Section:     "A",
			Status:      "present",
		},
		{
			ID:          2,
			Date:        "2026-11-02",
			StudentName: "Nusrat Jahan",
			Roll:        "12",
			Class:       "7",
			Section:     "B",
			Status:      "late",
		},
	}
	nextAttendanceID = 3
	attendanceMu     sync.Mutex
)

func attendanceHandler(w http.ResponseWriter, r *http.Request){
	w.Header().Set("Content-Type", "application/json; charset=utf-8")

	switch r.Method{
	case http.MethodGet:
		attendanceMu.Lock()
		defer attendanceMu.Unlock()

		if err := json.NewEncoder(w).Encode(attendanceRecords); err != nil{
			http.Error(w, "Could not encode attendance records", http.StatusInternalServerError)
		}

	case http.MethodPost:
		var record AttendanceRecord

		if err := json.NewDecoder(r.Body).Decode(&record); err != nil{
			http.Error(w, "Invalid attendance data", http.StatusBadRequest)
			return
		}

		record.Date = strings.TrimSpace(record.Date)
		record.StudentName = strings.TrimSpace(record.StudentName)
		record.Roll = strings.TrimSpace(record.Roll)
		record.Class = strings.TrimSpace(record.Class)
		record.Section = strings.TrimSpace(record.Section)
		record.Status = strings.ToLower(strings.TrimSpace(record.Status))

		if record.Date == "" || record.StudentName == "" || record.Class == "" || record.Section == "" {
			http.Error(w, "Date, student name, class, and section are required", http.StatusBadRequest)
			return
		}

		if record.Status != "present" && record.Status != "absent" && record.Status != "late" {
			http.Error(w, "Status must be present, absent, or late", http.StatusBadRequest)
			return
		}

		attendanceMu.Lock()
		record.ID = nextAttendanceID
		nextAttendanceID++
		attendanceRecords = append(attendanceRecords, record)
		attendanceMu.Unlock()

		w.WriteHeader(http.StatusCreated)

		if err := json.NewEncoder(w).Encode(record); err != nil{
			http.Error(w, "Could not encode attendance record", http.StatusInternalServerError)
		}

	default:
		w.Header().Set("Allow", "GET, POST")
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
	}
}