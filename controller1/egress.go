package controller1

import "time"

// Need to log secondary for a redirect!!
// Need to notify if there is an issue
// Need an egress redirect agent??
// How would this get started?? Maybe when a temp redirect is received then notify cloud??
// Cloud would need to know to stop failover.
// When the temporary

// Issues, questions, and decisions.
// Where does routing matching go? Client or Cloud?
// Can static configurations go on the client?? YES
// On startup, the default host names need to be downloaded to client and a notification needs to be
// sent to the appropriate Case officer to configure the Egress agents.
//  - What if agent is in a dynamic routing change? Is there any state that should be saved??
//  - What if agent is in a failover state? Should be OK to restart. State can be lost.

// EgressRoute -
// Cardinality - n
// Access - EgressAgent
// Update - User Changeset
type EgressRoute struct {
	EntryId   int       `json:"entry-id"`   // How to refer to the main entry
	VersionId string    `json:"version-id"` // How to version this artifact
	RouteName string    `json:"route-name"`
	CreatedTS time.Time `json:"created-ts"`
	AgentId   string    `json:"agent-id"` // Auditing

	RateLimiting bool `json:"rate-limiting"`

	// Need some cost metrics to determine when to route to a secondary?
	// Can this be user configurable??

	// Is there a need to configure host selection based on whether this is on startup/new pod
	// vs failover??
	// Can this be handled by Authority role? Startup only filters by role and returns all regions.
	// dynamic routing picks the closest one.

	Authority        string `json:"authority"` // github/advanced-go/observation: provider/account/repository
	AuthorityVersion string `json:"authority-version"`

	// Templates for host selection, "local" is valid as is "*". Blank does not include.
	// Only used for failover
	RegionT  string `json:"region-t"`
	ZoneT    string `json:"zone-t"`
	SubZoneT string `json:"sub-zone-t"`
}

/*
// EgressRoutePolicy - provides static and dynamic routing configuration + processing for egress routing
// These are always inserted with a date. Never updated
type EgressRoutePolicy struct {
	EntryId   int       `json:"entry-id"`   // How to refer to the main entry
	VersionId string    `json:"version-id"` // How to version this artifact
	RouteName string    `json:"route-name"`
	CreatedTS time.Time `json:"created-ts"`
	AgentId   string    `json:"agent-id"` // Auditing

	// Need some cost metrics to determine when to route to a secondary?
	// Can this be user configurable??

	// Is there a need to configure host selection based on whether this is on startup/new pod
	// vs failover??
	// Can this be handled by Authority role? Startup only filters by role and returns all regions.
	// dynamic routing picks the closest one.

	Authority        string `json:"authority"` // github/advanced-go/observation: provider/account/repository
	AuthorityVersion string `json:"authority-version"`

	// Templates for host selection, "local" is valid as is "*". Blank does not include.
	// Only used for failover
	RegionT  string `json:"region-t"`
	ZoneT    string `json:"zone-t"`
	SubZoneT string `json:"sub-zone-t"`
}


*/
