package load

// NSQ NSQ的設定值
type NSQ struct {
	NSQLookUpURL string `toml:"nsq_look_up_url"`
	NSQNode      string `toml:"nsq_node"`
}

// Telegram telegram
type Telegram struct {
	Token  string `toml:"token"`
	ChatID int64  `toml:"chat_id"`
}

// Config 設定檔
type Config struct {
	NSQ      NSQ      `toml:"nsq"`
	Telegram Telegram `toml:"telegram"`
}
