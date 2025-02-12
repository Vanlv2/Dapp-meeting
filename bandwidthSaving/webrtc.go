package main

import (
	"log"

	"github.com/pion/webrtc/v3"
)

// Tạo một peer connection mới
func createPeerConnection() (*webrtc.PeerConnection, error) {
	// Cấu hình WebRTC
	config := webrtc.Configuration{
		ICEServers: []webrtc.ICEServer{
			{
				URLs: []string{"stun:stun.l.google.com:19302"},
			},
		},
	}

	// Tạo một peer connection mới
	peerConnection, err := webrtc.NewPeerConnection(config)
	if err != nil {
		log.Println("Không thể tạo peer connection:", err)
		return nil, err
	}

	// Thêm các track vào peer connection (giả định rằng chúng là các stream video/audio)
	peerConnection.OnTrack(func(track *webrtc.TrackRemote, receiver *webrtc.RTPReceiver) {
		log.Println("Track received: ", track.Kind())
	})

	return peerConnection, nil
}
