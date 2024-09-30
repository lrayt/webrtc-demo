package main

import (
	"github.com/gin-gonic/gin"
	"github.com/gorilla/websocket"
	"log"
	"net/http"
	"time"
)

var upgrader = websocket.Upgrader{
	// 这个是校验请求来源
	// 在这里我们不做校验，直接return true
	CheckOrigin: func(r *http.Request) bool {
		return true
	},
}

func main() {
	r := gin.Default()
	r.GET("/ws", func(c *gin.Context) {
		// 将普通的http GET请求升级为websocket请求
		client, _ := upgrader.Upgrade(c.Writer, c.Request, nil)
		for {
			// 每隔两秒给前端推送一句消息“hello, WebSocket”
			err := client.WriteMessage(websocket.TextMessage, []byte("hello, WebSocket"))
			if err != nil {
				log.Println(err)
			}
			time.Sleep(time.Second * 2)
		}
	})
	if err := r.Run(":8080"); err != nil {
		log.Fatalf("gin run err:%s\n", err.Error())
	}
}
