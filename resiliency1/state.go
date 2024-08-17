package resiliency1

type IngressRedirectState struct {
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

func (r *IngressRedirectState) IsActive() bool     { return r.Percentage != -1 }
func (r *IngressRedirectState) IsConfigured() bool { return r.Location != "" }

type IngressResiliencyState struct {
	// Percentile SLO
	Percent int `json:"percent"` // Used for latency, traffic, status codes, counter, profile
	Latency int `json:"latency"` // Used for latency, saturation duration or traffic
	Minimum int `json:"minimum"` // Used for status codes to attenuate underflow, applied to the window interval

	// Rate Limiting action
	Limit float64 `json:"limit"`
	Burst int     `json:"burst"`
}

// No IsConfigured() as ingress resiliency is automatic without needing a plan

// IsActive - active
func (r *IngressResiliencyState) IsActive() bool { return r.Limit == -1 }

type EgressState struct {
	RouteName string `json:"route"`

	// Failover plan
	Scope     string `json:"scope"` // SubZone, Zone, Region, *, empty or none -> not configured
	Threshold int    `json:"threshold"`

	// Rate Limiting action
	Limit float64 `json:"limit"`
	Burst int     `json:"burst"`

	// Routing action
	Location   string `json:"location"`
	Percentage int    `json:"percentage"`
}

func (r *EgressState) IsRateLimitingActive() bool { return r.Limit != -1 }
func (r *EgressState) IsRoutingActive() bool      { return r.Percentage != -1 }
func (r *EgressState) IsConfigured() bool         { return r.Scope != "" }
