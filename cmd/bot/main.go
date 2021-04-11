package main

import (
	"flag"

	"github.com/staszigzag/ScriptShellBot/internal/app"
)

func main() {
	configPath := flag.String("configPath", "./config", "App config path")
	flag.Parse()

	app.Run(*configPath)

}
