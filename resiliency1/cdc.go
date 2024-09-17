package resiliency1

import "time"

// LastCDCId -
type LastCDCId struct {
	Entry    int
	Redirect int
	Egress   int
}

// CDCEntry - resiliency changes
type CDCEntry struct {
	CDCEntryId int       `json:"host1-entry-id"`
	EntryId    int       `json:"entry-id"`
	Route      string    `json:"route"`
	CreatedTS  time.Time `json:"created-ts"`
	Resource   string    `json:"resource"` // header,ingress-routing,ingress-percentile,egress-routing
	Action     string    `json:"action"`   // database update,delete,insert
}

// CDCRedirect - resiliency changes
type CDCRedirect struct {
	CDCEntryId int       `json:"host1-entry-id"`
	EntryId    int       `json:"entry-id"`
	Route      string    `json:"route"`
	CreatedTS  time.Time `json:"created-ts"`
	Resource   string    `json:"resource"` // header,ingress-routing,ingress-percentile,egress-routing
	Action     string    `json:"action"`   // database update,delete,insert
}

// CDCFailover - resiliency changes
type CDCFailover struct {
	CDCEntryId int       `json:"host1-entry-id"`
	EntryId    int       `json:"entry-id"`
	Route      string    `json:"route"`
	CreatedTS  time.Time `json:"created-ts"`
	Resource   string    `json:"resource"` // header,ingress-routing,ingress-percentile,egress-routing
	Action     string    `json:"action"`   // database update,delete,insert
}
