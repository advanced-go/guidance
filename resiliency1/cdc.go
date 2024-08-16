package resiliency1

import "time"

// CDCEntry - resiliency changes
type CDCEntry struct {
	CDCEntryId int       `json:"cdc-entry-id"`
	EntryId    int       `json:"entry-id"`
	RouteName  string    `json:"route"`
	CreatedTS  time.Time `json:"created-ts"`
	Resource   string    `json:"resource"` // header,ingress-routing,ingress-percentile,egress-routing
	Action     string    `json:"action"`   // database update,delete,insert
}

// CDCRedirect - resiliency changes
type CDCRedirect struct {
	CDCEntryId int       `json:"cdc-entry-id"`
	EntryId    int       `json:"entry-id"`
	RouteName  string    `json:"route"`
	CreatedTS  time.Time `json:"created-ts"`
	Resource   string    `json:"resource"` // header,ingress-routing,ingress-percentile,egress-routing
	Action     string    `json:"action"`   // database update,delete,insert
}

// CDCFailover - resiliency changes
type CDCFailover struct {
	CDCEntryId int       `json:"cdc-entry-id"`
	EntryId    int       `json:"entry-id"`
	RouteName  string    `json:"route"`
	CreatedTS  time.Time `json:"created-ts"`
	Resource   string    `json:"resource"` // header,ingress-routing,ingress-percentile,egress-routing
	Action     string    `json:"action"`   // database update,delete,insert
}
