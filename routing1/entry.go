package routing1

import (
	"time"
)

// TODO : how to replicate these changes for the same host in different regions, zones, sub-zones
// Maybe allow the origin to be only the host??

const (
	EntryIdName        = "entry_id"
	CreatedTSName      = "created_ts"
	UpdatedTSName      = "updated_ts"
	RegionName         = "region"
	ZoneName           = "zone"
	SubZoneName        = "sub_zone"
	HostName           = "host"
	IngressVersionName = "ingress_version"
	EgressVersionName  = "egress_version"
)

var (
	//safeEntry = common.NewSafe()
	entryData = []Entry{
		{Region: "us-west1", Zone: "a", Host: "www.host1.com", CreatedTS: time.Date(2024, 6, 10, 7, 120, 35, 0, time.UTC)},
		{Region: "us-west1", Zone: "a", Host: "www.host2.com", CreatedTS: time.Date(2024, 6, 10, 7, 120, 35, 0, time.UTC)},
	}
)

func lastEntry() Entry {
	return entryData[len(entryData)-1]
}

// Entry - host, utilize semantic versioning
type Entry struct {
	EntryId   int       `json:"entry-id"`
	Region    string    `json:"region"`
	Zone      string    `json:"zone"`
	SubZone   string    `json:"sub-zone"`
	Host      string    `json:"host"`
	CreatedTS time.Time `json:"created-ts"`
	AgentId   string    `json:"agent-id"`
}

// Detail - not used
type Detail struct {
	EntryId   int       `json:"entry-id"`
	RouteName string    `json:"route"`
	CreatedTS time.Time `json:"created-ts"`
	AgentId   string    `json:"agent-id"`
	Config    string    `json:"config"`
}
