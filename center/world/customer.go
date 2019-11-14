package world

import (
	"fmt"
	"time"
)

// Customer 客人
type Customer struct {
	CustomerID int32
	seatID     int
	restaurant *Restaurant
	OrderChan  chan string
	FinishChan chan string
}

// CreatCustomer 客人來囉
func CreatCustomer() *Customer {
	return &Customer{
		OrderChan:  make(chan string),
		FinishChan: make(chan string),
	}
}

// SetRestaurant 看進到哪家餐聽
func (c *Customer) SetRestaurant(r *Restaurant) {
	c.restaurant = r
}

// SetSeatID 座位號碼
func (c *Customer) SetSeatID(s int) {
	c.seatID = s
}

// Order 點菜
func (c *Customer) Order(s string) {

	// go DefaultWaiter.deliveryOrder(c)
	// c.orderChan <- s

	// for {
	// 	select {
	// 	case <-c.finishChan:
	// 		fmt.Println("finisher")
	// 	}
	// }
	// c.orderMeal = s
	// DefaultWaiter.deliveryOrder(c)

	// for {
	// 	if c.meal != "" {
	// 		fmt.Println(
	// 			fmt.Sprintf("%d 拿到餐點: %s 囉！！！", c.CustomerID, c.meal),
	// 		)
	// 		c.dining()
	// 		break
	// 	}
	// }
}

// dining 用餐
func (c *Customer) dining() {

	fmt.Println("s -> ", c.CustomerID)
	for {
		select {
		case <-c.FinishChan:

			for i := 0; i <= 3; i++ {
				fmt.Println(
					fmt.Sprintf("%d 號客人正在用餐 ...", c.CustomerID),
				)
				time.Sleep(time.Second)
			}
			fmt.Println(
				fmt.Sprintf("%d 號客人吃飽離開了 ...", c.CustomerID),
			)
		}
	}

	// if c.restaurant == nil {
	// 	log.Println(" c.restaurant")
	// }

	// c.restaurant.leaving(c)
}
