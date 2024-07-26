package controller1

import (
	"errors"
	"fmt"
	"github.com/advanced-go/stdlib/core"
	"net/url"
	"time"
)

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
	UpdatedTS time.Time `json:"updated-ts"` // Used to optimize restarts

	// Status - Active, Inactive, Removed. Is this needed??
	Status string `json:"status"`

	// Current version - auditing via CDC
	IngressVersion string `json:"ingress-version"`
	EgressVersion  string `json:"egress-version"`
}

func (e Entry) IsEmpty() bool {
	return e.EntryId <= 0
}

func (e Entry) Origin() core.Origin {
	return core.Origin{Region: e.Region, Zone: e.Zone, SubZone: e.SubZone, Host: e.Host}
}

func (Entry) Scan(columnNames []string, values []any) (e Entry, err error) {
	for i, name := range columnNames {
		switch name {
		case EntryIdName:
			e.EntryId = values[i].(int)

		case CreatedTSName:
			e.CreatedTS = values[i].(time.Time)

		case RegionName:
			e.Region = values[i].(string)
		case ZoneName:
			e.Zone = values[i].(string)
		case SubZoneName:
			e.SubZone = values[i].(string)
		case HostName:
			e.Host = values[i].(string)
		case IngressVersionName:
			e.IngressVersion = values[i].(string)
		case EgressVersionName:
			e.EgressVersion = values[i].(string)

		default:
			err = errors.New(fmt.Sprintf("invalid field name: %v", name))
			return
		}
	}
	return
}

func (e Entry) Values() []any {
	return []any{
		e.CreatedTS,

		e.Region,
		e.Zone,
		e.SubZone,
		e.Host,
		e.IngressVersion,
		e.EgressVersion,
	}
}

func (Entry) Rows(entries []Entry) [][]any {
	var values [][]any

	for _, e := range entries {
		values = append(values, e.Values())
	}
	return values
}

func validEntry(values url.Values, e Entry) bool {
	if values == nil {
		return false
	}
	filter := core.NewOrigin(values)
	target := core.Origin{Region: e.Region, Zone: e.Zone, SubZone: e.SubZone, Host: e.Host}
	if !core.OriginMatch(target, filter) {
		return false
	}
	// Additional filtering
	return true
}
