package nsq

import (
	"goPra/corvus/constant"
	"goPra/corvus/package/load"

	"log"

	"github.com/nsqio/go-nsq"
)

// Producer NSQ的傳送者
type Producer struct {
	producer *nsq.Producer
	customer *nsq.Consumer
}

// NewProducer 新的傳送NSQ的人
func NewProducer() *Producer {
	// 讀取設定檔
	conf := load.LoadConfig()

	config := nsq.NewConfig()
	nsqProducer, err := nsq.NewProducer(conf.NSQ.NSQNode, config)
	if err != nil {
		log.Println("Init NSQ Receiver Err :", err)
		return nil
	}

	np := &Producer{
		producer: nsqProducer,
	}

	return np
}

// SendToBackground 傳送給背景的確認訊息
func (p *Producer) SendToBackground(topic constant.NsqTopic, msg []byte) {
	err := p.producer.Publish(topic.String(), msg)
	if err != nil {
		log.Println("NSQ Producer Send MSG ERR :", err)
		return
	}
}
