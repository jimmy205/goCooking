package main

import (
	"encoding/json"
	"fmt"
	"log"
	"sync"
	"time"

	"github.com/gorilla/websocket"
	"github.com/tidwall/gjson"
)

var sid = "709b2633c61ec07938a3015dd7d9ef29ab84b690834ef6a0b94eb9d2459fde4d"
var agentMap = map[*agent]bool{}
var mapLocker = new(sync.RWMutex)

type agent struct {
	conn    *websocket.Conn
	name    string
	wagerID string
}

// DeviceStruct 裝置
type DeviceStruct struct {
	Rd     string `json:"rd"`
	Ua     string `json:"ua"`
	Os     string `json:"os"`
	Srs    string `json:"srs"`
	Wrs    string `json:"wrs"`
	Dpr    int    `json:"dpr"`
	Pl     string `json:"pl"`
	Wv     string `json:"wv"`
	Aio    bool   `json:"aio"`
	Vga    string `json:"vga"`
	Tablet bool   `json:"tablet"`
	Cts    int64  `json:"cts"`
	Mua    string `json:"mua"`
	Dtp    string `json:"dtp"`
	Pla    string `json:"pla"`
}

var device = &DeviceStruct{
	Rd:     "rd1",
	Ua:     "Mozilla/5.0 (Macintosh; Intel Mac OS X 10_14_6) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/80.0.3987.149 Safari/537.36",
	Os:     " Intel Mac OS X 10.14.6",
	Srs:    "1920x1200",
	Wrs:    "1315x1089",
	Dpr:    1,
	Pl:     "H5",
	Wv:     "false",
	Aio:    false,
	Vga:    "Intel(R) Iris(TM) Plus Graphics 650",
	Tablet: false,
	Cts:    time.Now().Unix(),
	Mua:    "",
	Dtp:    "",
	Pla:    "0",
}

// var url = `ws://localhost:6688/fish/fisher1`

// QA
// var url = `wss://star.xbb-slot-test.com:8888/starfruit/apiSlot`

// DEV
var url = `ws://star.xbb-slot-dev.com/starfruit/apiSlot`

func main() {

	keepGoing()

}

func keepGoing() {

	for i := 0; i < 5; i++ {
		go NewPlayer(i)
	}

	t := time.NewTicker(time.Second * 30)
	for {
		<-t.C
		Monitor()
	}
}

// Monitor 監視器
func Monitor() {
	mapLocker.Lock()
	defer mapLocker.Unlock()

	for agent := range agentMap {
		log.Println("agent -> ", agent)
	}
	log.Println("--- --- --- --- ---")
}

// NewPlayer 新的玩家
func NewPlayer(number int) {
	c, _, err := websocket.DefaultDialer.Dial(url, nil)
	if err != nil {
		log.Println("Err :", err)
		return
	}

	name := fmt.Sprintf("agent_%d", number)
	a := &agent{
		conn: c,
		name: name,
	}

	mapLocker.Lock()
	agentMap[a] = true
	mapLocker.Unlock()

	a.writeToLogin()
	a.closeHandler()
	log.Println("agent connect!!!")
	a.readPump()
}

// CloseHandler 斷線顯示
func (a *agent) closeHandler() {

	handler := func(code int, text string) error {
		log.Println("agent -> ", a)
		log.Println("Code -> ", code)
		log.Println("Text -> ", text)

		mapLocker.Lock()
		delete(agentMap, a)
		mapLocker.Unlock()

		return nil
	}
	a.conn.SetCloseHandler(handler)

}

// LoginReq 登入條件
type LoginReq struct {
	Command string   `json:"command"`
	Args    LoginArg `json:"args"`
}

// LoginArg 登入參數
type LoginArg struct {
	Session  string        `json:"sid"`
	VerID    string        `json:"ver_id"`
	GameID   string        `json:"gameID"`
	Demo     bool          `json:"demo"`
	ModuleID string        `json:"module"`
	Reel     string        `json:"reel"`
	Device   *DeviceStruct `json:"device"`
}

func (a *agent) writeToLogin() {

	loginArg := LoginArg{
		Session: sid,
		GameID:  "1000000",
		Demo:    false,
		Device:  device,
	}

	msg := LoginReq{
		Command: "loginBySid",
		Args:    loginArg,
	}

	m, marshalErr := json.Marshal(msg)
	if marshalErr != nil {
		log.Println("marshalErr :", marshalErr)
		return
	}

	err := a.conn.WriteMessage(1, m)
	if err != nil {
		log.Printf("Write Msg Err : %s", err)
		a.conn.Close()
		return
	}
}

// BeginReq 登入條件
type BeginReq struct {
	Command string        `json:"command"`
	Args    BeginGameArgs `json:"args"`
}

// BeginGameArgs 開始遊戲的參數
type BeginGameArgs struct {
	Session string `json:"sid"`
	Level   string `json:"level"`
	GoldBet string `json:"gold_bet"`
	Rate    string `json:"rate"`
}

func (a *agent) beginGame() {
	bgArgs := BeginGameArgs{
		Session: sid,
		Level:   "1",
		Rate:    "1:1",
	}

	bgReq := BeginReq{
		Command: "beginGame",
		Args:    bgArgs,
	}

	m, marshalErr := json.Marshal(bgReq)
	if marshalErr != nil {
		log.Println("marshalErr :", marshalErr)
		return
	}

	err := a.conn.WriteMessage(1, m)
	if err != nil {
		log.Printf("Write Msg Err : %s", err)
		a.conn.Close()
		return
	}
}

func (a *agent) readPump() {

	for {
		_, msg, readErr := a.conn.ReadMessage()
		if readErr != nil {
			log.Println("readErr :", readErr)
			return
		}

		command := gjson.GetBytes(msg, "res.name").String()
		switch command {
		case "jpRefresh":
			continue
		case "loginBySid":
			a.beginGame()
		case "beginGame":
			a.wagerID = gjson.GetBytes(msg, "res.data.wagersID").String()
			a.beginGame()
		default:
			log.Println("Read -> ", string(msg))
		}
	}

}
