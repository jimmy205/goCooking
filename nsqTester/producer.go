package main

import (
	"log"
	"time"

	"github.com/tidwall/sjson"
)

// PingTestProducer 定期向cuckoo傳送訊息
func (n *NSQHandler) PingTestProducer() {

	t := time.NewTicker(time.Second)
	for {
		<-t.C
		err := n.producer.Publish("nsq-test", []byte("ping"))
		if err != nil {
			log.Printf("[ PingTestProducer - 連線失敗 ] Err : [ %s ]", err)
			return
		}
	}

}

// MoneyHandlerProducer 開洗分製造機
func (n *NSQHandler) MoneyHandlerProducer() {
	msg := []byte("{}")
	msg, _ = sjson.SetBytes(msg, "balance", 500.12)
	msg, _ = sjson.SetBytes(msg, "opcode", 300001)
	msg, _ = sjson.SetBytes(msg, "partnerID", 12345)
	msg, _ = sjson.SetBytes(msg, "userID", 12345)
	msg, _ = sjson.SetBytes(msg, "txNum", "thisIsTxNum")
	msg, _ = sjson.SetBytes(msg, "uuid", "thisIsUUID")

	t := time.NewTicker(time.Second)
	for {
		<-t.C
		log.Println("send")
		err := n.producer.Publish("nsq-test", msg)
		if err != nil {
			log.Printf("[ PingTestProducer - 連線失敗 ] Err : [ %s ]", err)
			return
		}
	}

}
