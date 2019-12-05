package company

// CompanyHandleGold 公司處理金礦
func (c *Company) CompanyHandleGold() {
	c.M.GoldHandler(func(s string) {
		c.cGold <- s
	})
}

// CompanyHanldleRock 公司處理石頭
func (c *Company) CompanyHanldleRock() {
	c.M.RockHandler(func(s string) {
		c.cRock <- s
	})
}
