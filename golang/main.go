package main

import (
	"github.com/gin-gonic/gin"
	"github.com/gorilla/websocket"
	"log"
	"net/http"
)

// 升级 HTTP 连接为 WebSocket 连接的配置
var upgrader = websocket.Upgrader{
	CheckOrigin: func(r *http.Request) bool {
		return true
	},
}

func main() {
	r := gin.Default()
	r.GET("/ws", func(c *gin.Context) {
		conn, err := upgrader.Upgrade(c.Writer, c.Request, nil)
		if err != nil {
			log.Fatalf("upgrade: %s\n", err)
		}
		defer conn.Close()
		for {
			conn.ReadJSON()
		}
	})
	if err := r.Run(":8080"); err != nil {
		log.Fatalf("gin run err:%s\n", err.Error())
	}
}
