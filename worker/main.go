package main

import (
	"log"
	"simple-bot-with-temporal/bot"

	"go.temporal.io/sdk/client"
	"go.temporal.io/sdk/worker"
)

func register(w worker.Registry) {
	w.RegisterWorkflow(bot.BotSetTradeWorkflow)
	w.RegisterActivity(bot.BotSetTradeActivity)

	w.RegisterWorkflow(bot.BotGoldWorkflow)
	w.RegisterActivity(bot.BotGoldActivity)

}

func main() {
	// Create a client connection to Temporal service
	c, err := client.Dial(client.Options{
		// HostPort: client.DefaultHostPort,
		HostPort: "172.16.40.38:7233",
	})
	// http://172.16.40.38:8088/namespaces/default/task-queues/cron_bot_gold_taskqueue
	if err != nil {
		log.Fatalln("Unable to create client", err)
	}
	defer c.Close()
	w := worker.New(c, "bot_worker", worker.Options{})
	register(w)
	if err := w.Run(worker.InterruptCh()); err != nil {
		log.Fatalln("Unable to start goldPriceWorker", err)
	}

	// // Define and start workers for each task queue
	// go func() {
	// 	setTradeWorker := worker.New(c, "cron_bot_settrade_taskqueue", worker.Options{})
	// 	setTradeWorker.RegisterWorkflow(bot.BotSetTradeWorkflow)
	// 	setTradeWorker.RegisterActivity(bot.BotSetTradeActivity)

	// 	if err := setTradeWorker.Run(worker.InterruptCh()); err != nil {
	// 		log.Fatalln("Unable to start setTradeWorker", err)
	// 	}
	// }()

	// go func() {
	// 	goldPriceWorker := worker.New(c, "cron_bot_gold_taskqueue", worker.Options{})
	// 	goldPriceWorker.RegisterWorkflow(bot.BotGoldWorkflow)
	// 	goldPriceWorker.RegisterActivity(bot.BotGoldActivity)

	// 	if err := goldPriceWorker.Run(worker.InterruptCh()); err != nil {
	// 		log.Fatalln("Unable to start goldPriceWorker", err)
	// 	}
	// }()

	// Block indefinitely to keep the main goroutine running
	// select {}
}
