package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"
)

// Handler để cập nhật người tham gia thông qua request
func updateParticipantHandler(m *Meeting) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		name := r.URL.Query().Get("name")

		// Kiểm tra nếu participant không tồn tại
		if _, exists := m.Participants[name]; !exists {
			http.Error(w, "Participant does not exist", http.StatusNotFound)
			return
		}

		// Cập nhật từng trường nếu có trong request
		if spoken := r.URL.Query().Get("spoken"); spoken != "" {
			m.Participants[name].HasSpoken = spoken == "true"
		}

		if answered := r.URL.Query().Get("answered"); answered != "" {
			m.Participants[name].HasAnswered = answered == "true"
		}

		if tasksCompleted := r.URL.Query().Get("tasksCompleted"); tasksCompleted != "" {
			m.Participants[name].TasksCompleted = tasksCompleted == "true"
		}

		if attendanceTime := r.URL.Query().Get("attendanceTime"); attendanceTime != "" {
			if time, err := strconv.Atoi(attendanceTime); err == nil {
				m.Participants[name].AttendanceTime = time
			} else {
				http.Error(w, "Invalid attendance time", http.StatusBadRequest)
				return
			}
		}

		fmt.Fprintf(w, "Participant %s updated successfully", name)
	}
}

// Handler để kiểm tra trạng thái của người tham gia
func checkParticipantStatusHandler(m *Meeting) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		name := r.URL.Query().Get("name")
		if name == "" {
			http.Error(w, "Missing 'name' parameter", http.StatusBadRequest)
			return
		}

		// Gọi hàm CheckParticipantStatus để kiểm tra tình trạng của người tham gia
		status := m.CheckParticipantStatus(name)

		// Trả kết quả về dưới dạng JSON
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(status)
	}
}
