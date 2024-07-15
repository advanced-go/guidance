package routing1

import (
	"errors"
	"fmt"
	"github.com/advanced-go/stdlib/core"
	"net/url"
	"time"
)

// TODO :

const (
	RouteIdName   = "route_id"
	AuthorityName = "authority"
	VersionName   = "version-template"
	TrafficName   = "traffic"
	StatusName    = "status"
	RouteName     = "route"
)

var (
	//safeRoute = common.NewSafe()
	entryData = []Entry{
		{Region: "us-central1", Zone: "c", Host: "www.host1.com", Traffic: "egress", CreatedTS: time.Date(2024, 6, 10, 7, 120, 35, 0, time.UTC)},
		{Region: "us-central1", Zone: "c", Host: "www.host1.com", Traffic: "egress", CreatedTS: time.Date(2024, 6, 10, 7, 120, 35, 0, time.UTC)},
	}
)

// Entry - need to find some way to avoid duplicate processing, as host will start and stop due to
// Kubernetes scaling
type Entry struct {
	Region    string    `json:"region"`
	Zone      string    `json:"zone"`
	SubZone   string    `json:"sub-zone"`
	Host      string    `json:"host"`
	Traffic   string    `json:"traffic"`
	Route     string    `json:"route""`
	CreatedTS time.Time `json:"created-ts"`

	// Ingress or egress

	// Templates are only configured for egress traffic. Local is valid as is *. Blank does not include.
	VersionTemplate string `json:"version"` // Semantic versioning: 2.1.0
	RegionTemplate  string `json:"region"`
	ZoneTemplate    string `json:"zone"`
	SubZoneTemplate string `json:"sub-zone"`
}

/*
func (e Route) Origin() core.Origin {
	return core.Origin{Region: e.Region, Zone: e.Zone, SubZone: e.SubZone, Host: e.Host}
}


*/

func (Route) Scan(columnNames []string, values []any) (e Route, err error) {
	for i, name := range columnNames {
		switch name {
		case EntryIdName:
			e.EntryId = values[i].(int)
		case RouteIdName:
			e.EntryId = values[i].(int)
		case StatusName:
			e.Status = values[i].(string)
		case CreatedTSName:
			e.CreatedTS = values[i].(time.Time)

		case TrafficName:
			e.Traffic = values[i].(string)
		case RouteName:
			e.Name = values[i].(string)
		case AuthorityName:
			e.Authority = values[i].(string)

		case VersionName:
			e.Version = values[i].(string)
		case RegionName:
			e.Region = values[i].(string)
		case ZoneName:
			e.Zone = values[i].(string)
		case SubZoneName:
			e.SubZone = values[i].(string)

		default:
			err = errors.New(fmt.Sprintf("invalid field name: %v", name))
			return
		}
	}
	return
}

func (e Route) Values() []any {
	return []any{
		e.EntryId,
		e.RouteId,
		e.Status,
		e.CreatedTS,

		e.Traffic,
		e.Name,
		e.Authority,

		e.Version,
		e.Region,
		e.Zone,
		e.SubZone,
	}
}

func (Route) Rows(entries []Entry) [][]any {
	var values [][]any

	for _, e := range entries {
		values = append(values, e.Values())
	}
	return values
}

func validRoute(values url.Values, e Entry) bool {
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
