package signaling

import (
	"github.com/gorilla/websocket"
	"log"
	"webrtc-demo/utils"
)

type Client struct {
	Id   string
	room *Room
	conn *websocket.Conn
	send chan *Message
}

func NewClient(conn *websocket.Conn) *Client {
	return &Client{
		Id:   utils.GetUUID(),
		conn: conn,
		send: make(chan *Message, 8),
	}
}

//func (c Client) error(msg string) {
//	if err := c.conn.WriteJSON(&Message{Action: Error, Content: msg}); err != nil {
//		log.Printf("write message err:%s\n", err.Error())
//	}
//	if err := c.conn.Close(); err != nil {
//		log.Printf("conn close err:%s", err.Error())
//	}
//}
//
//func (c Client) joined() {
//	//c.send <- []byte(fmt.Sprintf(`{"type": "%s","content":"%s"}`, Joined, c.Id))
//	if err := c.conn.WriteJSON(&Message{Action: Joined, Content: c.Id}); err != nil {
//		c.error(err.Error())
//	}
//}

func (c *Client) readPump() {
	for {
		msg := new(Message)
		if err := c.conn.ReadJSON(msg); err != nil {
			log.Printf("read message err:%s", err.Error())
		}
	}
}

func (c *Client) writePump() {
	for {
		select {
		case msg := <-c.send:
			if err := c.conn.WriteJSON(msg); err != nil {
				log.Printf("write message err:%s", err.Error())
			}
		}
	}
}
