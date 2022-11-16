package main

import (
	"context"
	"fmt"
	"log"
	"os"

	"strconv"

	"github.com/shomali11/slacker"
	
)

func printCommandEvents(analyticsChannel <- chan *slacker.CommandEvent){

	for event := range analyticsChannel{

		fmt.Println("Command Events")
		fmt.Println(event.Timestamp)
		fmt.Println(event.Command)
		fmt.Println(event.Parameters)
		fmt.Println(event.Event)
		fmt.Println()

	}

}

func main(){

	os.Setenv("SLACK_BOT_TOKEN", "xoxb-4391442328641-4378846496754-eVQ3uJRjsAjuGFzttNymJ6lH")
	os.Setenv("SLACK_APP_TOKEN","xapp-1-A04B4LTBETD-4375870187925-ff657a937cc05e81f9ffd012acf0ddf3140204de9f20ca51de2b9929478a7eb4")

	bot := slacker.NewClient(os.Getenv("SLACK_BOT_TOKEN"), os.Getenv("SLACK_APP_TOKEN"))

	go printCommandEvents(bot.CommandEvents())

	bot.Command("my yob is <year>", &slacker.CommandDefinition{
		Description: "yob calculator",
		// Examples: "my yob is 2020",
		Handler: func(botCtx slacker.BotContext, request slacker.Request, response slacker.ResponseWriter){
			year := request.Param("year")
			yob, err := strconv.Atoi(year)
			if err!= nil{
				println("error")
			}
			age := 2022-yob
			r := fmt.Sprintf("age is %d\n",age)
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