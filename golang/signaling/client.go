package signaling

import (
	"fmt"
	"github.com/gorilla/websocket"
	"log"
	"time"
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

func (c Client) error(msg string) {
	if err := c.conn.WriteJSON(&Message{Action: Error, Content: msg}); err != nil {
		log.Printf("write message err:%s\n", err.Error())
	}
	if err := c.conn.Close(); err != nil {
		log.Printf("conn close err:%s", err.Error())
	}
}

func (c Client) joined() {
	//c.send <- []byte(fmt.Sprintf(`{"type": "%s","content":"%s"}`, Joined, c.Id))
	if err := c.conn.WriteJSON(&Message{Action: Joined, Content: c.Id}); err != nil {
		c.error(err.Error())
	}
}

func (c *Client) readPump() {
	defer func() {
		c.room.unregister <- c
		c.conn.Close()
	}()
	c.conn.SetReadLimit(maxMessageSize)
	c.conn.SetReadDeadline(time.Now().Add(pongWait))
	c.conn.SetPongHandler(func(string) error { c.conn.SetReadDeadline(time.Now().Add(pongWait)); return nil })
	for {
		msg := new(Message)
		if err := c.conn.ReadJSON(msg); err != nil {
			log.Printf("error: %s\n", err.Error())
		}
		switch msg.Action {
		case Join:
			c.room.register <- c
			c.send <- &Message{UID: c.Id, Action: Joined}
			c.room.broadcast <- &Message{UID: c.Id, Action: OtherJoin}
		default:
			fmt.Println("message:", msg.Action)
		}

		//_, message, err := c.conn.ReadMessage()
		//if err != nil {
		//	if websocket.IsUnexpectedCloseError(err, websocket.CloseGoingAway, websocket.CloseAbnormalClosure) {
		//		log.Printf("error: %v", err)
		//	}
		//	break
		//}
		//message = bytes.TrimSpace(bytes.Replace(message, newline, space, -1))
		fmt.Println("message:", msg)
		//c.room.broadcast <- message
	}
}

func (c *Client) writePump() {
	ticker := time.NewTicker(pingPeriod)
	defer func() {
		ticker.Stop()
		c.conn.Close()
	}()
	for {
		select {
		case message := <-c.send:
			if err := c.conn.WriteJSON(message); err != nil {
				log.Printf("write message err:%s\n", err.Error())
			}
			//c.conn.SetWriteDeadline(time.Now().Add(writeWait))
			//if !ok {
			//	// The hub closed the channel.
			//	c.conn.WriteMessage(websocket.CloseMessage, []byte{})
			//	return
			//}
			//
			//w, err := c.conn.NextWriter(websocket.TextMessage)
			//if err != nil {
			//	return
			//}
			//w.Write(message)
			//
			//// Add queued chat messages to the current websocket message.
			//n := len(c.send)
			//for i := 0; i < n; i++ {
			//	w.Write(newline)
			//	w.Write(<-c.send)
			//}
			//
			//if err := w.Close(); err != nil {
			//	return
			//}
		case <-ticker.C:
			c.conn.SetWriteDeadline(time.Now().Add(writeWait))
			if err := c.conn.WriteMessage(websocket.PingMessage, nil); err != nil {
				return
			}
		}
	}
}
