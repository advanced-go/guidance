package resiliency1

import "time"

var (
	DefaultPercentileSLO = &PercentileSLO{Percent: 99, Latency: 2000, Minimum: 100}
)

// PercentileSLO - percentile config
// TODO : need to add CDC to this table for Agent data change notifications.
type PercentileSLO struct {
	EntryId   int       `json:"entry-id"`
	Route     string    `json:"route"`
	AgentId   string    `json:"agent-id"`
	CreatedTS time.Time `json:"created-ts"`

	Percent int `json:"percent"` // Used for latency, traffic, status codes, counter, profile
	Latency int `json:"latency"` // Used for latency, saturation duration or traffic
	Minimum int `json:"minimum"` // Used for status codes to attenuate underflow, applied to the window interval
}

func NewPercentileSLO(s *PercentileSLO) {
	s.Minimum = -1
	s.Latency = -1
	s.Percent = -1
}
