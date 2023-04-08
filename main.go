package main

import (
	"log"
	"net/http"

	"github.com/vyas-git/go-meet-app/server"
)

func main() {
	http.HandleFunc("/create", server.CreateRoomRequestHandler)
	log.Println("Server listening on 8080 ..")
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		log.Fatal(err)
	}
}
