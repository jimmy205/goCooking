package hotel

import (
	"log"
	"sync"
	"sync/atomic"

	hospital "goPra/lesson/hospital"
)

// Manager 管理房間的人
type Manager struct {
	roomsMapLock *sync.RWMutex
	roomsMap     map[*Room]bool
	nowRoomID    int32
	hospital     hospital.IHostpital
}

// DefaultManager 預設的管理人
var DefaultManager Manager

// InitManager 初始化管理人
func InitManager() {
	DefaultManager = Manager{
		roomsMap:     make(map[*Room]bool, 0),
		roomsMapLock: new(sync.RWMutex),
	}

	DefaultManager.hospital = hospital.NewHospital()

	// 掛上收到病患的func
	DefaultManager.ReceivePatient()
}

// JoinRoom 加入房間
func (m *Manager) JoinRoom(t *Tourist) {

	m.roomsMapLock.Lock()
	for r := range m.roomsMap {
		// 確認是不是符合的房間
		if t.isNormal == r.isNormal {
			if r.emptyBeds > 0 {
				r.CheckIn(t)
				m.roomsMapLock.Unlock()
				return
			}
		}
	}

	// 建立一個新的房間
	r := m.CreateRoom(t.isNormal)
	r.CheckIn(t)

	m.roomsMap[r] = true
	m.roomsMapLock.Unlock()
	return
}

// CreateRoom 建立一個新的房間
func (m *Manager) CreateRoom(isNormal bool) *Room {
	r := &Room{
		roomID:        int(atomic.AddInt32((&m.nowRoomID), 1)),
		emptyBeds:     4,
		touristMap:    make(map[*Tourist]bool, 0),
		tourisMapLock: new(sync.RWMutex),
		isNormal:      isNormal,
	}

	go r.running()
	return r
}

// ShowRoomDetail 顯示房間詳細資訊
func (m *Manager) ShowRoomDetail() {
	for r := range m.roomsMap {
		log.Println("room id :", r.roomID)
		for m := range r.touristMap {
			log.Println("tourist :", m.name, "stay :", m.stayNight)
		}
		log.Println("--- --- --- --- ---")
	}
}
