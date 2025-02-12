package main

import (
	"log"
	"net/http"
	"sync"

	"github.com/gin-gonic/gin"
	"github.com/gorilla/websocket"
)

// Cấu hình WebSocket upgrader
var upgrader = websocket.Upgrader{
	CheckOrigin: func(r *http.Request) bool {
		return true
	},
}

// Dùng để lưu trữ thông tin signaling giữa các client
type SignalingServer struct {
	mu          sync.Mutex
	connections map[string]*websocket.Conn
}

// Khởi tạo signaling server
func NewSignalingServer() *SignalingServer {
	return &SignalingServer{
		connections: make(map[string]*websocket.Conn),
	}
}

// Hàm để thêm kết nối mới
func (s *SignalingServer) AddConnection(id string, conn *websocket.Conn) {
	s.mu.Lock()
	defer s.mu.Unlock()
	s.connections[id] = conn
}

// Hàm để gửi thông điệp giữa các client
func (s *SignalingServer) RelayMessage(id string, msgType int, message []byte) error {
	s.mu.Lock()
	defer s.mu.Unlock()

	for connID, conn := range s.connections {
		if connID != id {
			err := conn.WriteMessage(msgType, message)
			if err != nil {
				log.Println("Error sending message:", err)
				return err
			}
		}
	}
	return nil
}

func main() {
	router := gin.Default()
	signalingServer := NewSignalingServer()

	// Route xử lý WebSocket kết nối cho signaling
	router.GET("/ws/:id", func(c *gin.Context) {
		id := c.Param("id")
		conn, err := upgrader.Upgrade(c.Writer, c.Request, nil)
		if err != nil {
			log.Println("Failed to set websocket upgrade:", err)
			return
		}

		// Thêm kết nối mới vào signaling server
		signalingServer.AddConnection(id, conn)

		defer conn.Close()

		for {
			// Nhận message từ WebSocket
			msgType, message, err := conn.ReadMessage()
			if err != nil {
				log.Println("Read error:", err)
				break
			}

			// Gửi thông điệp đến peer khác
			err = signalingServer.RelayMessage(id, msgType, message)
			if err != nil {
				log.Println("Relay error:", err)
				break
			}
		}
	})

	// Chạy server tại port 8080
	router.Run(":8080")
}
