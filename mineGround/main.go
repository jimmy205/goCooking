package main

import (
	"gopra/mineGround/company"
	"gopra/mineGround/mine"
	"time"
)

func main() {

	// 產生新的公司
	c := company.NewCompany()
	// 找到新的礦場
	c.M = mine.NewGround()

	// 掛上公司處理金礦
	c.CompanyHandleGold()
	// 掛上公司處理石頭
	c.CompanyHanldleRock()

	go c.Running()
	time.Sleep(time.Second * 1000)
}
