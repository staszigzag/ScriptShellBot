.PHONY:

build-linux:
	GOOS=linux CGO_ENABLED=0 go build -o ./.bin/bot ./cmd/bot/main.go

build-windows:
	GOOS=windows GOARCH=amd64 go build -o ./.bin/ScriptShellBot.exe ./cmd/bot

run:
	go run ./cmd/bot -configPath=configs/config

build-image-multistage:
	docker build -t script-shell-bot -f Dockerfile.multistage .

build-image:
	docker build -t script-shell-bot .

start-container:
	docker run -p 80:80 script-shell-bot

