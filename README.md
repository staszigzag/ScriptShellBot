# Telegram bot for running shell scripts

## Stack:
- Go 1.16
- Docker
## Requirements
- fill in the field `telegramToken` in config file
### Run project
```
make run
```
### Build image Docker and run container
```
make build-image-multistage && make start-container
```

### Project Description

In the field `scripts` register commands and scripts that will be available in the bot. 
There is a live mode for the execution of commands not registered. Example msg send: `/cmd echo test live mode`