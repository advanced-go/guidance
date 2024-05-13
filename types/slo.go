package types

import (
	"errors"
	"fmt"
)

// SLO Attributes
// 1. Signal - latency,status code,saturation,traffic
// 2. Timeseries - ingress,egress,ping
// 3. Thresholds
//      Percent - latency, traffic, status codes
//      Value - latency, saturation, traffic
//      Minimum - status codes to attenuate underflow, applied to the window interval
// 4. Alerting - watch and warning percentage

//type SLOPercent uint8         // Range 1 - 99
//type SLOCategory uint8        // Latency, status codes, traffic, saturation, counter, profile
//type SLOLocalityScope uint8   // None, region, zone, all
//type SLOTrafficProtocol uint8 // gRPC, HTTP10, HTTP11, HTTP2, HTTP3
//type SLOInterval uint8        // Minutes
//type SLOInt uint16            // Value

const (
	//RPSComparisonPreviousDays - a value > 0
	RPSComparisonNone            = 0
	RPSComparisonCurrentLastWeek = -1
	RPSComparisonAllTime         = -2

	DefaultProcessingInterval = 5
	DefaultWindowInterval     = 7
	//DefaultProfileInterval    = 1
	//DefaultWindowIntervalMax = 60 minutes
)

// SLOEntry - contains the attributes to define and process an SLO.
type SLOEntry struct {
	// Unique identifiers
	Id         int32
	CustomerId int32

	// Category - latency, status codes, traffic, saturation-latency, saturation-utilization, counter, profile
	// TODO: Saturation SLOs need separate processing, less adaptive and more often. These are early warning
	Category int16

	// Traffic type - ingress, egress, ping, profile, counter
	TrafficType int16
	// Traffic protocol - gRPC, HTTP10, HTTP11, HTTP2, HTTP3
	TrafficProtocol int16

	// Interval range and window size
	// TODO: let adaptive determine optimal processing and window intervals
	//       using the configured value as being associated with the from interval
	//       The adaptive scaling needs to be logarithmic or exponential?
	//       The window interval should be larger than the from interval as this would catch issues
	//       that began at the end of a previous timeframe and ended at the beginning of the next window
	// TODO: RPS SLOs need an option for an interval of 1 day
	// TODO: Profile SLOs need to support granularity down to seconds
	ProcessingInterval int16 // Minutes
	WindowInterval     int16 // Minutes

	// Common parameters - applies to select/all categories
	WatchPercent     int16 // Range 1 - 99
	ThresholdPercent int16 // Used for latency, traffic, status codes, counter, profile
	ThresholdValue   int16 // Used for latency, saturation duration or traffic
	ThresholdMinimum int16 // Used for status codes to attenuate underflow, applied to the window interval

	// Traffic - this is current relative to a historical value
	// TODO : is previous days an average, or a comparison against each individual day?
	RPSLowComparison int16 // Values : None, Previous N Days, Same Day Last Week
	// The comparison for high would be a percentage of the metric
	RPSHighComparison int16 // Values : None, Previous N Days, Same Day Last Week, All-time

	// Processing flags - adaptive processing is enabled when an "ToIntervalMinutes" is configured
	// TODO : determine need for a use case comparing current RPS vs a static value
	LocalityScope     int16 // Values : None, Region, Zone, Default - uses all localities
	DisableProcessing bool  // Whether to process or not
	DisableTriage     bool

	// Status Codes
	// TODO : determine need for a use case comparing a ratio of status codes
	// TODO : add Envoy processing information to triage
	FilterStatusCodes string // Used for traffic related SLOs as a configurable filter
	StatusCodes       string // Comma seperated list of status codes, for a status code SLO

	// SLO name, unique within customer, service name and route
	Name        string
	Application string
	RouteName   string
}

func (s SLOEntry) IsEmpty() bool {
	return s.Id == 0
}

/*
func (s SLOEntry) IsAdaptiveProcessing() bool {
	return s.ToIntervalMinutes != 0
}

*/

func (s SLOEntry) IsRPSLowPreviousDays() bool {
	return s.RPSLowComparison > 0
}

func (s SLOEntry) IsRPSLowCurrentLastWeek() bool {
	return s.RPSLowComparison == RPSComparisonCurrentLastWeek
}

func (s SLOEntry) SetRPSLowPreviousDays(days int16) error {
	if days <= 0 {
		return errors.New(fmt.Sprintf("invalid number of days, days must be positive : %v", days))
	}
	s.RPSLowComparison = days
	return nil
}

func (s SLOEntry) SetRPSLowCurrentLastWeek() {
	s.RPSLowComparison = RPSComparisonCurrentLastWeek
}

func (s SLOEntry) IsRPSHighAllTimeTracking() bool {
	return s.Category == TrafficId && s.IsRPSHighAllTime()
}

func (s SLOEntry) IsRPSHighAllTime() bool {
	return s.RPSHighComparison == RPSComparisonAllTime
}

func (s SLOEntry) IsRPSHighPreviousDays() bool {
	return s.RPSHighComparison > 0
}

func (s SLOEntry) IsRPSHighCurrentLastWeek() bool {
	return s.RPSHighComparison == RPSComparisonCurrentLastWeek
}

func (s SLOEntry) SetRPSHighPreviousDays(days int16) error {
	if days <= 0 {
		return errors.New(fmt.Sprintf("invalid number of days, days must be positive : %v", days))
	}
	s.RPSHighComparison = days
	return nil
}

func (s SLOEntry) SetRPSHighCurrentLastWeek() {
	s.RPSHighComparison = RPSComparisonCurrentLastWeek
}

func (s SLOEntry) SetRPSHighAllTime() {
	s.RPSHighComparison = RPSComparisonAllTime
}

// SLOEntry - contains the attributes to define and process an SLO.
type SLOEntryAddenda struct {
	// Unique identifiers
	Id         int32
	SLOEntryId int32

	// Profile - metrics that function like a gauge
	// TODO : determine possible functions and required parameters
	// Can use saturation, percentages,ratio of peak
	MetricName string // Counter name/profile metric name, CPU, GO routines...

	// Counter
	// TODO : determine possible functions and required parameters
	// Aggregations applicable to counters: counter
	// Can also use ratio/percentages
	MetricNameSecondary string
	Function            string // Timeseries counter aggregation function selector
}
