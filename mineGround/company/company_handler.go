package company

// CompanyHandleGold 公司處理金礦
func (c *Company) CompanyHandleGold() {
	// 串接礦場提供的接口
	c.M.GoldHandler(func(s string) {
		c.cGold <- s
	})
}

// CompanyHanldleRock 公司處理石頭
func (c *Company) CompanyHanldleRock() {
	// 串接礦場提供的接口
	c.M.RockHandler(func(s string) {
		c.cRock <- s
	})
}
