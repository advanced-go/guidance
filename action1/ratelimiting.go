package action1

import "time"

// RateLimiting - ingress and egress
// AgentId     string    `json:"agent-id"`
// Need to represent 2 states:
// 1. Nil or not configured - both values == -1
// 2. Configured - both values >= 0
type RateLimiting struct {
	EntryId     int       `json:"entry-id"`
	RouteName   string    `json:"route"`
	CreatedTS   time.Time `json:"created-ts"`
	InferenceId int       `json:"inference-id"`
	Limit       float64   `json:"limit"`
	Burst       int       `json:"burst"`
}
