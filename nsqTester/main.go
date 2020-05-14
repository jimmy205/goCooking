package main

import (
	"log"
	"time"

	"github.com/nsqio/go-nsq"
)

// NSQHandler NSQHandler
type NSQHandler struct {
	producer *nsq.Producer
}

// DefaultNSQ 預設使用的NSQ
var DefaultNSQ *NSQHandler

func main() {

	InitNSQ()

	// DefaultNSQ.PingTestConsumer()
	// DefaultNSQ.PingTestProducer()
	DefaultNSQ.MoneyHandlerProducer()
	// DefaultNSQ.producer.Publish("nsq-test", []byte("ping"))

	time.Sleep(time.Minute * 30)

}

// InitNSQ 初始化NSQ
func InitNSQ() {
	config := nsq.NewConfig()

	nsqProducer, err := nsq.NewProducer("127.0.0.1:4150", config)
	if err != nil {
		log.Panic("[ InitNSQ 初始化NSQ失敗 ] - Err : [ ", err, " ]")
		return
	}

	DefaultNSQ = &NSQHandler{
		producer: nsqProducer,
	}

	return
}
