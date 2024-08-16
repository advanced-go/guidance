package resiliency1

type IngressState struct {
	// Redirect plan
	Location string `json:"location"`
	Status   string `json:"status"` // Scheduled,In-Progress,Completed,Failed

	// Percentile SLO
	Percent int `json:"percent"` // Used for latency, traffic, status codes, counter, profile
	Latency int `json:"latency"` // Used for latency, saturation duration or traffic
	Minimum int `json:"minimum"` // Used for status codes to attenuate underflow, applied to the window interval

	// Rate Limiting action
	Limit float64 `json:"limit"`
	Burst int     `json:"burst"`

	// Routing action
	Percentage int `json:"percentage"`
}

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
