package main

import (
	"log"

	"github.com/nsqio/go-nsq"
)

// PingTestConsumer 確認連線依然存在
func (n *NSQHandler) PingTestConsumer() {
	config := nsq.NewConfig()

	q, _ := nsq.NewConsumer("nsq-test", "test", config)

	q.AddConcurrentHandlers(nsq.HandlerFunc(func(msg *nsq.Message) error {
		log.Println("msg -> ", string(msg.Body))
		return nil
	}), 10)

	err := q.ConnectToNSQLookupd("127.0.0.1:4161")
	if err != nil {
		log.Println("err -> ", err)
		return
	}

}

// MoneyRertyHandler 開洗分重試
// func (n *NSQHandler) MoneyRertyHandler() {
// 	config := nsq.NewConfig()

// 	q, _ := nsq.NewConsumer("nsq-test", "test", config)

// 	q.AddConcurrentHandlers(nsq.HandlerFunc(func(msg *nsq.Message) error {
// 		log.Println("msg -> ", string(msg.Body))
// 		return nil
// 	}), 10)

// 	err := q.ConnectToNSQLookupd("127.0.0.1:4161")
// 	if err != nil {
// 		log.Println("err -> ", err)
// 		return
// 	}
// }
