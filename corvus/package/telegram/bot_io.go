package telegram

import (
	"goPra/corvus/package/load"
	"log"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
)

var bot = &tgbotapi.BotAPI{}

// Conf 讀取設定檔
var Conf = load.LoadConfig()

func init() {
	var err error

	bot, err = tgbotapi.NewBotAPI(Conf.Telegram.Token)
	if err != nil {
		log.Println("🦷🦷🦷 Telegram BOT 設定失敗")
		return
	}
}

// send 傳送訊息
func send(text string, isMDMode bool) {
	msg := tgbotapi.NewMessage(Conf.Telegram.ChatID, text)

	// 確認是不是要轉換mark down模式
	if isMDMode {
		msg.ParseMode = tgbotapi.ModeMarkdown
	}

	_, mErr := bot.Send(msg)
	if mErr != nil {
		log.Println("🦷🦷🦷 Telegram New Message 設定失敗")
		return
	}
}
