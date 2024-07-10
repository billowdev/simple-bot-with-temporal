package bot

import (
	"context"
	"time"

	"go.temporal.io/sdk/activity"
	"go.temporal.io/sdk/workflow"
)

func BotGoldWorkflow(ctx workflow.Context, url string) (*WorkflowResult, error) {
	workflow.GetLogger(ctx).Info("Cron workflow started.", "StartTime", workflow.Now(ctx))
	ao := workflow.ActivityOptions{
		StartToCloseTimeout: 10 * time.Second,
	}
	ctx1 := workflow.WithActivityOptions(ctx, ao)
	logger := workflow.GetLogger(ctx)
	logger.Info("Bot workflow started", "url", url)

	lastRunTime := time.Time{}
	if workflow.HasLastCompletionResult(ctx) {
		var lastResult WorkflowResult
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

	_ = workflow.ExecuteActivity(ctx1, BotGoldActivity2, lastRunTime, thisRunTime, url).Get(ctx, nil)

	// logger.Info("Bot activity finished", "result", result)
	// return result, nil
	return &WorkflowResult{
		RunTime: thisRunTime,
		Result:  result,
	}, nil
}

func BotGoldActivity(ctx context.Context, lastRunTime, thisRunTime time.Time, url string) ([]SGold, error) {
	// Query database, call external API, or do any other non-deterministic action.
	// logger := activity.GetLogger(ctx)
	// logger.Info("")
	activity.GetLogger(ctx).Info("Cron job running.", "lastRunTime_exclude", lastRunTime, "thisRunTime_include", thisRunTime)
	return ScraperGold(url)
	// return nil
}
func BotGoldActivity2(ctx context.Context, lastRunTime, thisRunTime time.Time, url string) ([]SGold, error) {
	// Query database, call external API, or do any other non-deterministic action.
	// logger := activity.GetLogger(ctx)
	// logger.Info("")
	activity.GetLogger(ctx).Info("Cron job running.", "lastRunTime_exclude", lastRunTime, "thisRunTime_include", thisRunTime)
	return ScraperGold2(url)
	// return nil
}
