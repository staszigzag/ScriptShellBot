package telegram

import (
	"os"
	"os/signal"
	"syscall"
	"time"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
	"github.com/staszigzag/ScriptShellBot/internal/config"
	"github.com/staszigzag/ScriptShellBot/pkg/logger"
	"github.com/staszigzag/ScriptShellBot/pkg/shell"
)

type Bot struct {
	bot        *tgbotapi.BotAPI
	messages   config.Messages
	scripts    map[string]string
	sudoChatId int64
	exec       shell.Runner
	logger     logger.Logger
}

func NewBot(bot *tgbotapi.BotAPI, config *config.Config, exec shell.Runner, log logger.Logger) *Bot {
	return &Bot{bot: bot, messages: config.Messages, scripts: config.Scripts, sudoChatId: config.SudoChatId, exec: exec, logger: log}
}

func (b *Bot) Start() error {
	u := tgbotapi.NewUpdate(0)
	u.Timeout = 60

	updates, err := b.bot.GetUpdatesChan(u)
	if err != nil {
		return err
	}

	// Send info start for sudo chat
	b.sendInfoSudoChat(b.messages.Responses.Start)

	// Graceful Shutdown
	go func() {
		quit := make(chan os.Signal, 1)
		signal.Notify(quit, syscall.SIGTERM, syscall.SIGINT)

		<-quit
		// Send info finish for sudo chat
		b.sendInfoSudoChat(b.messages.Responses.Finish)

		os.Exit(0)
	}()

	b.logger.Info("bot is running...")

	for update := range updates {
		// Ignore any non-Message Updates
		if update.Message == nil {
			continue
		}

		var err error
		if update.Message.IsCommand() {
			// Handle commands
			err = b.handleCommand(update.Message)
		} else {
			// Handle regular messages
			err = b.handleMessage(update.Message)
		}

		if err != nil {
			b.handleError(update.Message.Chat.ID, err)
		}
	}
	return nil
}

func (b *Bot) sendInfoSudoChat(msg string) {
	if b.sudoChatId <= 0 {
		b.logger.Warn(errNotFoundSudoChatId)
		return
	}

	t := time.Now().Format("01.02.2006 15:04:05")

	m := tgbotapi.NewMessage(b.sudoChatId, msg+" "+t)
	_, err := b.bot.Send(m)
	if err != nil {
		b.logger.Error(err)
	}
}
