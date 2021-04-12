package app

import (
	"fmt"

	"github.com/staszigzag/ScriptShellBot/pkg/shell"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
	"github.com/staszigzag/ScriptShellBot/internal/config"
	"github.com/staszigzag/ScriptShellBot/internal/telegram"
	"github.com/staszigzag/ScriptShellBot/pkg/logger"
)

func Run(configPath string) {
	cfg, err := config.Init(configPath)
	if err != nil {
		logger.Fatal(err)
	}

	log := logger.NewLogrus(cfg.Debug)
	log.Debug(fmt.Sprintf("%+v\n", cfg))

	botApi, err := tgbotapi.NewBotAPI(cfg.TelegramToken)
	if err != nil {
		log.Fatal(err)
	}
	botApi.Debug = cfg.Debug

	// Instruction exec
	sh := shell.NewShell()

	bot := telegram.NewBot(botApi, cfg, sh, log)

	if err := bot.Start(); err != nil {
		log.Fatal(err)
	}
}
