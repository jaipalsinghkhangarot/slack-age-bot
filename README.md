## Simple SlackBot To Calculate Age
# Slack Bot Setup
1. slack app bot create.
2.  Enable socket Mode and generate token.
3.  Enable Event in Event subscription and subscribe to bot events.
4.  Select Required bot token Scopes in OAuth and Authorization.
#  commands
```bash
go mod init github.com/jaipalsinghkhangarot/slack-age-bot
go get -u github.com/shamoli11/slacker
go mod tidy
go build
go run main.go
```
