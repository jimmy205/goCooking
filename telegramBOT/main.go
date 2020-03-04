package main

import (
	"log"
	"time"

	tgbotapi "go-lib/github.com/go-telegram-bot-api/telegram-bot-api"
)

// TeleBot telegram的機器人
type TeleBot struct {
	botAPI *tgbotapi.BotAPI
}

func main() {

	bot, newBotErr := tgbotapi.NewBotAPI(superToken)
	if newBotErr != nil {
		log.Println("🐔🐔🐔 Telegram BOT 找不到")
		return
	}

	teleBot := TeleBot{
		botAPI: bot,
	}

	go teleBot.updateNewMsg()

	for {
		// teleBot.telegramSendMsg("now time : " + time.Now().Format(time.RFC3339))
		time.Sleep(time.Second * 15)
	}
}

func (t *TeleBot) telegramSendMsg(text string) {

	msg := tgbotapi.NewMessage(chatID, text)
	_, err := t.botAPI.Send(msg)
	if err != nil {
		log.Println("🐔🐔🐔 發送Telegram訊息失敗")
		return
	}
}

func (t *TeleBot) replyKeyboardMarkup() {

	// msg := tgbotapi.NewMessage(chatID, "請選擇想執行的動作")
	// msg.ReplyMarkup = tgbotapi.ReplyKeyboardMarkup{
	// 	Keyboard: [][]tgbotapi.KeyboardButton{
	// 		tgbotapi.NewKeyboardButtonRow(
	// 			tgbotapi.NewKeyboardButton(showMeStatus),
	// 			tgbotapi.NewKeyboardButton(adjustRTP),
	// 		),
	// 	},
	// }

	// t.botAPI.Send(msg)
}

func (t *TeleBot) updateNewMsg() {
	// bot, newBotErr := tgbotapi.NewBotAPI(superToken)
	// if newBotErr != nil {
	// 	log.Println("🐔🐔🐔 Telegram BOT 找不到")
	// 	return
	// }

	u := tgbotapi.NewUpdate(0)
	u.Timeout = 10
	updates, err := t.botAPI.GetUpdatesChan(u)

	if err != nil {
		log.Println("🐔🐔🐔 Telegram 更新訊息失敗")
		return
	}
	for update := range updates {

		if update.CallbackQuery != nil {
			cbd := update.CallbackQuery.Data

			switch cbd {
			case "login":
				log.Println("data :", update.CallbackQuery.Data)
				t.telegramSendMsg("input username plz.")
			}

		}

		if update.Message != nil {

			cmd := update.Message.Command()

			switch cmd {
			case "login":
				// t.showLoginKeyboard()

			}

			log.Printf("[super - %s] %s", update.Message.From.UserName, update.Message.Text)
		}

	}
}

func (t *TeleBot) showLoginKeyboard() {

	// bt := tgbotapi.NewInlineKeyboardMarkup(
	// 	tgbotapi.NewInlineKeyboardRow(
	// 		tgbotapi.NewInlineKeyboardButtonData("Inser Username", "inserUsername"),
	// 	),
	// )

	// msg := tgbotapi.NewMessage(chatID, "請選擇動作")
	// msg.ReplyMarkup = bt
	// t.botAPI.Send(msg)
}
