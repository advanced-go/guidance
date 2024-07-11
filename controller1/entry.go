package controller1

import "time"

// Where do the routing rules go??
// On the client or in the cloud?

var (
	//safeEntry = common.NewSafe()
	entryData = []Entry{
		{Region: "us-west1", Zone: "a", Host: "www.host1.com", Route: "search", CreatedTS: time.Date(2024, 6, 10, 7, 120, 35, 0, time.UTC)},
		{Region: "us-west1", Zone: "a", Host: "www.host2.com", Route: "search", CreatedTS: time.Date(2024, 6, 10, 7, 120, 35, 0, time.UTC)},
		{Region: "us-central1", Zone: "c", Host: "www.host1.com", Route: "search", CreatedTS: time.Date(2024, 6, 10, 7, 120, 35, 0, time.UTC)},
		{Region: "us-central1", Zone: "c", Host: "www.host2.com", Route: "search", CreatedTS: time.Date(2024, 6, 10, 7, 120, 35, 0, time.UTC)},
		{Region: "us-central1", Zone: "d", Host: "www.host4.com", Route: "search", CreatedTS: time.Date(2024, 6, 10, 7, 120, 35, 0, time.UTC)},
	}
)

type Entry struct {
	Region    string    `json:"region"`
	Zone      string    `json:"zone"`
	SubZone   string    `json:"sub-zone"`
	Host      string    `json:"host"`
	Route     string    `json:"route"`
	Primary   string    `json:"primary"`
	AgentId   string    `json:"agent-id"` // Needed for updating the primary on a conversion
	CreatedTS time.Time `json:"created-ts"`
	UpdatedTS time.Time `json:"created-ts"`

	// Default - starts processing can be empty as a default will be used.
	Threshold Threshold
}

func lastEntry() Entry {
	return entryData[len(entryData)-1]
}
