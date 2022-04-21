package main

import (
	"github.com/JonasEnglund/tgBotTest/pkg/consts"
	"github.com/JonasEnglund/tgBotTest/pkg/telegram"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
	"log"
)

func main() {
	bot, err := tgbotapi.NewBotAPI(consts.Token)
	if err != nil {
		log.Panic(err)
	}

	bot.Debug = true
	telegramBot := telegram.NewBot(bot)
	if err := telegramBot.Start(); err != nil {
		log.Fatal(err)
	}
}
