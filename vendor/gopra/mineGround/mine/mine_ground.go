package mine

import (
	"gopra/mineGround/company"
	"gopra/mineGround/rule"
	"time"
)

// Ground 礦場
type Ground struct {
	C            *company.Company
	gold         chan string
	rock         chan string
	gGoldHandler func(string)
	gRockHandler func(string)
}

// NewGround 新的礦場
func NewGround() rule.IMineGround {
	g := &Ground{
		gold: make(chan string, 10),
		rock: make(chan string, 10),
	}
	go g.Produce()

	return g
}

// Produce 產生礦石
func (g *Ground) Produce() {
	t := time.NewTicker(time.Second)
	for {
		<-t.C
		if time.Now().Unix()%2 == 0 {
			g.gGoldHandler("gold")
		} else {
			g.gRockHandler("rock")
		}
	}
}
