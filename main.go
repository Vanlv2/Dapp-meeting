package main

import (
	"fmt"
	"net/http"
)

func main() {
	// Khởi tạo cuộc họp và tạo các participant giả lập
	meeting := &Meeting{
		Participants: make(map[string]*Participant),
	}

	// // Thêm các participant giả lập vào cuộc họp
	// meeting.Participants["John"] = &Participant{Name: "John", HasSpoken: false, HasAnswered: false, TasksCompleted: false, AttendanceTime: 50}

	// fmt.Println("== Trạng thái ban đầu của John ==")
	// statusJohn := meeting.CheckParticipantStatus("John")
	// fmt.Printf("Status of John: %+v\n", statusJohn)

	// // Cập nhật lại dữ liệu của John (giả sử thay đổi trạng thái tham gia)
	// meeting.UpdateParticipant("John", false, true, true, 60)

	// // Kiểm tra trạng thái sau khi cập nhật
	// fmt.Println("\n== Trạng thái sau khi cập nhật của John ==")
	// statusJohn = meeting.CheckParticipantStatus("John")
	// fmt.Printf("Status of John: %+v\n", statusJohn)

	// Đăng ký các handler cho HTTP kiểm tra với Postman
	http.HandleFunc("/update", updateParticipantHandler(meeting))
	http.HandleFunc("/status", checkParticipantStatusHandler(meeting))

	// Khởi chạy server trên cổng 8081
	fmt.Println("Server is running on port 8081")
	err := http.ListenAndServe(":8081", nil)
	if err != nil {
		fmt.Println("Error starting server:", err)
	}
}
