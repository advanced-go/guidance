package host1

import (
	"github.com/advanced-go/stdlib/core"
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

// LastCDCId -
type LastCDCId struct {
	Entry    int
	Redirect int
	Egress   int
}

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
// This needs to be updated every time a host starts up. Detail data needs to find the Entry based on the
// host name = entry detail key
// Need distinct constraint on: region+zone+subzone+host
type Entry struct {
	EntryId   int       `json:"entry-id"`
	Region    string    `json:"region"`
	Zone      string    `json:"zone"`
	SubZone   string    `json:"sub-zone"`
	Host      string    `json:"host"`
	CreatedTS time.Time `json:"created-ts"`
	DetailKey string    `json:"detail-key"` // How to query a detail entry, which is a part of the host
	Status    string    `json:"status"`     // active,in-active
}

func (e Entry) Origin() core.Origin {
	return core.Origin{
		Region:  e.Region,
		Zone:    e.Zone,
		SubZone: e.SubZone,
		Host:    e.Host,
		//InstanceId: "",
		//Route:      "",
	}
}
