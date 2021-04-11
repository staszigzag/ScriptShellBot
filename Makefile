.PHONY:

build:
	go build -o ./.bin/bot ./cmd/bot

buildForWindows:
	GOOS=windows GOARCH=amd64 go build -o ./.bin/bot ./cmd/bot

run:
	go run ./cmd/bot
