package signaling

import (
	"github.com/gorilla/websocket"
	"net/http"
	"time"
)

const (
	// Time allowed to write a message to the peer.
	writeWait = 10 * time.Second

	// Time allowed to read the next pong message from the peer.
	pongWait = 60 * time.Second

	// Send pings to peer with this period. Must be less than pongWait.
	pingPeriod = (pongWait * 9) / 10

	// Maximum message size allowed from peer.
	maxMessageSize = 512
)

var (
	newline = []byte{'\n'}
	space   = []byte{' '}
)

var upgrader = websocket.Upgrader{
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,
	CheckOrigin: func(r *http.Request) bool {
		return true // 允许所有源
	},
}

type RoomInfo struct {
	Id        string `json:"id"`
	Name      string `json:"name"`
	CreatedAt string `json:"created_at"`
}

type Actions string

const (
	Open      Actions = "open"
	Join              = "join"
	Joined            = "joined"
	OtherJoin         = "other-join"
	Error             = "error"
	Leave             = "leave"
)

type Message struct {
	UID     string  `json:"uid"`
	Action  Actions `json:"action"`
	Content any     `json:"content"`
}
