package main

import (
	"encoding/json"
	"net/http"
	"strings"
	"sync"
)

type Teacher struct {
	ID            int     `json:"id"`
	Name          string  `json:"name"`
	EmployeeID    string  `json:"employeeId"`
	Phone         string  `json:"phone"`
	Email         string  `json:"email"`
	Subject       string  `json:"subject"`
	AssignedClass string  `json:"assignedClass"`
	JoiningDate   string  `json:"joiningDate"`
	MonthlySalary float64 `json:"monthlySalary"`
}

var (
	teachers = []Teacher{
		{
			ID:            1,
			Name:          "Farzana Akter",
			EmployeeID:    "T-001",
			Phone:         "01711111111",
			Email:         "farzana.akter@bgpsc.edu.bd",
			Subject:       "Mathematics",
			AssignedClass: "Class 8 - Section A",
			JoiningDate:   "2022-01-10",
			MonthlySalary: 35000,
		},
	}
	nextTeacherID = 2
	teachersMu    sync.Mutex
)

func teachersHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json; charset=utf-8")

	switch r.Method {
	case http.MethodGet:
		teachersMu.Lock()
		defer teachersMu.Unlock()

		if err := json.NewEncoder(w).Encode(teachers); err != nil {
			http.Error(w, "Could not encode teachers", http.StatusInternalServerError)
		}

	case http.MethodPost:
		var teacher Teacher

		if err := json.NewDecoder(r.Body).Decode(&teacher); err != nil {
			http.Error(w, "Invalid teacher data", http.StatusBadRequest)
			return
		}

		teacher.Name = strings.TrimSpace(teacher.Name)
		teacher.EmployeeID = strings.TrimSpace(teacher.EmployeeID)
		teacher.Subject = strings.TrimSpace(teacher.Subject)
		teacher.AssignedClass = strings.TrimSpace(teacher.AssignedClass)

		if teacher.Name == "" || teacher.EmployeeID == "" || teacher.Subject == "" {
			http.Error(w, "Name, employee ID, and subject are required", http.StatusBadRequest)
			return
		}

		teachersMu.Lock()
		teacher.ID = nextTeacherID
		nextTeacherID++
		teachers = append(teachers, teacher)
		teachersMu.Unlock()

		w.WriteHeader(http.StatusCreated)

		if err := json.NewEncoder(w).Encode(teacher); err != nil {
			http.Error(w, "Could not encode teacher", http.StatusInternalServerError)
		}

	default:
		w.Header().Set("Allow", "GET, POST")
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
	}
}