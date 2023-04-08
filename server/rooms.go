package server

import (
	"log"
	"math/rand"
	"sync"

	"github.com/gorilla/websocket"
)

type Participant struct {
	Host bool
	Conn *websocket.Conn
}
type RoomMap struct {
	Mutex sync.RWMutex
	Map   map[string][]Participant
}

func (r *RoomMap) Init() {
	r.Map = make(map[string][]Participant)
}
func (r *RoomMap) Get(roomId string) []Participant {
	r.Mutex.RLock()
	defer r.Mutex.Unlock()
	return r.Map[roomId]
}
func (r *RoomMap) CreateRoom() string {
	var letters = []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ1234567890")
	var unique_str = make([]rune, 8)
	for i := range unique_str {
		unique_str[i] = letters[rand.Intn(len(letters))]
	}
	roomId := string(unique_str)
	r.Map[roomId] = []Participant{}
	return roomId
}
func (r *RoomMap) InsertRoom(roomId string, host bool, conn *websocket.Conn) {
	p := Participant{host, conn}
	log.Println("Inserting to Room with roomId :", roomId)
	r.Map[roomId] = append(r.Map[roomId], p)
}
func (r *RoomMap) DeleteRoom(roomId string) {
	delete(r.Map, roomId)
}
