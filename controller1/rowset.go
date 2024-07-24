package controller1

import "time"

type Rowset struct {
	EntryId   int       `json:"entry-id"`
	Region    string    `json:"region"`
	Zone      string    `json:"zone"`
	SubZone   string    `json:"sub-zone"`
	Host      string    `json:"host"`
	CreatedTS time.Time `json:"created-ts"`
	Status    string    `json:"status"` // Active, inactive, deleted?
	//UpdatedTS time.Time `json:"updated-ts"` this is in CDC

	// Current version - auditing via CDC
	Version string `json:"version"`
}
