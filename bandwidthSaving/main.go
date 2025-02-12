package main

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func main() {
	r := mux.NewRouter()

	// WHIP endpoint
	r.HandleFunc("/whip", WHIPEndpoint).Methods("POST")

	// WHEP endpoint
	r.HandleFunc("/whep", WHEPEndpoint).Methods("GET")

	// Khởi động server
	log.Println("Server đang chạy tại :8080")
	log.Fatal(http.ListenAndServe(":8080", r))
}
