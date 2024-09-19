package redirect1

import (
	"github.com/advanced-go/stdlib/core"
	"time"
)

// Note: ingress redirects are permanent and process based. No changes are allowed to the Redirect information,
// as the Redirect can be in progress. Terminating a running redirect can be achieved by adding a "terminated"
// status
//
// Note: CDC needs to be enabled for Entry and Status
//

// Entry - configuration for a permanent or temporary redirect
// Origin scope
// sub-zone - not possible as the location would be the current host
// zone     - allows redirect between sub-zones
// region   - allows redirect between zones
// global   - allows redirect between regions
type Entry struct {
	EntryId    int         `json:"entry-id"`
	RedirectId int         `json:"redirect-id"`
	Origin     core.Origin `json:"origin"`
	CreatedTS  time.Time   `json:"created-ts"`

	// 307 - temporary, 308 permanent
	StatusCode string `json:"status-code"`

	// Host selection
	Location string `json:"location"`
	Scope    string `json:"scope"` // zone, region, global

	// Processing configuration
	Policy Policy
}

// Policy - defines how the redirect is processed. Steps are the levels of traffic.
// Duration and deadline limit the amount of time.
// Note: there is no configuration to determine when to redirect on egress traffic. In other
// words, how long is spent on rate limiting vs redirect is left to experience.
type Policy struct {
	// Filter for traffic to be redirected
	Methods string            `json:"methods"` // List of comma seperated values, GET,PUT,POST, or * or empty
	Headers map[string]string `json:"headers"` // JSON list of header name value pairs

	// Traffic rollout
	StepRetries    int           `json:"step-retries"`    // Number of times to retry a step before failure
	StepThresholds string        `json:"step-thresholds"` // List of comma seperated percentages,a default of 10,20,40,70,100
	StepDuration   time.Duration `json:"step-duration"`   // How long to process a step

	// Time attributes for temporary redirect, start also applies for permanent
	StartTS    time.Time     `json:"start-ts"`    // redirect start time
	DeadlineTS time.Time     `json:"deadline-ts"` // redirect until a deadline
	Duration   time.Duration `json:"duration"`    // redirect for a duration
}

// CDCEntry - resiliency changes
type CDCEntry struct {
	CDCEntryId int `json:"host1-entry-id"`
	EntryId    int `json:"entry-id"`
	//Route      string    `json:"route"`
	CreatedTS time.Time `json:"created-ts"`
	//Resource   string    `json:"resource"` // header,ingress-routing,ingress-percentile,egress-routing
	SQLCommand string `json:"sql-command"` // database update,delete,insert
}

// Status - status changes to redirect only for permanent, temporary redirect status is in access log
type Status struct {
	RedirectId int         `json:"redirect-id"`
	StatusId   int         `json:"status-id"`
	Origin     core.Origin `json:"origin"`
	AgentId    string      `json:"agent-id"`
	CreatedTS  time.Time   `json:"created-ts"`
	Status     string      `json:"status"` // Scheduled,In-Progress,Completed,Failed,Terminated
}

// CDCStatus - resiliency changes
type CDCStatus struct {
	CDCEntryId int `json:"host1-entry-id"`
	EntryId    int `json:"entry-id"`
	//Route      string    `json:"route"`
	CreatedTS time.Time `json:"created-ts"`
	//Resource   string    `json:"resource"` // header,ingress-routing,ingress-percentile,egress-routing
	SQLCommand string `json:"sql-command"` // database update,delete,insert
}
