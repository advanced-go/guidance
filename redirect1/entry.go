package redirect1

import (
	"github.com/advanced-go/stdlib/core"
	"time"
)

const (
	DefaultStepDuration   = time.Minute * 6
	DefaultStepThresholds = "10,20,40,70,100"
)

// Note: ingress redirects are permanent and process based. No changes are allowed to the Redirect information,
// as the Redirect can be in progress. Terminating a running redirect can be achieved by adding a "terminated"
// status
//
// Note: CDC needs to be enabled for IngressEntry, EgressEntry and IngressStatus
//       No EgressStatus

// IngressEntry - configuration for a permanent or temporary redirect
// TODO: Determine if status codes should be configured to determine failure.
type IngressEntry struct {
	Origin core.Origin `json:"origin"`
	Status string      `json:"status"` // Current status

	// 307 - temporary, 308 permanent
	StatusCode string `json:"status-code"`

	// Host URL or template
	Location string `json:"location"`

	// Failure thresholds
	FailureStatusCodes string `json:"failure-status-codes"` // Comma seperated status codes, including templates
	FailureThreshold   int    `json:"failure-threshold"`    // Percentage of traffic

	// Traffic rollout  - optional, defaults: 10,20,40,70,100 / 6 minutes
	StepThresholds string        `json:"step-thresholds"` // List of comma seperated percentages
	StepDuration   time.Duration `json:"step-duration"`   // Step processing duration, must be >= 5 minutes

	// Time attributes - optional
	StartTS  time.Time     `json:"start-ts"` // start time
	Duration time.Duration `json:"duration"` // duration for a temporary redirect
}

func SetRolloutDefaults(e *IngressEntry) {
	if e.StepThresholds == "" {
		e.StepThresholds = DefaultStepThresholds
	}
	if e.StepDuration == 0 {
		e.StepDuration = DefaultStepDuration
	}
}

// EgressEntry - used for redirecting traffic when upstream host is saturating
//
// Scope - new host selection
//
//	sub-zone - not possible as the location would be the current host
//	zone     - allows redirect between sub-zones
//	region   - allows redirect between zones
//	global   - allows redirect between regions
//
// Threshold - percentage of failure traffic needed to trigger a redirect
//
//	value == -1  -> let system determine, first try rate limiting, if still failing, then redirect
//	value ==  0  -> no threshold, re-routing immediately when failures occur
//	value  >  0  -> re-routing when failure threshold is met
type EgressEntry struct {
	Origin core.Origin `json:"origin"`
	Status string      `json:"status"` // Current status

	Scope     string `json:"scope"`
	Threshold int    `json:"threshold"`
}

/*

// CDCIngressEntry - resiliency changes
type CDCIngressEntry struct {
	CDCEntryId int `json:"host1-entry-id"`
	EntryId    int `json:"entry-id"`
	//Route      string    `json:"route"`
	CreatedTS time.Time `json:"created-ts"`
	//Resource   string    `json:"resource"` // header,ingress-routing,ingress-percentile,egress-routing
	SQLCommand string `json:"sql-command"` // database update,delete,insert
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


*/

/*

// Status2 - status changes to redirect only for permanent, temporary redirect status is in access log
type Status2 struct {
	//RedirectId int         `json:"redirect-id"`
	//StatusId   int         `json:"status-id"`
	Origin    core.Origin `json:"origin"`
	AgentId   string      `json:"agent-id"`
	CreatedTS time.Time   `json:"created-ts"`
	Status    string      `json:"status"` // Scheduled,In-Progress,Completed,Failed,Terminated
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

*/
