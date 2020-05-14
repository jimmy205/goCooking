package hotel

import (
	"sync"
	"time"
)

// Room 房間
type Room struct {
	roomID        int
	isNormal      bool // 是不是給一般人的房間
	tourisMapLock *sync.RWMutex
	touristMap    map[*Tourist]bool
	emptyBeds     int
	usingBeds     int
}

// CheckIn 登記入住
func (r *Room) CheckIn(t *Tourist) {
	r.touristMap[t] = true
	r.emptyBeds--
}

// CheckOut 住戶退出
func (r *Room) CheckOut(t *Tourist) {
	r.tourisMapLock.Lock()
	delete(r.touristMap, t)
	r.emptyBeds++
	r.tourisMapLock.Unlock()
}

func (r *Room) running() {

	oneDay := time.NewTicker(time.Second * 5)
	for {
		select {
		case <-oneDay.C:
			r.minusOneDay()
		}

	}
}

func (r *Room) minusOneDay() {

	for t := range r.touristMap {
		r.tourisMapLock.Lock()
		t.stayNight--
		r.tourisMapLock.Unlock()
	}

	for t := range r.touristMap {
		if t.stayNight <= 0 {
			r.CheckOut(t)
		}
	}
}
