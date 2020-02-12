package nsq

import (
	"goPra/corvus/constant"
	"goPra/corvus/package/load"
	"log"

	"github.com/nsqio/go-nsq"
)

// SetReceiverNSQ 設定NSQ
func SetReceiverNSQ(topic constant.NsqTopic, channel constant.NsqChannel, fn func(msg []byte)) {
	conf := load.LoadConfig()

	nsqConf := nsq.NewConfig()
	q, _ := nsq.NewConsumer(topic.String(), channel.String(), nsqConf)
	q.AddHandler(nsq.HandlerFunc(func(msg *nsq.Message) error {
		fn(msg.Body)
		return nil
	}))

	nsqConnErr := q.ConnectToNSQLookupd(conf.NSQ.NSQLookUpURL)
	if nsqConnErr != nil {
		log.Println("nsqConnErr :", nsqConnErr)
	}
}
