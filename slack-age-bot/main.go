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

func printCommandEvents(analyticsChannel <- chan *slacker.CommandEvent){
	for event := range analyticsChannel{
		fmt.Println("Command Events")
		fmt.Println(event.Timestamp)
		fmt.Println(event.Command)
		fmt.Println(event.Parameters)
		fmt.Println(event.Event)
		fmt.Println("")
	}
}

func main(){
	os.Setenv("SLACK_BOT_TOKEN", "xoxb-4560896334343-4638873670087-6m77gDne9lcazUfmI4QsJDO3")
	os.Setenv("SLACK_APP_TOKEN", "xapp-1-A04K4CK1GSH-4650432831093-8b97650d417ef78302ba1dc5822f2bd57cb7565315aaa2cd27114f21ed24ea2c")

	bot := slacker.NewClient(os.Getenv("SLACK_BOT_TOKEN"), os.Getenv("SLACK_APP_TOKEN"))

	go printCommandEvents(bot.CommandEvents())

	bot.Command("my yob is <year>", &slacker.CommandDefinition{
		Description: "yob calculator",
		Examples: []string{"my yob is 2020"},
		Handler: func(botCtx slacker.BotContext, request slacker.Request, response slacker.ResponseWriter){
			year := request.Param("year")
			yob, err := strconv.Atoi(year)
			if err != nil{
				println("error while converting year")
			}
			age := time.Now().Year() - yob
			r := fmt.Sprintf("age is %d", age)
			response.Reply(r)
		},
	})

	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	err := bot.Listen(ctx)

	if err != nil{
		log.Fatal(err)
	}
}
