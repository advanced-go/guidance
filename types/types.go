package types

import "github.com/advanced-go/stdlib/core"

// https://localhost:8081/github/advanced-go/guidance:region/zone/subzone/app
// guidance:us-west.zone1.sub-zone.app
// guidance:us-east.zone2..app

type CostFunction struct {
	Threshold string `json:"threshold"`
}

type EntryV1 struct {
	Origin    core.Origin
	Status    string `json:"status"`
	CreatedTS string `json:"created-ts"`
	UpdatedTS string `json:"updated-ts"`

	CostFunction CostFunction

	// Routing
	PrimaryRoute   string `json:"primary-route"`
	SecondaryRoute string `json:"secondary-route"`
	SecondaryPct   string `json:"secondary-pct"`

	// Timeout
	Timeout string `json:"timeout"`

	// Rate Limiting
	RateLimit string `json:"rate-limit"`
	RateBurst string `json:"rate-burst"`
}
