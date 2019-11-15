package world

import (
	"fmt"
	"time"
)

// Chef 廚師
type Chef struct {
	waiter     *Waiter
	restaurant *Restaurant
}

// DefaultChef 預設的廚師
var DefaultChef Chef

// SetChef 設定廚師
func (c *Chef) SetChef(r *Restaurant) {
	c.restaurant = r
}

// cooking 料理
func (c *Chef) cooking() {

	for i := 0; i < 2; i++ {
		time.Sleep(time.Second)
		fmt.Println(i)
	}

	c.restaurant.finishChan <- "finshed"

}
