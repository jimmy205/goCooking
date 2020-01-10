package main

import (
	"log"

	"github.com/nsqio/go-nsq"
)

func main() {
	config := nsq.NewConfig()

	w, _ := nsq.NewProducer("127.0.0.1:4150", config)

	err := w.Publish("test", []byte("hello,world"))
	if err != nil {
		log.Println("err :", err)
	}

	w.Stop()
}
