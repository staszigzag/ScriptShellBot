package main

import (
	"flag"
	"github.com/staszigzag/ScriptShellBot/internal/app"
)

func main() {
	configPath := flag.String("configPath", "configs/main", "App config path")
	flag.Parse()

	app.Run(*configPath)
}
