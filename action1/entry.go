package action1

import (
	"time"
)

const (
	//accessLogSelect = "SELECT * FROM access_log {where} order by start_time limit 2"
	accessLogSelect = "SELECT region,customer_id,start_time,duration_str,traffic,rate_limit FROM access_log {where} order by start_time desc limit 2"
	accessLogInsert = "INSERT INTO access_log (" +
		"customer_id,start_time,duration_ms,duration_str,traffic," +
		"region,zone,sub_zone,service,instance_id,route_name," +
		"request_id,url,protocol,method,host,path,status_code,bytes_sent,status_flags," +
		"timeout,rate_limit,rate_burst) VALUES"

	EntryIdName   = "entry_id"
	AgentIdName   = "agent_id"
	CreatedTSName = "created_ts"
	RegionName    = "region"
	ZoneName      = "zone"
	SubZoneName   = "sub_zone"
	HostName      = "host"
	RouteName     = "route"
	DetailsName   = "details"
	ActionName    = "action"
)

var (
	//safeEntry = common.NewSafe()
	entryData = []Entry{
		{Region: "us-west1", Zone: "a", Host: "www.host1.com", AgentId: "agent-id", RouteName: "host", Details: "information", Action: "processed", CreatedTS: time.Date(2024, 6, 10, 7, 120, 35, 0, time.UTC)},
		{Region: "us-west1", Zone: "a", Host: "www.host2.com", AgentId: "agent-id", RouteName: "host", Details: "text", Action: "processed", CreatedTS: time.Date(2024, 6, 10, 7, 120, 35, 0, time.UTC)},
	}
)

// Entry - host
type Entry struct {
	Region    string `json:"region"`
	Zone      string `json:"zone"`
	SubZone   string `json:"sub-zone"`
	Host      string `json:"host"`
	RouteName string `json:"route"`

	CreatedTS time.Time `json:"created-ts"`
	AgentId   string    `json:"agent-id"`

	// RateLimiting
	Limit float64
	Burst int

	// Routing
	Primary    string
	Secondary  string // Location for ingress
	Percentage int
	//Code       string

	// Redirect

	// Details + action
	Details string `json:"details"`
	Action  string `json:"action"`
}

type RateLimitingEntry struct {
	Region    string    `json:"region"`
	Zone      string    `json:"zone"`
	SubZone   string    `json:"sub-zone"`
	Host      string    `json:"host"`
	RouteName string    `json:"route"`
	CreatedTS time.Time `json:"created-ts"`
	AgentId   string    `json:"agent-id"`

	// Need to represent 2 states:
	// 1. Nil or not configured - both values == -1
	// 2. Configured - both values >= 0
	Limit float64 `json:"limit"`
	Burst int     `json:"burst"`
}

type RoutingEntry struct {
	Region    string    `json:"region"`
	Zone      string    `json:"zone"`
	SubZone   string    `json:"sub-zone"`
	Host      string    `json:"host"`
	RouteName string    `json:"route"`
	CreatedTS time.Time `json:"created-ts"`
	AgentId   string    `json:"agent-id"`

	Location   string `json:"location"`
	Percentage int    `json:"percentage"`
	// Need to determine how to represent 3 states:
	// 1. Nil or not configured - location - empty, percentage = -1, code = empty
	// 2. Re-routing in progress
	// 3. Permanent redirect
	//Code       string `json:"code"` // Can this be used as "RD" for pre redirect, and the
}

// RedirectEntry - applies for both ingress and egress
type RedirectEntry struct {
	Region    string    `json:"region"`
	Zone      string    `json:"zone"`
	SubZone   string    `json:"sub-zone"`
	Host      string    `json:"host"`
	RouteName string    `json:"route"`
	CreatedTS time.Time `json:"created-ts"`
	AgentId   string    `json:"agent-id"`

	Location   string `json:"location"`
	StatusCode string `json:"status-code"` // Only for ingress
}
