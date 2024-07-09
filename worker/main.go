package main

import (
	"fmt"
	"log"
	"simple-bot-with-temporal/bot"

	"go.temporal.io/sdk/client"
	"go.temporal.io/sdk/worker"
)

func main() {
	c, err := client.Dial(client.Options{
		HostPort: client.DefaultHostPort,
	})
	if err != nil {
		log.Fatalln("Unable to create client", err)
	}
	defer c.Close()

	w := worker.New(c, "cron_bot", worker.Options{})
	w.RegisterWorkflow(bot.BotWorkflow)
	w.RegisterActivity(bot.BotActivity)

	err = w.Run(worker.InterruptCh())
	if err != nil {
		fmt.Println("-------err----------")
		fmt.Println(err)
		fmt.Println("-------err----------")
		log.Fatalln("Unable to start worker", err)
	}
}
