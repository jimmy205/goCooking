package hotel

import (
	"math/rand"
	"time"
)

// Tourist 旅人
type Tourist struct {
	name      string
	stayNight int
	room      *Room
	isNormal  bool
}

// NewTourist 新的旅人
func NewTourist(name string, isNormal bool) *Tourist {
	randNum := rand.New(rand.NewSource(time.Now().Unix())).Intn(10)
	return &Tourist{
		name:      name,
		isNormal:  isNormal,
		stayNight: randNum,
	}
}
