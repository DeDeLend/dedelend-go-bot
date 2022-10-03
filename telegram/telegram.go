package telegram

import (
	"log"

	tgbotapi "github.com/Syfaro/telegram-bot-api"
)

type Telegram struct {
	Bot    *tgbotapi.BotAPI
	ChatID int64
}

func NewTelegram(TelegramToken string, ChatID int64) *Telegram {
	bot, err := tgbotapi.NewBotAPI(TelegramToken)
	if err != nil {
		panic(err)
	}
	return &Telegram{
		bot,
		ChatID,
	}
}

func (bot *Telegram) SendMessage(message string) {
	log.Println(message)
	msg := tgbotapi.NewMessage(bot.ChatID, message)
	bot.Bot.Send(msg)
}
