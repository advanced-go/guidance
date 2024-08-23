package resiliency1

import (
	"github.com/advanced-go/stdlib/core"
	"time"
)

// EgressConfig - egress routing
// TODO : need to add CDC to this table for Agent data change notifications.
// Threshold - when routing changes occur.
// Value == -1 -> let system determine
// Value == 0  -> no threshold, re-routing immediately
// Value > 0   -> re-routing when threshold is met
type EgressConfig struct {
	EntryId    int       `json:"entry-id"`
	Region     string    `json:"region"`
	Zone       string    `json:"zone"`
	SubZone    string    `json:"sub-zone"`
	Host       string    `json:"host"`
	RouteName  string    `json:"route"`
	AgentId    string    `json:"agent-id"`
	SQLCommand string    `json:"sql-command"` // insert,update,delete
	CreatedTS  time.Time `json:"created-ts"`

	// Failover configuration
	Scope     string `json:"scope"`     // SubZone, Zone, Region, *, empty or none -> not configured
	Threshold int    `json:"threshold"` // Percentage of traffic
}

func (p EgressConfig) Origin() core.Origin {
	return core.Origin{
		Region:  p.Region,
		Zone:    p.Zone,
		SubZone: p.SubZone,
		Host:    p.Host,
		Route:   p.RouteName,
	}
}

// RedirectConfig - ingress redirect
// TODO : need to add CDC to this table for Agent data change notifications.
type RedirectConfig struct {
	EntryId    int       `json:"entry-id"`
	Region     string    `json:"region"`
	Zone       string    `json:"zone"`
	SubZone    string    `json:"sub-zone"`
	Host       string    `json:"host"`
	RouteName  string    `json:"route"`
	AgentId    string    `json:"agent-id"`
	SQLCommand string    `json:"sql-command"` // insert,update,delete
	CreatedTS  time.Time `json:"created-ts"`
	UpdatedTS  time.Time `json:"updated-ts"`
	Location   string    `json:"location"`
	Status     string    `json:"status"` // Scheduled,In-Progress,Completed,Failed
}

func (p RedirectConfig) Origin() core.Origin {
	return core.Origin{
		Region:  p.Region,
		Zone:    p.Zone,
		SubZone: p.SubZone,
		Host:    p.Host,
		Route:   p.RouteName,
	}
}
