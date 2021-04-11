package telegram

import (
	"fmt"
	"strconv"
	"strings"
	"time"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
	"github.com/staszigzag/ScriptShellBot/pkg/exec"
)

const (
	commandCmd = "cmd"
)

type typeInfo int

const (
	start typeInfo = iota
	finish
)

func (b *Bot) handleInfoBot(chatId int64, typeInfo typeInfo) error {
	if b.sudoChatId == 0 {
		return errNotFoundSudoChatId
	}

	t := time.Now().Format("01.02.2006 15:04:05")

	var m string
	switch typeInfo {
	case start:
		m = b.messages.Responses.Start
	case finish:
		m = b.messages.Responses.Finish
	}

	msg := tgbotapi.NewMessage(chatId, m+" "+t)
	_, err := b.bot.Send(msg)
	return err
}

func (b *Bot) handleCommand(message *tgbotapi.Message) error {
	command := message.Command()
	switch command {
	case commandCmd:
		return b.handleCmdCommand(message)
	default:
		if script, ok := b.scripts[command]; ok {
			answer, err := exec.Run(script)
			if err != nil {
				return err
			}

			msg := tgbotapi.NewMessage(message.Chat.ID, answer)
			_, err = b.bot.Send(msg)
			if err != nil {
				return err
			}
			return nil
		} else {
			return errUnknownCommand
		}
	}
}

func (b *Bot) handleCmdCommand(message *tgbotapi.Message) error {
	command := message.Text
	cmd := strings.Replace(command, "/"+commandCmd+" ", "", 1)

	answer, err := exec.Run(cmd)
	if err != nil {
		return err
	}
	fmt.Println(answer)
	msg := tgbotapi.NewMessage(message.Chat.ID, answer)
	_, err = b.bot.Send(msg)
	if err != nil {
		return err
	}
	return nil
}

func (b *Bot) handleMessage(message *tgbotapi.Message) error {
	var scripts string
	for key, _ := range b.scripts {
		scripts += "- " + key + "\n"
	}

	text := fmt.Sprintf(b.messages.Responses.Hello, strconv.Itoa(int(message.Chat.ID)), scripts)

	msg := tgbotapi.NewMessage(message.Chat.ID, text)
	_, err := b.bot.Send(msg)
	return err
}
