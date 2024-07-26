package controller1

import "time"

// Need to log secondary for a redirect!!
// Need to notify if there is an issue
// Need an egress redirect agent??
// How would this get started?? Maybe when a temp redirect is received then notify cloud??
// Cloud would need to know to stop failover.
// When the temporary

// Egress -
// Cardinality - n
// Access - EgressAgent
// Update - User Changeset
type Egress struct {
	EntryId   int       `json:"entry-id"`   // How to refer to the main entry
	VersionId string    `json:"version-id"` // How to version this artifact
	RouteName string    `json:"route-name"`
	CreatedTS time.Time `json:"created-ts"`
	AgentId   string    `json:"agent-id"` // Auditing

	Location string `json:"location"` // Redirect location
	// Is there a need for a list of secondary authorities??
	// Or maybe some sort of authority template if the authority name changes between scopes
	// AutuhorityT  - authority template
	FailoverScope string `json:"failover-scope"` // SubZone, Zone, Region, *, empty or none -> not configured
	// FailureThreshold - when routing changes occur.
	// Value == -1 -> let system determine
	// Value == 0  -> no threshold, failover immediately
	// Value > 0   -> failover when threshold is met
	FailoverThreshold int `json:"failover-threshold"`

	// Need some cost metrics to determine when to route to a secondary?
	// Can this be user configurable??

}

func (e Egress) IsEmpty() bool {
	return e.EntryId <= 0
}

func (e Egress) IsRedirect() bool {
	return e.Location != ""
}

func (e Egress) IsFailover() bool {
	return e.FailoverScope != ""
}
