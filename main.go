package main

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"os"

	"github.com/Krognol/go-wolfram"
	"github.com/joho/godotenv"
	"github.com/shomali11/slacker"
	"github.com/tidwall/gjson"
	witai "github.com/wit-ai/wit-go/v2"
)

func main() {
	godotenv.Load()
	APP_TOKEN := os.Getenv("APP_TOKEN")
	BOT_TOKEN := os.Getenv("BOT_TOKEN")
	WIT_AI_TOKEN := os.Getenv("WIT_AI_TOKEN")
	WOLFRAM_APP_ID := os.Getenv("WOLFRAM_APP_ID")

	bot := slacker.NewClient(BOT_TOKEN, APP_TOKEN)

	client := witai.NewClient(WIT_AI_TOKEN)

	c := &wolfram.Client{AppID: WOLFRAM_APP_ID}

	bot.Command("query <sentence>", &slacker.CommandDefinition{
		Description: "Ask a question!",
		Examples:    []string{"Who is the prime minister of INDIA?"},
		Handler: func(botCtx slacker.BotContext, request slacker.Request, response slacker.ResponseWriter) {
			sentence := request.Param("sentence")
			msg, _ := client.Parse(&witai.MessageRequest{
				Query: sentence,
			})
			data, _ := json.MarshalIndent(msg, "", "    ")
			rough := string(data[:])
			fmt.Println(rough)
			val := gjson.Get(rough, "entities.wit$wolfram_search_query:wolfram_search_query.0.value")
			query := val.String()
			fmt.Println(query)

			res, err := c.GetSpokentAnswerQuery(query, wolfram.Metric, 1000)
			if err != nil {
				log.Fatal("Error occured while using wolfram api :", err)
			}
			fmt.Println(res)
			response.Reply(res)
		},
	})

	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	err := bot.Listen(ctx)
	if err != nil {
		log.Fatal(err)
	}
}
