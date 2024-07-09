package bot

import (
	"context"
	"time"

	"go.temporal.io/sdk/activity"
)

func BotSetTradeActivity(ctx context.Context, lastRunTime, thisRunTime time.Time, url string) ([]SSetTrade, error) {
	// Query database, call external API, or do any other non-deterministic action.
	// logger := activity.GetLogger(ctx)
	// logger.Info("")
	activity.GetLogger(ctx).Info("Cron job running.", "lastRunTime_exclude", lastRunTime, "thisRunTime_include", thisRunTime)
	return ScraperSet50(url)
	// return nil
}

func BotGoldActivity(ctx context.Context, lastRunTime, thisRunTime time.Time, url string) ([]SGold, error) {
	// Query database, call external API, or do any other non-deterministic action.
	// logger := activity.GetLogger(ctx)
	// logger.Info("")
	activity.GetLogger(ctx).Info("Cron job running.", "lastRunTime_exclude", lastRunTime, "thisRunTime_include", thisRunTime)
	return ScraperGold(url)
	// return nil
}
