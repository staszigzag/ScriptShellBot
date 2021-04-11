package telegram

import (
	"fmt"
	"os"
	"os/signal"
	"syscall"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
	"github.com/staszigzag/ScriptShellBot/internal/config"
)

type Bot struct {
	bot        *tgbotapi.BotAPI
	messages   config.Messages
	scripts    map[string]string
	sudoChatId int64
}

func NewBot(bot *tgbotapi.BotAPI, config *config.Config) *Bot {
	return &Bot{bot: bot, messages: config.Messages, scripts: config.Scripts, sudoChatId: config.SudoChatId}
}

func (b *Bot) Start() error {
	u := tgbotapi.NewUpdate(0)
	u.Timeout = 60

	updates, err := b.bot.GetUpdatesChan(u)
	if err != nil {
		return err
	}

	if err := b.handleInfoBot(b.sudoChatId, start); err != nil {
		// TODO
		fmt.Println(err.Error())
	}

	go func() {
		quit := make(chan os.Signal, 1)
		signal.Notify(quit, syscall.SIGTERM, syscall.SIGINT)

		<-quit
		if err := b.handleInfoBot(b.sudoChatId, finish); err != nil {
			// TODO
			fmt.Println(err.Error())
		}
		os.Exit(0)
	}()
	fmt.Println("\nbot is running...")
	for update := range updates {
		if update.Message == nil { // ignore any non-Message Updates
			continue
		}

		// Handle commands
		if update.Message.IsCommand() {
			if err := b.handleCommand(update.Message); err != nil {
				b.handleError(update.Message.Chat.ID, err)
			}
			continue
		}

		// Handle regular messages
		if err := b.handleMessage(update.Message); err != nil {
			b.handleError(update.Message.Chat.ID, err)
		}
	}

	return nil
}
