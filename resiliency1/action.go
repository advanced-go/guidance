package resiliency1

import "time"

// Actions - no updates as history with corresponding inference is needed

// RateLimitingAction - ingress and egress
// AgentId     string    `json:"agent-id"`
// Need to represent 2 states:
// 1. Nil or not configured - both values == -1
// 2. Configured - both values >= 0
type RateLimitingAction struct {
	EntryId     int       `json:"entry-id"`
	RouteName   string    `json:"route"`
	CreatedTS   time.Time `json:"created-ts"`
	InferenceId int       `json:"inference-id"`
	Limit       float64   `json:"limit"`
	Burst       int       `json:"burst"`
}

// RoutingAction - ingress and egress
// Need to determine how to represent 2 states:
// 1. Nil or not configured - location empty, percentage = -1
// 2. Re-routing in progress - location not empty, percentage >= 0
type RoutingAction struct {
	EntryId     int       `json:"entry-id"`
	RouteName   string    `json:"route"`
	CreatedTS   time.Time `json:"created-ts"`
	InferenceId int       `json:"inference-id"`
	Location    string    `json:"location"`
	Percentage  int       `json:"percentage"`
}

// RedirectAction - ingress and egress
type RedirectAction struct {
	EntryId     int       `json:"entry-id"`
	RouteName   string    `json:"route"`
	CreatedTS   time.Time `json:"created-ts"`
	InferenceId int       `json:"inference-id"`
	Location    string    `json:"location"`
	StatusCode  string    `json:"status-code"` // Only for ingress
}

type Actions struct {
	RouteName        string  `json:"route"`
	RedirectLocation string  `json:"redirect-location"`
	StatusCode       string  `json:"status-code"`
	Limit            float64 `json:"limit"`
	Burst            int     `json:"burst"`
	RouteLocation    string  `json:"route-location"`
	Percentage       int     `json:"percentage"`
}
