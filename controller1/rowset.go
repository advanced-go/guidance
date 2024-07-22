package controller1

// Version policy
// Startup 2.3.* - any patch, prefer the latest
// Conversion2.3.* - automatically update dependency with a change request
// Version supports wildcards allowing control of dependency updates. so least restrictive to most restrictive would be:
//  2.*
//  2.3.*
//  2.3.12

// RoutingPolicy - routing information
type RoutingPolicy struct {
	// Allows wildcards so that host selection is less restrictive than static version
	// Usually a less restrictive version, so that for failover, a host with an acceptable
	// version can be accepted.
	AuthVersion string `json:"auth-version"`
	// Scope for authority selection
	Scope string `json:"scope"` // SubZone, Zone, Region, *.
	// FailureThreshold - when routing changes occur.
	// Value == -1 -> let system determine
	// Value == 0  -> no threshold, failover immediately
	// Value > 0   -> failover when threshold is met
	FailureThreshold int `json:"failure-threshold"`
}

type DependencyPolicy struct {
	// Changeset approval needed. Ever change gets a changeset
	ChangesetApproval bool `json:"changeset-approval"`
	// How to choose a host, version with wildcards. This version is less restrictive than in the rowset.
	// Scope defaults to 'Zone'
	Scope       string `json:"scope"` // SubZone if used, then Zone, use most restrictive
	AuthVersion string `json:"auth-version"`
	HourFrom    int    `json:"hour-from"`
	HourTo      int    `json:"hour-to"`
	Include     bool   `json:"include"`
	Days        string `json:"days"`
}

type Rowset struct {
	Version    string           `json:"version"`
	Conversion DependencyPolicy `json:"dependency-policy"`

	RouteName string `json:"route"`
	// Always favor a primary authority
	Authority string `json:"authority"` // github/advanced-go/observation: provider/account/repository
	// This has to be specific, so that version changes can be determined for failover, conversion.
	// 2.3.12
	// Changing the auth version should change the failover policy version
	AuthVersion string `json:"auth-version"`

	//StartupPolicy  RoutingPolicy `json:"startup-policy"`
	// If a failure on startup, then go to failover.
	FailoverPolicy RoutingPolicy `json:"failover-policy"`
	// For failover routing
	//AuthVersionT string`json:"auth-version-t"`
	//FailureThreshold int `json:"failure-threshold"`
}

//ProcessingScheduleId string `json:"processing-schedule-id"`
//DependencyUpdates    bool   `json:"dependency-updates"`
//RateLimiting bool   `json:"rate-limiting"`
//RegionT      string `json:"region-t"`
//ZoneT        string `json:"zone-t"`
//SubZoneT     string `json:"sub-zone-t"`
//HostT        string `json:"host-t"`
// Need to determine a cost metric. This choice is between consistency vs availability.
// All consistency means only local zone routing.
// Startup
