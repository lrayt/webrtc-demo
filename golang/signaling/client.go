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
	client := &Client{
		Id:   utils.GetUUID(),
		conn: conn,
		send: make(chan *Message, 8),
	}
	go client.writePump()
	go client.readPump()
	return client
}

func (c *Client) readPump() {
	defer func() {
		if c.room != nil {
			c.room.unregister <- c
		}
	}()
	for {
		msg := new(Message)
		err := c.conn.ReadJSON(msg)
		if websocket.IsCloseError(err, websocket.CloseGoingAway) {
			log.Println("client disconnected")
			break
		} else if err != nil {
			log.Printf("read message err:%s", err.Error())
			break
		}
	}
}

func (c *Client) writePump() {
	defer func() {
		c.conn.Close()
	}()

	for {
		select {
		case msg, ok := <-c.send:
			if !ok {
				c.conn.WriteMessage(websocket.CloseMessage, []byte{})
				return
			}
			if err := c.conn.WriteJSON(msg); err != nil {
				log.Printf("write message err:%s", err.Error())
			}
		}
	}
}

func (c Client) userInfo() *UserInfo {
	return &UserInfo{
		Id:        c.Id,
		Name:      "Mary",
		CreatedAt: "2024-09-20",
	}
}
