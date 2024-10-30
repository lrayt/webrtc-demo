package signaling

import (
	"fmt"
	"time"
	"webrtc-demo/utils"
)

type Room struct {
	id         string
	name       string
	CreateTime time.Time
	clients    map[string]*Client
	broadcast  chan *Message
	register   chan *Client
	unregister chan *Client
}

func NewRoom(name string) *Room {
	room := &Room{
		id:         utils.GetUUID(),
		name:       name,
		CreateTime: time.Now(),
		clients:    make(map[string]*Client),
		broadcast:  make(chan *Message, 8),
		register:   make(chan *Client),
		unregister: make(chan *Client),
	}
	go room.Run()
	return room
}

func (r Room) Run() {
	for {
		select {
		case client, ok := <-r.register:
			if !ok {
				continue
			}
			r.clients[client.Id] = client
			client.send <- &Message{UID: client.Id, Action: Joined}
			if len(r.clients) > 1 {
				r.broadcast <- &Message{UID: client.Id, Action: OtherJoin}
			}
		case client := <-r.unregister:
			if _, ok := r.clients[client.Id]; ok {
				delete(r.clients, client.Id)
				close(client.send)
			}
			r.broadcast <- &Message{UID: client.Id, Action: Leave}
		case message, ok := <-r.broadcast:
			if !ok {
				continue
			}
			for _, client := range r.clients {
				fmt.Println("--->", message.UID, client.Id)
				if message.UID != client.Id {
					client.send <- message
				}
			}
		}
	}
}

func (r Room) RoomInfo() *RoomInfo {
	return &RoomInfo{
		Id:        r.id,
		Name:      r.name,
		CreatedAt: r.CreateTime.String(),
	}
}
