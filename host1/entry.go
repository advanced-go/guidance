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

// CDCState - last ids
type CDCState struct {
	Resource  string
	LastCDCId int
}

// NewCDCState - create a new CDCState for a given resource: host,ingress,egress
func NewCDCState(rsc string) *CDCState {
	l := new(CDCState)
	l.LastCDCId = -1
	l.Resource = rsc
	return l
}

// Entry - host, utilize semantic versioning
// This needs to be updated every time a host starts up. Detail data needs to find the Entry based on the
// host name = entry detail key
// Need distinct constraint on: region+zone+subzone+host
type Entry struct {
	EntryId   int         `json:"entry-id"`
	Origin    core.Origin `json:"origin"`
	CreatedTS time.Time   `json:"created-ts"`
	DetailKey string      `json:"detail-key"` // How to query a detail entry, which is a part of the host
	Status    string      `json:"status"`     // active,in-active
}

type EntryStatus struct {
	Origin       core.Origin `json:"origin"`
	EntryCDCId   int         `json:"entry-cdc-id"`
	IngressCDCId int         `json:"ingress-cdc-id"`
	EgressCDCId  int         `json:"egress-cdc-id"`
}
