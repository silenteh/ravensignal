package main

// Plans should also contain various location around the world and the data retention

var (
	// kFreePlan allows for N amount of hosts to be scanned
	kFreePlan = Plan{
		ID:         "1",
		TotalHosts: 1,
		IntervalMS: 60000,
	}

	// kSilverPlan allows for N amount of hosts to be scanned
	kSilverPlan = Plan{
		ID:         "2",
		TotalHosts: 10,
		IntervalMS: 15000,
	}

	kGoldPlan = Plan{
		ID:         "3",
		TotalHosts: -1,
		IntervalMS: 1000,
	}
)

// Plan types of plan and based on the how much we charge and its limits
type Plan struct {
	ID         string `json:"id"`
	TotalHosts int    `json:"total_hosts"`
	IntervalMS int    `json:"interval_ms"` // the lowest interval the checks can have betwen one another
}
