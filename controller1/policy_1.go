package controller1

import (
	"time"
)

// Policy design:
//
// Need a default policy, plus optional failover policy, and additional parallel/canary and conversion
//
// How to effective date failover policies so that a smooth transition can occur?
// Or can we just add a new one and test? Need a way to test a new failover policy.
//
// Activation model :
// 1. Effective data range
// 2. Threshold - percentage of filtered traffic
//
// Action - what processing does the agent do
// 1. Conversion - converting hosts
// 2. Failover - a list of secondary hosts
//
//
// Can have only 1 failover policy, which is threshold activated
// Multiple conversion or parallel policies
// How to have a permanent multiple host routing? Does that make sense with Kubernetes

var (
	//safeStatus = common.NewSafe()
	policyData = []EntryPolicy{
		{EntryId: 1, Status: "active", CreatedTS: time.Date(2024, 6, 10, 7, 120, 35, 0, time.UTC)},
		{EntryId: 1, Status: "inactive", CreatedTS: time.Date(2024, 6, 10, 7, 120, 35, 0, time.UTC)},
	}
)

func lastPolicy() EntryPolicy {
	return policyData[len(policyData)-1]
}

type Threshold struct {
	Percent int    `json:"percent"`
	Codes   string `json:"codes"`
}

type EffectiveDate struct {
	FromTS time.Time `json:"from-ts"`
	ToTS   time.Time `json:"to-ts"`
}

// Conversion - need to change primary after completion
type Conversion struct {
	Timespan     EffectiveDate
	Secondary    string        `json:"secondary"`
	StepPercent  string        `json:"step-percent"` // 10,20,100
	StepDuration time.Duration `json:"step-duration"`
	Threshold    Threshold     // Applied only during conversion
}

// Parallel - parallel for testing? Is this really a Canary?
type Parallel struct {
	Timespan  EffectiveDate
	Secondary string `json:"secondary"` // List of secondaries, percentage would be equal
	//StepPercent string    `json:"step-percent"` // 10,20,100
	//StepDuration time.Duration `json:"step-duration"`
	Threshold Threshold // Applied only during conversion
}

// Failover - for processing, need to determine how long to failover before failing back.
// Can we monitor route egress traffic?
// the wildcard should try close Zones, SubZones before a new region.
// May need some sort of cost metric
type Failover struct {
	Activation Threshold
	Secondary  string `json:"secondary"` // A specific host, or a wildcard to let Backbone decide, or list
	//StepPercent string    `json:"step-percent"` // 10,20,100
	//StepDuration time.Duration `json:"step-duration"`
	Threshold Threshold // Applied only during conversion
}

// EntryPolicy - policy
type EntryPolicy struct {
	EntryId   int       `json:"entry-id"`
	Status    string    `json:"status"` // Active, inactive, expired
	CreatedTS time.Time `json:"created-ts"`

	// Activation
	Range EffectiveDate

	// Processing
	Conversion Conversion
}
