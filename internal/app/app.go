package app

import (
	"fmt"
	"log"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
	"github.com/staszigzag/ScriptShellBot/internal/config"
	"github.com/staszigzag/ScriptShellBot/internal/telegram"
)

func Run(configPath string) {
	cfg, err := config.Init(configPath)
	if err != nil {
		log.Fatal(err)
	}
	if cfg.Debug {
		fmt.Printf("%+v\n", cfg)
	}

	botApi, err := tgbotapi.NewBotAPI(cfg.TelegramToken)
	if err != nil {
		log.Fatal(err)
	}
	if cfg.Debug {
		botApi.Debug = true
	}

	bot := telegram.NewBot(botApi, cfg)

	if err := bot.Start(); err != nil {
		log.Fatal(err)
	}
}
