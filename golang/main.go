package main

import (
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
	"webrtc-demo/signaling"
)

func main() {
	r := gin.Default()
	hub := signaling.NewRoomHub()

	r.Use(func(c *gin.Context) {
		method := c.Request.Method
		origin := c.Request.Header.Get("Origin")
		if origin != "" {
			c.Header("Access-Control-Allow-Origin", "*") // 可将将 * 替换为指定的域名
			c.Header("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, DELETE, UPDATE")
			c.Header("Access-Control-Allow-Headers", "Origin, X-Requested-With, Content-Type, Accept, Authorization")
			c.Header("Access-Control-Expose-Headers", "Content-Length, Access-Control-Allow-Origin, Access-Control-Allow-Headers, Cache-Control, Content-Language, Content-Type")
			c.Header("Access-Control-Allow-Credentials", "true")
		}
		if method == "OPTIONS" {
			c.AbortWithStatus(http.StatusNoContent)
		}
		c.Next()
	})

	rg := r.Group("api/v1/room")
	{
		rg.GET("all", hub.All)
		rg.POST("add", hub.Add)
		rg.POST("del", hub.Del)
	}
	r.GET("/ws", hub.WS)
	if err := r.Run(":8080"); err != nil {
		log.Fatalf("server run error %s", err.Error())
	}
}
