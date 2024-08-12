package resiliency1

import "time"

// Percentile - percentile config
// TODO : need to add CDC to this table for Agent data change notifications.
type Percentile struct {
	EntryId   int       `json:"entry-id"`
	CreatedTS time.Time `json:"created-ts"`
	AgentId   string    `json:"agent-id"`
	RouteName string    `json:"route"`
	//Watch   int // Range 1 - 99
	Percent int // Used for latency, traffic, status codes, counter, profile
	Latency int // Used for latency, saturation duration or traffic
	Minimum int // Used for status codes to attenuate underflow, applied to the window interval
}
