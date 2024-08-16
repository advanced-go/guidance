package resiliency1

import "time"

// FailoverPlan - egress routing
// TODO : need to add CDC to this table for Agent data change notifications.
// Threshold - when routing changes occur.
// Value == -1 -> let system determine
// Value == 0  -> no threshold, re-routing immediately
// Value > 0   -> re-routing when threshold is met
type FailoverPlan struct {
	EntryId   int       `json:"entry-id"`
	RouteName string    `json:"route"`
	AgentId   string    `json:"agent-id"`
	CreatedTS time.Time `json:"created-ts"`
	Scope     string    `json:"scope"` // SubZone, Zone, Region, *, empty or none -> not configured
	Threshold int       `json:"threshold"`
}

// RedirectPlan - ingress redirect
// TODO : need to add CDC to this table for Agent data change notifications.
type RedirectPlan struct {
	EntryId   int       `json:"entry-id"`
	RouteName string    `json:"route"`
	AgentId   string    `json:"agent-id"`
	CreatedTS time.Time `json:"created-ts"`
	UpdatedTS time.Time `json:"updated-ts"`
	Location  string    `json:"location"`
	Status    string    `json:"status"` // Scheduled,In-Progress,Completed,Failed
}
