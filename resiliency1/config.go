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
	Route      string    `json:"route"`
	AgentId    string    `json:"agent-id"`
	SQLCommand string    `json:"sql-command"` // insert,update,delete
	CreatedTS  time.Time `json:"created-ts"`

	// Failover configuration
	FailoverScope     string `json:"scope"`     // SubZone, Zone, Region, *, empty or none -> not configured
	FailoverThreshold int    `json:"threshold"` // Percentage of traffic
}

func (p EgressConfig) Origin() core.Origin {
	return core.Origin{
		Region:  p.Region,
		Zone:    p.Zone,
		SubZone: p.SubZone,
		Host:    p.Host,
		Route:   p.Route,
	}
}

// RedirectConfig - ingress redirect, can be permanent or temporary
// TODO : need to add CDC to this table for Agent data change notifications.
type RedirectConfig struct {
	EntryId    int       `json:"entry-id"`
	Region     string    `json:"region"`
	Zone       string    `json:"zone"`
	SubZone    string    `json:"sub-zone"`
	Host       string    `json:"host"`
	Route      string    `json:"route"`
	AgentId    string    `json:"agent-id"`
	SQLCommand string    `json:"sql-command"` // insert,update,delete
	CreatedTS  time.Time `json:"created-ts"`
	UpdatedTS  time.Time `json:"updated-ts"`

	// Redirection attributes.
	StatusCode int           `json:"status-code"` // 307 - temporary, 308 permanent
	Filter     string        `json:"filter"`      // JSON redirect filter for traffic, using request attributes
	Duration   time.Duration `json:"duration"`    // redirect for a duration
	DeadlineTS time.Time     `json:"deadline-ts"` // redirect until a deadline
	Location   string        `json:"location"`    // URL or just host name?
	Status     string        `json:"status"`      // Scheduled,In-Progress,Completed,Failed
}

func (p RedirectConfig) Origin() core.Origin {
	return core.Origin{
		Region:  p.Region,
		Zone:    p.Zone,
		SubZone: p.SubZone,
		Host:    p.Host,
		Route:   p.Route,
	}
}

// Redirect - configuration for a permanent or temporary redirect
type Redirect struct {
	StatusCode string `json:"status-code"` // 307 - temporary, 308 permanent

	// Host name and origin.
	Location RedirectLocation
	Filter   RedirectFilter
	Policy   RedirectPolicy
}

// RedirectLocation - configured redirect host name and origin
type RedirectLocation struct {
	Host string `json:"host"` // host name

	// Origin template to determine where a host is located.
	// If nothing is configured, then the current Origin is used.
	// Configurations act as filters/templates to select a particular host in a region
	RegionT  string `json:"region-t"`
	ZoneT    string `json:"zone-t"`
	SubZoneT string `json:"sub-zone-t"`
	HostT    string `json:"host-t"`
}

// RedirectFilter - defines how to select traffic to be redirected
type RedirectFilter struct {
	// Percentage - percentage of traffic, only allow multiples of 10: 10, 30, 50
	//Percentage int //

	// Or, request filtering, with methods and headers logical AND
	Methods string            // List of comma seperated values, GET,PUT,POST, or * or empty
	Headers map[string]string // JSON list of header name value pairs

}

// RedirectPolicy - defines how the redirect is processed. Steps are the levels of traffic.
// Duration and deadline limit the amount of time.
// Note: there is no configuration to determine when to redirect on egress traffic. In other
// words, how long is spent on rate limiting vs redirect is left to experience.
type RedirectPolicy struct {
	Steps string // List of comma seperated values,a default of 10,20,40,70,100

	Duration   time.Duration `json:"duration"`    // redirect for a duration
	DeadlineTS time.Time     `json:"deadline-ts"` // redirect until a deadline

}
