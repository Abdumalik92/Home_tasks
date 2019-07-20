package main

import (
	"fmt"

	"github.com/gin-gonic/gin"

	_ "github.com/jinzhu/gorm/dialects/mysql"
)

func main() {
	r := gin.Default()
	r.LoadHTMLGlob("templates/*")
	r.GET("/", MainPage)
	r.POST("/auth", Login)
	r.POST("/chat", Chat)
	r.GET("/ws", func(c *gin.Context) {
		wshandler(c.Writer, c.Request)
	})

	go handleMessages()
	fmt.Println("Server is listening...")
	r.Run(":8080")
}
