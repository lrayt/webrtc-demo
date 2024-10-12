package signaling

import "fmt"

type Room struct {
	clients    map[string]*Client
	broadcast  chan []byte
	register   chan *Client
	unregister chan *Client
}

func NewRoom() *Room {
	return &Room{
		clients:    make(map[string]*Client),
		broadcast:  make(chan []byte),
		register:   make(chan *Client),
		unregister: make(chan *Client),
	}
}

func (r *Room) Run() {
	for {
		select {
		case client := <-r.register:
			r.clients[client.Id] = client
			client.send <- []byte(fmt.Sprintf(`{"event":"joined","content":"%s"'}`, client.Id))
		case client := <-r.unregister:
			if _, ok := r.clients[client.Id]; ok {
				delete(r.clients, client.Id)
				close(client.send)
			}
			// todo: leave
		case message := <-r.broadcast:

		}
	}
}
