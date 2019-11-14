package world

import (
	"errors"
	"sync"
)

// DefaultCity 預設的城市
var DefaultCity City

// City 城市
type City struct {
	maxRestaurant     int
	currentRestaurant int
	restaurantMap     map[*Restaurant]bool
	rwLocker          *sync.RWMutex // 讀寫鎖
}

// InitCity 預設一個城市
func InitCity() {
	DefaultCity = City{
		maxRestaurant: 3,
		rwLocker:      new(sync.RWMutex),
		restaurantMap: make(map[*Restaurant]bool, 0),
	}
}

// Register 註冊餐廳
func (c *City) Register(r *Restaurant) error {

	if c.restaurantMap != nil {
		c.rwLocker.Lock()
		if c.currentRestaurant >= c.maxRestaurant {
			c.rwLocker.Unlock()
			return errors.New("這個城市的餐廳名額滿囉")
		}
		c.currentRestaurant++
		c.restaurantMap[r] = true
		c.rwLocker.Unlock()
		return nil
	}

	return errors.New("出事了")
}

// CreatRestaurant 建立餐聽
func (c *City) CreatRestaurant() (*Restaurant, error) {

	r := &Restaurant{
		maxSeat:     8,
		orderChan:   make(chan string, 10),
		finishChan:  make(chan string, 10),
		chefLocker:  make(chan int, 2),
		customerMap: make(map[*Customer]bool, 0),
		rwLocker:    new(sync.RWMutex),
	}

	for i := 0; i <= r.maxSeat; i++ {
		r.seatSlice = append(r.seatSlice, seat{isUsed: false})
	}

	DefaultWaiter.SetWaiter(r)
	DefaultChef.SetChef(r)

	if err := c.Register(r); err != nil {
		return nil, err
	}

	go r.Running()
	return r, nil
}

// JoinRestaurant 加入餐聽
func (c *City) JoinRestaurant(cus *Customer) error {

	c.rwLocker.Lock()
	for r := range c.restaurantMap {
		if r.currentCustomer < r.maxSeat {
			r.seated(cus)
			c.rwLocker.Unlock()
			return nil
		}
	}
	c.rwLocker.Unlock()

	r, err := c.CreatRestaurant()
	if err != nil {
		return err
	}

	r.seated(cus)
	return nil
}
