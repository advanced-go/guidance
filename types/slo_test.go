package types

import (
	"encoding/json"
	"fmt"
	"unsafe"
)

// SLOEntry Sizeof : 120
func _ExampleSLOEntrySize() {
	// 36 bytes  = integers
	// 104 bytes + 4 byte alignment + TrafficStatusCodes + SLOName + Application + RouteName
	// 168 bytes + StatusCodes + MetricName + MetricNameSecondary + Function
	//strings = 36 bytes, 16 bytes per string + blocking to 64 bit address
	// 88 bytes = integers + status codes 2 strings + metric name
	var s SLOEntry

	fmt.Printf("SLOEntry Sizeof : %v\n", unsafe.Sizeof(s))

	//Output:
	// fail

}

func ExampleSLOEntryMarshal() {

	entry := SLOEntry{
		Id:              1001,
		CustomerId:      99,
		Category:        LatencyId,
		TrafficType:     IngressTrafficId,
		TrafficProtocol: Http11TrafficProtocolId,

		ProcessingInterval: DefaultProcessingInterval,
		WindowInterval:     DefaultWindowInterval,
		WatchPercent:       90,
		ThresholdPercent:   95,
		ThresholdValue:     500,

		DisableProcessing: false,
		DisableTriage:     false,
		LocalityScope:     DefaultId,

		RPSLowComparison:  RPSComparisonNone,
		RPSHighComparison: RPSComparisonNone,
		FilterStatusCodes: "",
		StatusCodes:       "",

		Name:        "lat-95-500",
		Application: "domain.sub-domain",
		RouteName:   "upstream-origin",
	}

	buf, err := json.Marshal(&entry)
	fmt.Printf("Error : %v\n", err)
	fmt.Printf("Entry  : %v\n", string(buf))

	//Output:
	// fail
}
