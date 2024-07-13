package controller1

import (
	"github.com/advanced-go/stdlib/core"
	"time"
)

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
	Region    string      `json:"region"`
	Zone      string      `json:"zone"`
	SubZone   string      `json:"sub-zone"`
	Host      string      `json:"host"`
	Route     string      `json:"route"`
	CreatedTS time.Time   `json:"created-ts"`
	BeginTS   time.Time   `json:"begin-ts"`
	EndTS     time.Time   `json:"end-ts"`
	Include   core.Origin `json:"include"`
	Exclude   core.Origin `json:"exclude"`
}

type EntryStatus struct {
	Region    string    `json:"region"`
	Zone      string    `json:"zone"`
	SubZone   string    `json:"sub-zone"`
	Host      string    `json:"host"`
	Route     string    `json:"route"`
	Status    string    `json:"status"` // Values: Scheduled,Active,Completed
	CreatedTS time.Time `json:"created-ts"`
}

func lastEntry() Entry {
	return entryData[len(entryData)-1]
}
