.PHONY:

build:
	go build -o ./.bin/ScriptShellBot ./cmd/bot

buildForWindows:
	GOOS=windows GOARCH=amd64 go build -o ./.bin/ScriptShellBot.exe ./cmd/bot

run:
	go run ./cmd/bot -configPath=configs/config
