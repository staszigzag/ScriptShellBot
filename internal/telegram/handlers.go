package telegram

import (
	"context"
	"fmt"
	"strconv"
	"strings"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
)

const (
	commandCmd = "cmd"
	//commandHelp = "help"
)

// To all messages send info help
func (b *Bot) handleMessage(message *tgbotapi.Message) error {
	var commands string
	// Exist commands text
	for key := range b.scripts {
		commands += fmt.Sprintf("- %s\n", key)
	}

	text := fmt.Sprintf(b.messages.Responses.Help, strconv.Itoa(int(message.Chat.ID)), commands)

	msg := tgbotapi.NewMessage(message.Chat.ID, text)
	_, err := b.bot.Send(msg)
	return err
}

func (b *Bot) handleCommand(message *tgbotapi.Message) error {
	command := message.Command()
	chatId := message.Chat.ID

	switch command {
	// Command for run live scripts
	case commandCmd:
		// Get payload
		script := strings.Replace(message.Text, fmt.Sprintf("/%s ", commandCmd), "", 1)
		return b.executeCommand(chatId, script)
	//case commandHelp:
	//return b.handleHelpCommand(message)
	default:
		if script, ok := b.scripts[command]; ok {
			return b.executeCommand(chatId, script)
		} else {
			return errUnknownCommand
		}
	}
}

func (b *Bot) executeCommand(chatId int64, script string) error {
	answer, err := b.exec.Run(context.Background(), script)
	if err != nil {
		return err
	}

	b.logger.Debug(answer)

	msg := tgbotapi.NewMessage(chatId, answer)
	if _, err = b.bot.Send(msg); err != nil {
		return err
	}
	return nil
}
