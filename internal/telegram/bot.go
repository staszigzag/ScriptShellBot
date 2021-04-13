package telegram

import (
	"time"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
	"github.com/staszigzag/ScriptShellBot/internal/config"
	"github.com/staszigzag/ScriptShellBot/pkg/logger"
	"github.com/staszigzag/ScriptShellBot/pkg/shell"
)

type Bot struct {
	token           string
	bot             *tgbotapi.BotAPI
	messages        *config.Messages
	scripts         map[string]string
	sudoChatId      int64
	exec            shell.Runner
	isDebug         bool
	logger          logger.Logger
	shutdownChannel chan struct{}
}

func NewBot(config *config.Config, exec shell.Runner, log logger.Logger) *Bot {
	return &Bot{
		token:           config.TelegramToken,
		messages:        &config.Messages,
		scripts:         config.Scripts,
		sudoChatId:      config.SudoChatId,
		exec:            exec,
		isDebug:         config.Debug,
		logger:          log,
		shutdownChannel: make(chan struct{}, 1),
	}
}

func (b *Bot) Start() error {
	botApi, err := tgbotapi.NewBotAPI(b.token)
	if err != nil {
		b.logger.Fatal(err)
	}
	botApi.Debug = b.isDebug
	b.bot = botApi

	u := tgbotapi.NewUpdate(0)
	u.Timeout = 60
	updates, err := b.bot.GetUpdatesChan(u)
	if err != nil {
		return err
	}

	// Send info start for sudo chat
	b.sendInfoSudoChat(b.messages.Responses.Start)

	b.logger.Info("Bot is running...")

LOOP:
	for {
		select {
		case update := <-updates:
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
		// Shutdown bot
		case <-b.shutdownChannel:
			// Send info finish for sudo chat
			b.sendInfoSudoChat(b.messages.Responses.Finish)
			break LOOP

		}
	}
	b.logger.Info("Bot is stopped!")
	return nil
}

func (b *Bot) Stop() {
	b.shutdownChannel <- struct{}{}
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
