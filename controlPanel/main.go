package main

import (
	"log"
	"time"

	"gopkg.in/olahol/melody.v1"
)

var m = melody.New()

func main() {

	// StartServer()

	m := NewControlPanel()
	go m.Start()

	for {
		time.Sleep(time.Second * 2)
		log.Println("n :", m.robotNum)
	}
}
