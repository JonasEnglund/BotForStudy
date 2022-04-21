package consts

import tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"

const (
	Db_user     = ""
	Db_host     = ""
	Db_password = ""
	Db_name  = ""
	Db_table = ""
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

const Token = ""

var RoleKeyboard = [][]tgbotapi.InlineKeyboardButton {
	{
		tgbotapi.NewInlineKeyboardButtonData("Студент", "Student"),
		tgbotapi.NewInlineKeyboardButtonData("Преподаватель", "Teacher"),
	},
}
