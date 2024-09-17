package state1

import "github.com/advanced-go/stdlib/core"

// IngressRedirectState - ingress redirect state
type IngressRedirectState struct {
	EntryId int    `json:"entry-id"`
	Route   string `json:"route"`

	// Redirect plan
	Location string `json:"location"`
	Status   string `json:"status"` // Scheduled,In-Progress,Completed,Failed

	// Percentile SLO
	Percent int `json:"percent"` // Used for latency, traffic, status codes, counter, profile
	Latency int `json:"latency"` // Used for latency, saturation duration or traffic
	Minimum int `json:"minimum"` // Used for status codes to attenuate underflow, applied to the window interval

	// Routing action
	Percentage int `json:"percentage"`
}

// NewIngressRedirectState - initialize
func NewIngressRedirectState(s *IngressRedirectState) {
	s.Location = ""
	s.Status = ""
	s.Percent = DefaultPercentileSLO.Percent
	s.Latency = DefaultPercentileSLO.Latency
	s.Minimum = DefaultPercentileSLO.Minimum
	s.Percentage = -1
}

func (r *IngressRedirectState) IsActive() bool     { return r.Percentage != -1 }
func (r *IngressRedirectState) IsConfigured() bool { return r.Location != "" }

// IngressResiliencyState - ingress resiliency state
type IngressResiliencyState struct {
	EntryId int    `json:"entry-id"`
	Route   string `json:"route"`

	// Percentile SLO
	Percent int `json:"percent"` // Used for latency, traffic, status codes, counter, profile
	Latency int `json:"latency"` // Used for latency, saturation duration or traffic
	Minimum int `json:"minimum"` // Used for status codes to attenuate underflow, applied to the window interval

	// Rate Limiting action
	Limit float64 `json:"limit"`
	Burst int     `json:"burst"`
}

// NewIngressResiliencyState - initialize
func NewIngressResiliencyState(s *IngressResiliencyState) {
	s.Percent = DefaultPercentileSLO.Percent
	s.Latency = DefaultPercentileSLO.Latency
	s.Minimum = DefaultPercentileSLO.Minimum
	s.Limit = -1
	s.Burst = -1
}

// No IsConfigured() as ingress resiliency is automatic without needing a plan

// IsActive - active
func (r *IngressResiliencyState) IsActive() bool { return r.Limit == -1 }

// EgressState - egress state, needs origin for agent creation
type EgressState struct {
	EntryId int    `json:"entry-id"`
	Region  string `json:"region"`
	Zone    string `json:"zone"`
	SubZone string `json:"sub-zone"`
	Host    string `json:"host"`
	Route   string `json:"route"`

	// Failover configuration
	FailoverScope     string `json:"scope"` // SubZone, Zone, Region, *, empty or none -> not configured
	FailoverThreshold int    `json:"threshold"`

	// Rate Limiting action
	Limit float64 `json:"limit"`
	Burst int     `json:"burst"`

	// Routing action
	Location   string `json:"location"`
	Percentage int    `json:"percentage"`
}

// NewEgressState - initialize
func NewEgressState(s *EgressState) {
	s.FailoverScope = ""
	s.FailoverThreshold = -1
	s.Limit = -1
	s.Burst = -1
	s.Location = ""
	s.Percentage = -1
}

func (r *EgressState) IsRateLimitingActive() bool { return r.Limit != -1 }
func (r *EgressState) IsRoutingActive() bool      { return r.Percentage != -1 }
func (r *EgressState) IsConfigured() bool         { return r.FailoverScope != "" }
func (r *EgressState) Origin() core.Origin {
	return core.Origin{
		Region:     r.Region,
		Zone:       r.Zone,
		SubZone:    r.SubZone,
		Host:       r.Host,
		InstanceId: "",
		Route:      r.Route,
	}
}
