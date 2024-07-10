package main

import (
	"context"
	"log"
	"simple-bot-with-temporal/bot"

	"go.temporal.io/sdk/client"
)

func RunSetTrade(c client.Client) {
	// workflowID := "bot_set_trade_" + uuid.New().String()
	workflowID := "bot_set_trade"
	wo := client.StartWorkflowOptions{
		ID:        workflowID,
		TaskQueue: "bot_worker",
		// CronSchedule: "* * * * *",
	}
	link := "https://www.set.or.th/th/home"
	we, err := c.ExecuteWorkflow(context.Background(), wo, bot.BotSetTradeWorkflow, link)
	if err != nil {
		log.Fatalln("Unable to execute the workflow", err)
	}
	log.Println("Started workflow", "WorkflowID", we.GetID(), "RunID", we.GetRunID())
}

func RunGold(c client.Client) {
	// workflowID := "bot_gold_price_today_" + uuid.New().String()
	workflowID := "bot_gold_price_today"
	wo := client.StartWorkflowOptions{
		ID:        workflowID,
		TaskQueue: "bot_worker",
		// CronSchedule: "* * * * *",
	}
	link := "https://xn--42cah7d0cxcvbbb9x.com/"
	we, err := c.ExecuteWorkflow(context.Background(), wo, bot.BotGoldWorkflow, link)
	if err != nil {
		log.Fatalln("Unable to execute the workflow", err)
	}
	log.Println("Started workflow", "WorkflowID", we.GetID(), "RunID", we.GetRunID())
}

func main() {
	c, err := client.Dial(client.Options{
		// HostPort: client.DefaultHostPort,
		HostPort: "172.16.40.38:7233",
	})
	if err != nil {
		log.Fatalln("Unable to connect to the client", err)
	}
	defer c.Close()
	RunSetTrade(c)
	RunGold(c)
}
