package schedule1

import "time"

type Entry struct {
	Region    string    `json:"region"`
	Zone      string    `json:"zone"`
	SubZone   string    `json:"sub-zone"`
	Host      string    `json:"host"`
	AgentId   string    `json:"agent-id"`
	CreatedTS time.Time `json:"created-ts"`

	Start int // Used for latency, traffic, status codes, counter, profile
	Stop  int // Used for latency, saturation duration or traffic

}

var (
	//safeEntry = common.NewSafe()
	entryData = []Entry{
		{Region: "us-west1", Zone: "a", Host: "www.host1.com", AgentId: "test-agent", Start: 9, Stop: 11, CreatedTS: time.Date(2024, 6, 10, 7, 120, 35, 0, time.UTC)},
		{Region: "us-west1", Zone: "a", Host: "www.host2.com", AgentId: "test-agent", Start: 9, Stop: 11, CreatedTS: time.Date(2024, 6, 10, 7, 120, 35, 0, time.UTC)},
		{Region: "us-central1", Zone: "c", Host: "www.host1.com", AgentId: "test-agent", Start: 9, Stop: 11, CreatedTS: time.Date(2024, 6, 10, 7, 120, 35, 0, time.UTC)},
		{Region: "us-central1", Zone: "c", Host: "www.host2.com", AgentId: "test-agent", Start: 9, Stop: 11, CreatedTS: time.Date(2024, 6, 10, 7, 120, 35, 0, time.UTC)},
		{Region: "us-central1", Zone: "d", Host: "www.host4.com", AgentId: "test-agent", Start: 9, Stop: 11, CreatedTS: time.Date(2024, 6, 10, 7, 120, 35, 0, time.UTC)},
	}
)
