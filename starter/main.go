package main

import (
	"context"
	"log"
	"simple-bot-with-temporal/bot"

	"go.temporal.io/sdk/client"
)

func main() {
	c, err := client.Dial(client.Options{
		HostPort: client.DefaultHostPort,
	})
	if err != nil {
		log.Fatalln("Unable to connect to the client", err)
	}
	defer c.Close()
	workflowID := "cron_bot_set_trade"
	wo := client.StartWorkflowOptions{
		ID:        workflowID,
		TaskQueue: "cron_bot",
		// CronSchedule: "* * * * *",
	}
	link := "https://www.set.or.th/th/home"
	we, err := c.ExecuteWorkflow(context.Background(), wo, bot.BotWorkflow, link)
	if err != nil {
		log.Fatalln("Unable to execute the workflow", err)
	}
	log.Println("Started workflow", "WorkflowID", we.GetID(), "RunID", we.GetRunID())
}
