package main

import (
	"github.com/gin-gonic/gin"
	"github.com/olahol/melody"
	"net/http"
)

func main() {
	r := gin.Default()
	m := melody.New()

	// 设定 index.html
	r.GET("/", func(c *gin.Context) {
		http.ServeFile(c.Writer, c.Request, "index.html")
	})

	// 开始接收请求
	r.GET("/ws", func(c *gin.Context) {
		m.HandleRequest(c.Writer, c.Request)
	})

	// 当收到消息时，广播
	m.HandleMessage(func(s *melody.Session, msg []byte) {
		m.Broadcast(msg)
	})

	r.Run("")
}