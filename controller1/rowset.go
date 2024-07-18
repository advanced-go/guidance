package controller1

type Rowset struct {
	Version              string `json:"version"`
	ProcessingScheduleId string `json:"processing-schedule-id"`
	DependencyUpdates    bool   `json:"dependency-updates"`
	DependencyScheduleId string `json:"dependency-schedule-id"`

	RouteName    string `json:"route"`
	RateLimiting bool   `json:"rate-limiting"`
	RegionT      string `json:"region-t"`
	ZoneT        string `json:"zone-t"`
	SubZoneT     string `json:"sub-zone-t"`
	HostT        string `json:"host-t"`
	Authority    string `json:"authority"` // github/advanced-go/observation: provider/account/repository
	AuthVersion  string `json:"auth-version"`

	// Notifications
	Email string
	Slack string
}