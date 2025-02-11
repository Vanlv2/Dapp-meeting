package main

import (
	"fmt"
	"sync"
)

// Meeting struct lưu trữ danh sách người tham gia và quản lý đồng thời
type Meeting struct {
	Participants map[string]*Participant
	mu           sync.Mutex
}

// CheckParticipantStatus kiểm tra trạng thái của một người tham gia
func (m *Meeting) CheckParticipantStatus(participantName string) map[string]bool {
	m.mu.Lock()
	defer m.mu.Unlock()

	status := map[string]bool{
		"HasSpoken":      false,
		"HasAnswered":    false,
		"TasksCompleted": false,
		"AttendanceTime": false,
	}

	p, exists := m.Participants[participantName]
	if !exists {
		fmt.Printf("Participant %s does not exist.\n", participantName)
		return status
	}

	fmt.Printf("Checking status of %s: %+v\n", participantName, p) // In ra thông tin participant để kiểm tra

	status["HasSpoken"] = p.HasSpoken
	status["HasAnswered"] = p.HasAnswered
	status["TasksCompleted"] = p.TasksCompleted
	status["AttendanceTime"] = p.AttendanceTime >= 80

	fmt.Printf("Participant %s status: %+v\n", participantName, status)

	return status
}
