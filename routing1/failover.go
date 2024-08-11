package routing1

// Failover - routing config
// TODO : need to add CDC to this table for Agent data change notifications.
// Threshold - when routing changes occur.
// Value == -1 -> let system determine
// Value == 0  -> no threshold, re-routing immediately
// Value > 0   -> re-routing when threshold is met
type Failover struct {
	//EntryId   int       `json:"entry-id"`
	RouteName string `json:"route"`
	//CreatedTS time.Time `json:"created-ts"`
	//AgentId   string    `json:"agent-id"`
	Scope     string `json:"scope"` // SubZone, Zone, Region, *, empty or none -> not configured
	Threshold int    `json:"threshold"`
}
