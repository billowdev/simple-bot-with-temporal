package bot

import (
	"context"
	"time"

	"go.temporal.io/sdk/activity"
)

func BotActivity(ctx context.Context, lastRunTime, thisRunTime time.Time, url string) ([]SSetTrade, error) {
	// logger := activity.GetLogger(ctx)
	// logger.Info("")
	activity.GetLogger(ctx).Info("Cron job running.", "lastRunTime_exclude", lastRunTime, "thisRunTime_include", thisRunTime)
	// // Query database, call external API, or do any other non-deterministic action.
	return Scraper(url)
	// return nil
}
