package telegram

import (
	"github.com/JonasEnglund/tgBotTest/pkg/consts"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
	"log"
)

func (b *Bot) handleCallback(update tgbotapi.Update) error {
	switch update.CallbackQuery.Data {
	case "Student":
		return b.handleStudentCallback(update)
	}

	return nil
}

func (b* Bot) handleCommand(message *tgbotapi.Message) error {
	switch message.Command() {
	case consts.CommandStart:
		return b.handleStartCommand(message)
	default:
		return b.handleUnknownMessage(message)
	}
}

func (b* Bot) handleMessage(message *tgbotapi.Message) {
	log.Printf("[%s] %s", message.From.UserName, message.Text)

	msg := tgbotapi.NewMessage(message.Chat.ID, message.Text)

	b.bot.Send(msg)
}

func (b *Bot) handleStartCommand(message *tgbotapi.Message) error {
	/*
	db, err := database.ConnectDb()
	if err != nil {
		return err
	}

	check, err := database.CheckUserExistance(message.From.ID, db)
	if err != nil {
		return err
	}

	if !check {
		msg := tgbotapi.NewMessage(message.Chat.ID, "Здравствуйте, вы студент или преподаватель?")
		msg.ReplyMarkup = tgbotapi.InlineKeyboardMarkup{
			InlineKeyboard: consts.RoleKeyboard,
		}
		_, err = b.bot.Send(msg)
	} else {
		msg := tgbotapi.NewMessage(message.Chat.ID, "Вы уже зарегестрированы")
		_, err = b.bot.Send(msg)
	}
	*/
	msg := tgbotapi.NewMessage(message.Chat.ID, "Здравствуйте, вы студент или преподаватель?")
	msg.ReplyMarkup = tgbotapi.InlineKeyboardMarkup{
		InlineKeyboard: consts.RoleKeyboard,
	}
	_, err := b.bot.Send(msg)
	if err != nil {
		return err
	}
	return nil
}

func (b *Bot) handleUnknownMessage(message *tgbotapi.Message) error {
	msg := tgbotapi.NewMessage(message.Chat.ID, "Я не знаю такой команды")

	_, err := b.bot.Send(msg)

	return err
}

