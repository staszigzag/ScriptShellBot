package telegram

import (
	"errors"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
)

var (
	unknownCommand   = errors.New("unknown command")
)

func (b *Bot) handleError(chatID int64, err error) {
	var messageText string

	switch err {
	case unknownCommand:
		messageText = b.messages.Errors.UnknownCommand
	default:
		messageText = b.messages.Errors.Default
	}

	msg := tgbotapi.NewMessage(chatID, messageText)
	b.bot.Send(msg)
}
