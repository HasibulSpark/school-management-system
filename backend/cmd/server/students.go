package main

import (
	"encoding/json"
	"net/http"
	"strings"
	"sync"
)

type Student struct {
	ID            int    `json:"id"`
	Name          string `json:"name"`
	DateOfBirth   string `json:"dateOfBirth"`
	Class         string `json:"class"`
	Section       string `json:"section"`
	Roll          string `json:"roll"`
	GuardianName  string `json:"guardianName"`
	GuardianPhone string `json:"guardianPhone"`
	Address       string `json:"address"`
}

var (
	students = []Student{
		{
			ID:            1,
			Name:          "Arafat Rahman",
			DateOfBirth:   "2010-05-12",
			Class:         "8",
			Section:       "A",
			Roll:          "01",
			GuardianName:  "Md. Rahman",
			GuardianPhone: "01700000000",
			Address:       "Sylhet, Bangladesh",
		},
	}
	nextStudentID = 2
	studentsMu    sync.Mutex
)

func studentsHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json; charset=utf-8")

	switch r.Method {
	case http.MethodGet:
		studentsMu.Lock()
		defer studentsMu.Unlock()

		if err := json.NewEncoder(w).Encode(students); err != nil {
			http.Error(w, "Could not encode students", http.StatusInternalServerError)
		}

	case http.MethodPost:
		var student Student

		if err := json.NewDecoder(r.Body).Decode(&student); err != nil {
			http.Error(w, "Invalid student data", http.StatusBadRequest)
			return
		}

		student.Name = strings.TrimSpace(student.Name)
		student.Class = strings.TrimSpace(student.Class)
		student.Section = strings.TrimSpace(student.Section)
		student.Roll = strings.TrimSpace(student.Roll)

		if student.Name == "" || student.Class == "" || student.Section == "" || student.Roll == "" {
			http.Error(w, "Name, class, section, and roll are required", http.StatusBadRequest)
			return
		}

		studentsMu.Lock()
		student.ID = nextStudentID
		nextStudentID++
		students = append(students, student)
		studentsMu.Unlock()

		w.WriteHeader(http.StatusCreated)

		if err := json.NewEncoder(w).Encode(student); err != nil {
			http.Error(w, "Could not encode student", http.StatusInternalServerError)
		}

	default:
		w.Header().Set("Allow", "GET, POST")
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
	}
}