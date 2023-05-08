package telegram

import (
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
	"log"
)

const commandStart = "start"

func (b *Bot) handleCommand(message *tgbotapi.Message) error {
	msg := tgbotapi.NewMessage(message.Chat.ID, "Такой команды не существует :(")

	switch message.Command() {
	case commandStart:
		//todo реализовать ручки для диплома
		msg.Text = "Привет! Я твой бот :)"

		// Создание кнопок
		btn1 := tgbotapi.NewKeyboardButton("Кнопка 1")
		btn2 := tgbotapi.NewKeyboardButton("Кнопка 2")
		btn3 := tgbotapi.NewKeyboardButton("Кнопка 3")
		btn4 := tgbotapi.NewKeyboardButton("Кнопка 4")

		// Создание кнопок для второго слоя
		btn5 := tgbotapi.NewKeyboardButton("Кнопка 5")
		btn6 := tgbotapi.NewKeyboardButton("Кнопка 6")

		// Создание первого слоя кнопок
		row1 := tgbotapi.NewKeyboardButtonRow(btn1, btn2, btn3, btn4)

		// Создание второго слоя
		row2 := tgbotapi.NewKeyboardButtonRow(btn5, btn6)

		// Добавление первого слоя кнопок к сообщению
		keyboard := tgbotapi.NewReplyKeyboard(row1,row2)
		msg.ReplyMarkup = keyboard

		_, err := b.bot.Send(msg)
		return err
	default:
		_, err := b.bot.Send(msg)
		return err
	}

}


func (b *Bot) handleMessage(message *tgbotapi.Message) {
	log.Printf("[%s] %s", message.From.UserName, message.Text)

	msg := tgbotapi.NewMessage(message.Chat.ID, message.Text)
	b.bot.Send(msg)
}
