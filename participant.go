package main

// Participant chứa thông tin về người tham gia cuộc họp
type Participant struct {
	Name           string // Tên người tham gia
	HasSpoken      bool   // Phát biểu trong cuộc hợp
	HasAnswered    bool   // Đã trả lời câu hỏi đặt ra chưa
	TasksCompleted bool   // Hoàn thành nhiệm vụ được giao(tham gia thảo luận, mini game,...)
	AttendanceTime int    // Thời gian tham gia cuộc họp (phút)
}
