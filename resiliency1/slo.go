package resiliency1

import "time"

// PercentileSLO - percentile config
// TODO : need to add CDC to this table for Agent data change notifications.
type PercentileSLO struct {
	EntryId   int       `json:"entry-id"`
	RouteName string    `json:"route"`
	AgentId   string    `json:"agent-id"`
	CreatedTS time.Time `json:"created-ts"`
	//Watch   int // Range 1 - 99
	Percent int // Used for latency, traffic, status codes, counter, profile
	Latency int // Used for latency, saturation duration or traffic
	Minimum int // Used for status codes to attenuate underflow, applied to the window interval
}
