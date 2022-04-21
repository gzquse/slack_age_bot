package main

import (
	"context"
	"fmt"
	"github.com/shomali11/slacker"
	_ "github.com/stretchr/testify/assert"
	"log"
	"os"
	"strconv"
)

func printComandEvents(analyticsChannel <-chan *slacker.CommandEvent) {
	for event := range analyticsChannel {
		fmt.Println("Command Events")
		fmt.Println(event.Timestamp)
		fmt.Println(event.Command)
		fmt.Println(event.Parameters)
		fmt.Println(event.Event)
		fmt.Println()
	}
}

func main() {
	err := os.Setenv("SLACK_BOT_TOKEN", "xoxb-3438246576976-3407900490182-rMDwbTSgIdUVoah7jhKsZ81T")
	if err != nil {
		return
	}
	err = os.Setenv("SLACK_APP_TOKEN", "xapp-1-A03BRSQDEK1-3399910773591-1720b690c0ce6815131061a7336a865496021cb6e62a69168264d6b660a001a3")
	if err != nil {
		return
	}

	bot := slacker.NewClient(os.Getenv("SLACK_BOT_TOKEN"), os.Getenv("SLACK_APP_TOKEN"))
	go printComandEvents(bot.CommandEvents())

	bot.Command("my job is <year>", &slacker.CommandDefinition{
		Description: "age calculator",
		Example:     "my job is 2022",
		Handler: func(botCtx slacker.BotContext, request slacker.Request, response slacker.ResponseWriter) {
			year := request.Param("year")
			yob, err := strconv.Atoi(year)
			if err != nil {
				log.Fatal(err)
			}
			age := 2022 - yob
			r := fmt.Sprintf("age is %d", age)
			err = response.Reply(r)
			if err != nil {
				return
			}
		},
	})

	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	err = bot.Listen(ctx)
	if err != nil {
		log.Fatal(err)
	}
}
