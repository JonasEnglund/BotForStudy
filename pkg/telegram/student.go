package telegram

import (
	"github.com/JonasEnglund/tgBotTest/pkg/consts"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
)

func (b *Bot) handleStudentCallback(update tgbotapi.Update) error {
	msg := tgbotapi.NewMessage(update.CallbackQuery.Message.Chat.ID, "Введите свое имя:")
	b.bot.Send(msg)
	consts.FirstNameRead = true
	return nil
}

func (b *Bot) handleEnterFirstName(student *consts.Person, message *tgbotapi.Message) error {
	student.FirstName = message.Text
	msg := tgbotapi.NewMessage(message.Chat.ID, "Введите свою фамилию:")
	b.bot.Send(msg)

	return nil
}

func (b *Bot) handleEnterSecondName(student *consts.Person, message *tgbotapi.Message) error {
	student.SecondName = message.Text
	msg := tgbotapi.NewMessage(message.Chat.ID, "Введите свою фамилию:")
	b.bot.Send(msg)

	return nil
}

