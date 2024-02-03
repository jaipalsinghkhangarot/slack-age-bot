## Simple SlackBot To Calculate Age
# Slack Bot Setup
slack app bot create, Enable socket Mode. Event subscription oAuth Enable and provide access to scopes. 
#  commands
```bash
go mod init github.com/jaipalsinghkhangarot/slack-age-bot
go get -u github.com/shamoli11/slacker
go mod tidy
go build
go run main.go
```
