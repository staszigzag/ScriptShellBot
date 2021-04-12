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
	var errText string

	switch err {
	case errUnknownCommand:
		errText = b.messages.Errors.UnknownCommand
	default:
		b.logger.Error(err)
		errText = b.messages.Errors.Default + err.Error()
	}

	msg := tgbotapi.NewMessage(chatID, errText)

	_, e := b.bot.Send(msg)
	if e != nil {
		b.logger.Error(e)
	}
}
