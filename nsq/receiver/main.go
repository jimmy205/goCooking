package main

import (
	"log"

	"github.com/nsqio/go-nsq"
)

func main() {
	// r := make(chan bool)

	config := nsq.NewConfig()

	q, _ := nsq.NewConsumer("test", "test_channel", config)

	q.AddHandler(nsq.HandlerFunc(func(msg *nsq.Message) error {
		log.Println("got something :", string(msg.Body))
		// r <- true
		return nil
	}))

	err := q.ConnectToNSQLookupd("127.0.0.1:4161")
	if err != nil {
		log.Println("connect fail")
	}

	// <-r
}
