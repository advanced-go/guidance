package controller1

type RoutingPolicy struct {
	Primary     bool   `json:"primary"`
	Authority   string `json:"authority"` // github/advanced-go/observation: provider/account/repository
	AuthVersion string `json:"auth-version"`
	Prioritize  string `json:"prioritize"` // Origin, availability
	// OriginThreshold - when to prioritize by availability instead of origin
	OriginThreshold int `json:"origin-threshold"`
}

type DependencySchedule struct {
	HourFrom int    `json:"hour-from"`
	HourTo   int    `json:"hour-to"`
	Include  bool   `json:"include"`
	Days     string `json:"days"`
}

type Rowset struct {
	Version string `json:"version"`
	//ProcessingScheduleId string `json:"processing-schedule-id"`
	//DependencyUpdates    bool   `json:"dependency-updates"`
	Schedule DependencySchedule `json:"schedule"`

	RouteName string `json:"route"`
	//RateLimiting bool   `json:"rate-limiting"`
	//RegionT      string `json:"region-t"`
	//ZoneT        string `json:"zone-t"`
	//SubZoneT     string `json:"sub-zone-t"`
	//HostT        string `json:"host-t"`
	// Need to determine a cost metric. This choice is between consistency vs availability.
	// All consistency means only local zone routing.
	// Startup
	//Authority      string        `json:"authority"` // github/advanced-go/observation: provider/account/repository
	//AuthVersion    string        `json:"auth-version"`
	StartupPolicy  RoutingPolicy `json:"startup-policy"`
	FailoverPolicy RoutingPolicy `json:"failover-policy"`

	// Notifications
	Email string
	Slack string
}
