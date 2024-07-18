package controller1

import "time"

// No more static vs dynamic routing, all routing is a primary with a secondary used for failover
// and conversion.
// May add dynamic routing later, meaning a primary is set at runtime,
// Re-evaluate if needed. Static routing - can use dynamic routing for failover if templates are configured
//StaticHost string `json:"host"`

// Policy - provides static and dynamic routing configuration + processing for egress routing
// These are always inserted with a date. Never updated
type Policy struct {
	Region    string    `json:"region"`
	Zone      string    `json:"zone"`
	SubZone   string    `json:"sub-zone"`
	Host      string    `json:"host"`
	RouteName string    `json:"route"`
	PolicyId  int       `json:"policy-id"`
	Version   string    `json:"version"` // Current Semantic version: 2.1.0
	CreatedTS time.Time `json:"created-ts"`
	AgentId   string    `json:"agent-id"` // Could be a user or agent id

	// Need some cost metrics to determine when to route to a secondary?
	// Can this be user configurable??

	// Current version
	AuthorityVersion string `json:"auth-version"` // Current Semantic version: 2.1.0
	AuthorityHost    string `json:"auth-host"`    // Primary host override.

	// Templates for host selection, "local" is valid as is "*". Blank does not include.
	RegionT  string `json:"region-t"`
	ZoneT    string `json:"zone-t"`
	SubZoneT string `json:"sub-zone-t"`

	// Dynamic conversion
	VersionT string    `json:"version-t"`
	FromTS   time.Time `json:"from-ts"` // Timespan for conversion
	ToTS     time.Time `json:"to-ts"`
}
