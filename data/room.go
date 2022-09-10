package data

import (
	"fmt"
	"github.com/gorilla/websocket"
)

var (
	rooms = make(map[string]*Room, 0)
)

const (
	MainRoom = "Main"
)

type Room struct {
	Name     string
	Messages []ReceivedMessage
	Channel  chan ReceivedMessage
	Sockets  []*websocket.Conn
}

func NewRoom(name string) *Room {
	return &Room{
		Name:     name,
		Messages: make([]ReceivedMessage, 0),
		Channel:  make(chan ReceivedMessage, 10),
	}
}

func AddRoom(room *Room) {
	rooms[room.Name] = room
}

func GetRoomByName(name string) *Room {
	return rooms[name]
}

func (r *Room) AnnounceNewUser(user User) {
	r.Messages = append(r.Messages, NewReceivedMessage(user, fmt.Sprintf("User %s logged on.", user.Name)))
	r.Channel <- r.Messages[len(r.Messages)-1]
}

func (r *Room) Listen() {
	go func() {
		for {
			i := <-r.Channel
			r.Broadcast(i)
		}
	}()
}

func (r *Room) Broadcast(message ReceivedMessage) {
	for _, ws := range r.Sockets {
		ws.WriteJSON(message)
	}
}

func (r *Room) RegisterSocket(ws *websocket.Conn) {
	r.Sockets = append(r.Sockets, ws)
}

func (r *Room) UnregisterSocket(ws *websocket.Conn) {
	for index, s := range r.Sockets {
		if s == ws {
			r.Sockets = append(r.Sockets[0:index], r.Sockets[index+1:]...)
		}
	}
}
