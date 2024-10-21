package signaling

import "time"

type Room struct {
	id         string
	name       string
	CreateTime time.Time
	clients    map[string]*Client
	broadcast  chan []byte
	register   chan *Client
	unregister chan *Client
}

func NewRoom(name string) *Room {
	return &Room{
		id:         Message,
		name:       name,
		CreateTime: time.Now(),
		clients:    make(map[string]*Client),
		broadcast:  make(chan []byte),
		register:   make(chan *Client),
		unregister: make(chan *Client),
	}
}

func (r Room) Run() {
	for {
		select {
		case client := <-r.register:
			r.clients[client.Id] = client
			client.joined()
		case client := <-r.unregister:
			if _, ok := r.clients[client.Id]; ok {
				delete(r.clients, client.Id)
				close(client.send)
			}
			// todo: leave
			//case message := <-r.broadcast:

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
