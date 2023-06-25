package main

import (
	"context"
	"fmt"
	"log"
	"os"

	"github.com/joho/godotenv"
	"github.com/shomali11/slacker"
	witai "github.com/wit-ai/wit-go/v2"
)

func main() {
	godotenv.Load()
	APP_TOKEN := os.Getenv("APP_TOKEN")
	BOT_TOKEN := os.Getenv("BOT_TOKEN")
	WIT_AI_TOKEN := os.Getenv("WIT_AI_TOKEN")

	bot := slacker.NewClient(BOT_TOKEN, APP_TOKEN)

	client := witai.NewClient(WIT_AI_TOKEN)

	bot.Command("<sentence>", &slacker.CommandDefinition{
		Description: "Say a sentence!",
		Examples:    []string{"say hello there everyone!"},
		Handler: func(botCtx slacker.BotContext, request slacker.Request, response slacker.ResponseWriter) {
			sentence := request.Param("sentence")
			fmt.Println(sentence)
			msg, _ := client.Parse(&witai.MessageRequest{
				Query: sentence,
			})
			fmt.Printf("%v", msg.Entities)
			fmt.Printf("%v", msg.ID)
			fmt.Printf("%v", msg.Intents)
			fmt.Printf("%v", msg.Text)
			fmt.Printf("%v", msg.Traits)
			response.Reply(sentence)
		},
	})

	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	err := bot.Listen(ctx)
	if err != nil {
		log.Fatal(err)
	}
}
