package types

// https://localhost:8081/github/advanced-go/guidance:region/zone/subzone/app

type Timeout struct {
	Duration string `json:"duration"`
}

type RateLimiter struct {
	Limit string `json:"limit"`
	Burst int    `json:"burst"`
}

type CostFunction struct {
	Threshold string `json:"threshold"`
}

type Routing struct {
	Active    string
	Primary   string
	Secondary string
	//FailureThreshold int // For failures
	StepPercentage int // Steps for traffic routing
}

type EntryV1 struct {
	Name         string `json:"name"`
	Status       int    `json:"status"`
	CreatedTS    string `json:"created-ts"`
	UpdatedTS    string `json:"updated-ts"`
	CostFunction CostFunction
	Routing      Routing
	Timeout      Timeout
	RateLimiter  RateLimiter
}
