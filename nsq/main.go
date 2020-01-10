package main

import (
	"encoding/json"
	"log"
	"time"

	"github.com/nsqio/go-nsq"
)

func main() {

	go producer()
	// go receiver()

	time.Sleep(time.Minute)
}

func producer() {
	config := nsq.NewConfig()

	p, _ := nsq.NewProducer("127.0.0.1:4150", config)
	defer p.Stop()

	t := time.NewTicker(time.Second)
	defer t.Stop()

	type PingInfo struct {
		Action string `json:"action"`  // 動作
		UUID   string `json:"uuid"`    // 使用者的uuid
		UserID int    `json:"user_id"` // 使用者ID
	}

	m := PingInfo{
		Action: "ping",
		UUID:   "d5a920ce-ea8a-4cb8-9c26-b119ece2c88e",
		UserID: 129,
	}

	b, _ := json.Marshal(m)

	for {
		<-t.C
		// u := uuid.New()
		err := p.Publish("check_connect", b)
		if err != nil {
			log.Println("err :", err)
		}
	}

}

func receiver() {

	r := make(chan bool)
	config := nsq.NewConfig()

	c, _ := nsq.NewConsumer("test", "test_channel", config)

	c.AddHandler(nsq.HandlerFunc(func(msg *nsq.Message) error {

		log.Println("msg :", string(msg.Body))
		return nil
	}))

	err := c.ConnectToNSQLookupd("127.0.0.1:4161")
	if err != nil {
		log.Println("err :", err)
	}

	<-r
}
