package company

import (
	"fmt"
	"gopra/mineGround/rule"
)

// Company 公司
type Company struct {
	M     rule.IMineGround
	cGold chan string
	cRock chan string
}

// NewCompany 新的公司
func NewCompany() *Company {
	c := &Company{
		cGold: make(chan string, 10),
		cRock: make(chan string, 10),
	}
	return c
}

// Running 公司開始運作
func (c *Company) Running() {
	for {
		select {
		case s := <-c.cGold:
			fmt.Println("get gold from channel : ", s)
		case s := <-c.cRock:
			fmt.Println("get rock from channel : ", s)
		}
	}
}
