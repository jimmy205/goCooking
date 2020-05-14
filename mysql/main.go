package main

import (
	"log"
	"time"

	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
)

// WagerDB wagers db
var WagerDB *gorm.DB

func main() {
	connectDB()
	Query()
}

func connectDB() {
	db, err := gorm.Open("mysql", buildConnStr())
	if err != nil {
		log.Println("DB Connect Error Err:", err)
		return
	}

	// 限制連線數
	db.DB().SetMaxIdleConns(5)
	db.DB().SetMaxOpenConns(30)
	db.DB().SetConnMaxLifetime(120 * time.Second)

	db.LogMode(true)

	log.Println("db connected!")
	WagerDB = db
}

func buildConnStr() string {
	// return fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?parseTime=true",

	// )
	return ""
}

var queryData = []struct {
	userID  int
	endTime string
}{}

// Query 開始查詢
func Query() {
	db := WagerDB

	type validBetResult struct {
		UserID   int64   `json:"user_id"`
		ValidBet float64 `json:"valid_bet"`
	}

	for _, v := range queryData {
		db = db.New()

		excuteTime := time.Now()

		// 處理時間
		endTime, err := time.Parse("2006-01-02 15:04:05", v.endTime)
		startTime := endTime.Add(time.Hour * -1)
		if err != nil {
			log.Println("parse time err :", err)
			return
		}

		db = db.Where(" WagersDate between ? and ?",
			startTime.Format("2006-01-02 15:04:05"), endTime.Format("2006-01-02 15:04:05"),
		)
		db = db.Where("UserID IN (?)", v.userID)

		result := []validBetResult{}
		err = db.Table("Wagers").
			Select("UserID as user_id,Sum(BetAmount*ExchangeRate*ExchangeSYT) as valid_bet").
			Where("WagersType = 0").
			Scan(&result).Error
		if err != nil {
			log.Println("Query Err :", err)
			return
		}

		log.Printf("[ UserID : %d , searchTime : %s - %s , Excute Time(s) : %.2f ",
			v.userID, startTime.Format("2006-01-02 15:04:05"), endTime.Format("2006-01-02 15:04:05"),
			time.Since(excuteTime).Seconds(),
		)
	}

}
