package main

import (
	"github.com/gin-gonic/gin"
	"github.com/tidwall/gjson"
	"gopkg.in/olahol/melody.v1"
)

// ControlPanel 控制面板
type ControlPanel struct {
	m        *melody.Melody
	robotNum int64
}

// NewControlPanel 新創一個控制面板
func NewControlPanel() ControlPanel {
	return ControlPanel{
		m: melody.New(),
	}
}

// Start 開始一個控制面板
func (c *ControlPanel) Start() {
	r := gin.Default()

	r.GET("/ws", func(c *gin.Context) {
		m.HandleRequest(c.Writer, c.Request)
	})

	m.HandleMessage(func(s *melody.Session, msg []byte) {
		c.ControlEntry(s, string(msg))
	})

	r.Run(":8880")
}

// ControlEntry 控制的入口
func (c *ControlPanel) ControlEntry(s *melody.Session, msg string) {

	action := gjson.Get(msg, "action").String()

	switch action {
	case "addMember":
		n := gjson.Get(msg, "number").Int()
		c.robotNum += n

	case "deletMember":
		n := gjson.Get(msg, "number").Int()
		c.robotNum -= n
	}

}
