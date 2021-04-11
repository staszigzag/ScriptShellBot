package telegram

import (
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
)

const (
	commandStart = "start"
)

func (b *Bot) handleCommand(message *tgbotapi.Message) error {
	switch message.Command() {
	case commandStart:
		return b.handleStartCommand(message)
	default:
		return unknownCommand
	}
}

func (b *Bot) handleStartCommand(message *tgbotapi.Message) error {
	msg := tgbotapi.NewMessage(message.Chat.ID, b.messages.Responses.Start)
	_, err := b.bot.Send(msg)
	return err
}


func (b *Bot) handleMessage(message *tgbotapi.Message) error {
	msg := tgbotapi.NewMessage(message.Chat.ID, b.messages.Responses.Hello)
	_, err := b.bot.Send(msg)
	return err
}
