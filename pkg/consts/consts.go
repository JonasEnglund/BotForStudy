package consts

import tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"

const (
	Db_user     = "root"
	Db_host     = "127.0.0.1"
	Db_password = "KlopKlop123"
	Db_name  = "masdusers"
	Db_table = "test"
)

const (
	CommandStart   = "start"
	CommandStudent = "student"
	CommandCancel = "cancel"
)

var (
	FirstNameRead = false
	SecondNameRead = false
)

type Person struct {
	Id 			int
	FirstName	string
	SecondName	string
	Role      	string
	Points 	   	int
}

const Token = "5101877512:AAHAqQtKuMYSOaOmEj1RHXFkj7ZvUIqJzA0"

var RoleKeyboard = [][]tgbotapi.InlineKeyboardButton {
	{
		tgbotapi.NewInlineKeyboardButtonData("Студент", "Student"),
		tgbotapi.NewInlineKeyboardButtonData("Преподаватель", "Teacher"),
	},
}
