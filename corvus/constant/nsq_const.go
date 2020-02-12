package constant

// NsqTopic 有設定的TopicName
type NsqTopic string

// NSQ Topic 的名稱設定
const (
	NsqTopicServerMonitor = NsqTopic("serverMonitor") // NSQ 服務監視器
	NsqTopicCorvusMonitor = NsqTopic("corvusMonitor") // NSQ 服務監視器
)

// IsValid 確認Topic是否有用
func (nt NsqTopic) IsValid() bool {
	switch nt {
	case NsqTopicServerMonitor,
		NsqTopicCorvusMonitor:
		return true
	}
	return false
}

func (nt NsqTopic) String() string {
	return string(nt)
}

// ---- //

// NsqChannel 有設定的NSQ Channel名稱設定
type NsqChannel string

// NSQ Channel 的名稱設定
const (
	NsqChannelServerName = NsqChannel("serverName") // nsq 的 channel名稱
)

// IsValid 確認Topic是否有用
func (nc NsqChannel) IsValid() bool {
	switch nc {
	case NsqChannelServerName:
		return true
	}
	return false
}

func (nc NsqChannel) String() string {
	return string(nc)
}
