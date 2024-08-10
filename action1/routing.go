package action1

import "time"

// Routing - ingress and egress
// Need to determine how to represent 2 states:
// 1. Nil or not configured - location empty, percentage = -1
// 2. Re-routing in progress - location not empty, percentage >= 0
type Routing struct {
	EntryId     int       `json:"entry-id"`
	RouteName   string    `json:"route"`
	CreatedTS   time.Time `json:"created-ts"`
	InferenceId int       `json:"inference-id"`
	Location    string    `json:"location"`
	Percentage  int       `json:"percentage"`
}
