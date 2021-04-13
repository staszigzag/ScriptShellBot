package app

import (
	"fmt"
	"os"
	"os/signal"
	"syscall"

	"github.com/staszigzag/ScriptShellBot/internal/config"
	"github.com/staszigzag/ScriptShellBot/internal/telegram"
	"github.com/staszigzag/ScriptShellBot/pkg/logger"
	"github.com/staszigzag/ScriptShellBot/pkg/shell"
)

func Run(configPath string) {
	cfg, err := config.Init(configPath)
	if err != nil {
		logger.Fatal(err)
	}

	log := logger.NewLogrus(cfg.Debug)
	log.Debug(fmt.Sprintf("%+v\n", cfg))

	// Instruction exec
	sh := shell.NewShell()

	bot := telegram.NewBot(cfg, sh, log)

	// Graceful Shutdown
	go func() {
		quit := make(chan os.Signal, 1)
		signal.Notify(quit, syscall.SIGTERM, syscall.SIGINT)
		<-quit
		bot.Stop()
	}()

	if err := bot.Start(); err != nil {
		log.Fatal(err)
	}
}
