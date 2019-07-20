package main

import (
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/gorilla/websocket"
	"github.com/jinzhu/gorm"
)

var onlineUsers = make(map[*websocket.Conn]string)
var users = make(map[*websocket.Conn]bool)
var broadcast = make(chan Message)

var wsupgrader = websocket.Upgrader{
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,
}
var user []User
var messages []Message
var dbs *gorm.DB

func MainPage(c *gin.Context) {
	c.HTML(http.StatusOK, "user.html", gin.H{
		"title": "Welcome",
	})
}
func Chat(c *gin.Context) {
	c.HTML(http.StatusOK, "chat.html", gin.H{})
}
func Login(c *gin.Context) {
	var err error
	var logs Log
	dbs, err = gorm.Open("mysql", "root:1111@tcp(localhost:3306)/chat?charset=utf8&parseTime=True&loc=Local")
	if err != nil {
		log.Panic("Couldn't opendatabase", err.Error())
	}
	name := c.PostForm("username")
	pass := c.PostForm("password")
	dbs.Find(&user)

	for i, _ := range user {

		if user[i].Username == name {
			if user[i].Password == pass {
				logs.Username = name

				dbs.Create(&logs)
				c.Redirect(307, "http://localhost:8080/chat")

				return
			}

		}
	}

	MainPage(c)
	c.JSON(400, "Wrong password")
	return
}

func wshandler(w http.ResponseWriter, r *http.Request) {
	conn, err := wsupgrader.Upgrade(w, r, nil)
	if err != nil {
		fmt.Println("Failed to set websocket upgrade: %+v", err)
		return
	}

	dbs, err = gorm.Open("mysql", "root:1111@tcp(localhost:3306)/chat?charset=utf8&parseTime=True&loc=Local")
	if err != nil {
		log.Panic("Couldn't opendatabase", err.Error())
	}
	var logs Log
	dbs.First(&logs)
	dbs.Delete(&logs)
	defer conn.Close()

	var journal []Journal
	onlineUsers[conn] = logs.Username
	users[conn] = true

	dbs.Find(&journal)

	for _, row := range journal {
		historyMsg := Message{
			Username: row.Username,
			Message:  row.Messag,
			Date:     row.Time,
		}
		conn.WriteJSON(historyMsg)
	}

	for {
		var msg Message
		// Read in a new message as JSON and map it to a Message object
		err := conn.ReadJSON(&msg)
		if err != nil {
			log.Printf("error: %v", err)
			delete(users, conn)
			delete(onlineUsers, conn)
			break
		}
		if msg.Message == "connect" {
			msg.Username = onlineUsers[conn]
			msg.Message = "test.conn"
			conn.WriteJSON(msg)
		} else {
			msg.Username = onlineUsers[conn]
			// Send the newly received message to the broadcast channel
			broadcast <- msg
		}
	}
}

func handleMessages() {
	var err error
	dbs, err = gorm.Open("mysql", "root:1111@tcp(localhost:3306)/chat?charset=utf8&parseTime=True&loc=Local")
	if err != nil {
		log.Panic("Couldn't opendatabase", err.Error())
	}
	now := time.Now().Format("02.01.2006 15:04:05")
	for {
		var history Journal
		// Grab the next message from the broadcast channel
		msg := <-broadcast
		if msg.Message != " is online" {
			history.Username = msg.Username
			history.Messag = msg.Message
			history.Time = now

			dbs.Create(&history)
		}

		// Send it out to every user that is currently connected
		for user := range users {
			msg.Date = now
			err := user.WriteJSON(msg)
			if err != nil {
				log.Printf("error: %v", err)
				user.Close()
				delete(users, user)
				delete(onlineUsers, user)
			}
		}
	}
}
