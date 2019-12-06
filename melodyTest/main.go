package main

import (
	"log"

	"github.com/tidwall/gjson"

	"github.com/gin-gonic/gin"
	"gopkg.in/olahol/melody.v1"
)

var r = gin.Default()
var m = melody.New()

func main() {
	HandleConnect()
	HanldeMessage()
	HanleDisconnect()
	goWs()
}

func goWs() {

	r.GET("/ws", func(c *gin.Context) {
		m.HandleRequest(c.Writer, c.Request)
	})

	r.Run(":8000")
}

func printName(msg string) {
	name := gjson.Get(msg, "name.last").String()
	log.Println("name : ", name)
}

// HanldeMessage 處理收到的訊息
func HanldeMessage() {
	m.HandleMessage(func(s *melody.Session, msg []byte) {
		log.Println("handle message")
		printName(string(msg))
	})
}

// HandleConnect 處理連線(連上)
func HandleConnect() {
	m.HandleConnect(func(s *melody.Session) {
		log.Println("someone connected")
	})
}

// HanleDisconnect 處理連線(離開)
func HanleDisconnect() {
	m.HandleDisconnect(func(s *melody.Session) {
		log.Println("someone leaving...")
	})
}
