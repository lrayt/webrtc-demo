package signaling

import "github.com/gorilla/websocket"

type Client struct {
	Id   string
	room *Room
	conn *websocket.Conn
	send chan []byte
}
