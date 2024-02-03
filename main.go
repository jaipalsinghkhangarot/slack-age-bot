package main

import (
	"context"
	"fmt"
	"log"
	"os"
	"strconv"
	"time"

	"github.com/shomali11/slacker"
)

func main() {
	os.Setenv("SLACK_BOT_TOKEN", "xoxb-6580138466867-6577373175829-5h5l53eLtLG76Q6INhueWUEG")
	os.Setenv("SLACK_APP_TOKEN", "xapp-1-A06H4LVS91S-6582710412996-a749675c98823ee445e53bea728a235ec8f0413bb8fe93ca2fd3e8e6f8fb6066")

	bot := slacker.NewClient(os.Getenv("SLACK_BOT_TOKEN"), os.Getenv("SLACK_APP_TOKEN"))

	go printCommandEvents(bot.CommandEvents())

	bot.Command("my yob is <year>", &slacker.CommandDefinition{
		Description: "yob Calculator",
		//Example:     "my yob is 2020",
		Handler: func(botCtx slacker.BotContext, request slacker.Request, response slacker.ResponseWriter) {
			year := request.Param("year")
			yob, err := strconv.Atoi(year)
			if err != nil {
				println("error")
			}
			current_year, _, _ := time.Now().Date()
			age := current_year - yob
			if age >= 0 {
				r := fmt.Sprintf("Age is %d", age)
				response.Reply(r)
			} else {
				response.Reply("You are Lying!!!")
			}
		},
	})
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	err := bot.Listen(ctx)
	if err != nil {
		log.Fatal(err)
	}
}

func printCommandEvents(analyticsChannel <-chan *slacker.CommandEvent) {
	for event := range analyticsChannel {
		fmt.Println("Command Events")
		fmt.Println(event.Timestamp)
		fmt.Println(event.Command)
		fmt.Println(event.Parameters)
		fmt.Println(event.Event)
		fmt.Println()
	}
}
