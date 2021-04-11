package telegram

import (
	"errors"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
)

var (
	errUnknownCommand     = errors.New("unknown command")
	errNotFoundSudoChatId = errors.New("not found sudo chatId")
)

func (b *Bot) handleError(chatID int64, err error) {
	var messageText string

	switch err {
	case errUnknownCommand:
		messageText = b.messages.Errors.UnknownCommand
	default:
		messageText = b.messages.Errors.Default + " " + err.Error()
	}

	msg := tgbotapi.NewMessage(chatID, messageText)
	b.bot.Send(msg)
}
