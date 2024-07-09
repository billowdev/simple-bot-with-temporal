package bot

import (
	"time"

	"go.temporal.io/sdk/workflow"
)

// CronResult is used to return data from one cron run to the next
type CronResult struct {
	RunTime time.Time
	Result  interface{}
}

func BotSetTradeWorkflow(ctx workflow.Context, url string) (*CronResult, error) {
	workflow.GetLogger(ctx).Info("Cron workflow started.", "StartTime", workflow.Now(ctx))
	ao := workflow.ActivityOptions{
		StartToCloseTimeout: 10 * time.Second,
	}
	ctx1 := workflow.WithActivityOptions(ctx, ao)
	logger := workflow.GetLogger(ctx)
	logger.Info("Bot workflow started", "url", url)

	lastRunTime := time.Time{}
	if workflow.HasLastCompletionResult(ctx) {
		var lastResult CronResult
		if err := workflow.GetLastCompletionResult(ctx, &lastResult); err != nil {
			lastRunTime = lastResult.RunTime
		}
	}
	thisRunTime := workflow.Now(ctx)

	var result []SSetTrade

	err := workflow.ExecuteActivity(ctx1, BotSetTradeActivity, lastRunTime, thisRunTime, url).Get(ctx, &result)

	if err != nil {
		// logger.Error("Error executing bot activity", "error", err)
		workflow.GetLogger(ctx).Error("Cron job bot activity failed.", "Error", err)
		return nil, err
	}
	// logger.Info("Bot activity finished", "result", result)
	// return result, nil
	return &CronResult{
		RunTime: thisRunTime,
		Result:  result,
	}, nil
}

func BotGoldWorkflow(ctx workflow.Context, url string) (*CronResult, error) {
	workflow.GetLogger(ctx).Info("Cron workflow started.", "StartTime", workflow.Now(ctx))
	ao := workflow.ActivityOptions{
		StartToCloseTimeout: 10 * time.Second,
	}
	ctx1 := workflow.WithActivityOptions(ctx, ao)
	logger := workflow.GetLogger(ctx)
	logger.Info("Bot workflow started", "url", url)

	lastRunTime := time.Time{}
	if workflow.HasLastCompletionResult(ctx) {
		var lastResult CronResult
		if err := workflow.GetLastCompletionResult(ctx, &lastResult); err != nil {
			lastRunTime = lastResult.RunTime
		}
	}
	thisRunTime := workflow.Now(ctx)

	var result []SGold

	err := workflow.ExecuteActivity(ctx1, BotGoldActivity, lastRunTime, thisRunTime, url).Get(ctx, &result)

	if err != nil {
		// logger.Error("Error executing bot activity", "error", err)
		workflow.GetLogger(ctx).Error("Cron job bot activity failed.", "Error", err)
		return nil, err
	}
	// logger.Info("Bot activity finished", "result", result)
	// return result, nil
	return &CronResult{
		RunTime: thisRunTime,
		Result:  result,
	}, nil
}
