package bot

import "time"

// WorkflowResult is used to return data from one cron run to the next
type WorkflowResult struct {
	RunTime time.Time
	Result  interface{}
}

type SSetTrade struct {
	Index  string
	Price  string
	Change string
	Volume string
	Value  string
}

type SGold struct {
	Type string
	Buy  string
	Sell string
}
