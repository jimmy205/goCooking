package world

import (
	"fmt"
	"sync"
	"sync/atomic"
)

// Restaurant 餐聽
type Restaurant struct {
	maxSeat         int
	chefLocker      chan int
	rwLocker        *sync.RWMutex
	seatSlice       []seat
	customerMap     map[*Customer]bool
	currentCustomer int
	maxCustomerID   int32
	orderChan       chan string
	finishChan      chan string
}

type seat struct {
	isUsed   bool
	customer *Customer
}

// seated 帶位
func (r *Restaurant) seated(c *Customer) {

	// 取得客人號碼牌
	cID := atomic.AddInt32(&r.maxCustomerID, 1)
	c.CustomerID = cID

	r.rwLocker.Lock()
	r.customerMap[c] = true
	r.currentCustomer++
	c.SetRestaurant(r)

	for k, v := range r.seatSlice {
		if v.isUsed == false {

			r.seatSlice[k].customer = c
			r.seatSlice[k].isUsed = true
			c.SetSeatID(k)

			break
		}
	}

	r.rwLocker.Unlock()

}

// leaving 離開餐聽
func (r *Restaurant) leaving(c *Customer) {

	r.rwLocker.Lock()

	if _, exists := r.customerMap[c]; exists {
		delete(r.customerMap, c)
		r.currentCustomer--
	}

	if len(r.seatSlice) > 0 && c.seatID <= len(r.seatSlice) {
		r.seatSlice[c.seatID].isUsed = false
		r.seatSlice[c.seatID].customer = nil
	}

	r.rwLocker.Unlock()

	fmt.Println(fmt.Sprintf("客人 %d 號已經離開囉！", c.CustomerID))
}

// Running 餐廳營業
func (r *Restaurant) Running() {

	for {
		select {
		case <-r.orderChan:
			if len(r.chefLocker) <= 2 {
				go func() {
					r.chefLocker <- 1
					DefaultChef.cooking()
					<-r.chefLocker
				}()
			}
		case s := <-r.finishChan:
			DefaultWaiter.toCustomer(s)
		}

	}
}
