package rule

// IMineGround 礦場提供的func
type IMineGround interface {
	GoldHandler(func(string))
	RockHandler(func(string))
}
