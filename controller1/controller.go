package controller1

// ControllerRedirect - ingress redirection
type ControllerRedirect struct {
	Location string `json:"location"`
	Status   string `json:"status"`
}

// ControllerRouting - egress routing
type ControllerRouting struct {
	Scope string `json:"scope"` // SubZone, Zone, Region, *, empty or none -> not configured
	// Threshold - when routing changes occur.
	// Value == -1 -> let system determine
	// Value == 0  -> no threshold, re-routing immediately
	// Value > 0   -> re-routing when threshold is met
	Threshold int `json:"threshold"`
}
