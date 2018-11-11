package main

var (
	kFreePlan = Plan{
		ID:         "1",
		TotalHosts: 1,
	}

	kSilverPlan = Plan{
		ID:         "2",
		TotalHosts: 10,
	}

	kGoldPlan = Plan{
		ID:         "3",
		TotalHosts: -1,
	}
)

// Plan types of plan and based on the how much we charge and its limits
type Plan struct {
	ID         string
	TotalHosts int
}
