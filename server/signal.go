package server

import (
	"fmt"
	"net/http"
)

func CreateRoomRequestHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "Hii")
}
