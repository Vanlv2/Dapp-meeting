package main

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/pion/webrtc/v3"
)

// WHIPEndpoint nhận tín hiệu từ client (POST)
func WHIPEndpoint(w http.ResponseWriter, r *http.Request) {
	// Parse tín hiệu (offer) từ body của request
	var offer webrtc.SessionDescription
	err := json.NewDecoder(r.Body).Decode(&offer)
	if err != nil {
		http.Error(w, "Invalid offer", http.StatusBadRequest)
		return
	}

	// Tạo một peer connection mới
	peerConnection, err := createPeerConnection()
	if err != nil {
		http.Error(w, "Could not create peer connection", http.StatusInternalServerError)
		return
	}

	// Thiết lập offer từ client
	err = peerConnection.SetRemoteDescription(offer)
	if err != nil {
		log.Println("SetRemoteDescription failed:", err)
		http.Error(w, "Failed to set remote description", http.StatusInternalServerError)
		return
	}

	// Tạo answer và gửi về cho client
	answer, err := peerConnection.CreateAnswer(nil)
	if err != nil {
		log.Println("CreateAnswer failed:", err)
		http.Error(w, "Failed to create answer", http.StatusInternalServerError)
		return
	}

	err = peerConnection.SetLocalDescription(answer)
	if err != nil {
		log.Println("SetLocalDescription failed:", err)
		http.Error(w, "Failed to set local description", http.StatusInternalServerError)
		return
	}

	// Trả về answer cho client
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(answer)
}

// WHEPEndpoint gửi stream tới client (GET)
func WHEPEndpoint(w http.ResponseWriter, r *http.Request) {
	// Logic xử lý khi client yêu cầu stream (sử dụng WebRTC API để thiết lập egress)
	// Bạn có thể gửi stream video/audio thông qua WebRTC ở đây
}
