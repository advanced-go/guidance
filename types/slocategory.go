package types

const (
	Category              = "category"
	Latency               = "latency"
	StatusCodes           = "status-codes"
	Traffic               = "traffic"
	SaturationLatency     = "saturation-latency"
	SaturationUtilization = "saturation-utilization"
	Counter               = "counter"
	Profile               = "profile"
)

const (
	LatencyId int16 = iota + 1
	StatusCodesId
	TrafficId
	SaturationLatencyId
	SaturationUtilizationId
	CounterId
	ProfileId
	CategoryIdInvalid
)

/*
func (id int16) String() string {
	if id <= 0 || id >= CategoryIdInvalid {
		return ""
	}
	return [...]string{Latency, StatusCodes, Traffic, Saturation, Counter, Profile}[id-1]
}

*/

/*
func CategoryFromAny(a any) int16 {
	if util.IsNil(a) {
		return 0
	}
	if v, ok := a.(int32); ok {
		return CategoryFromId(v)
	}
	if v, ok := a.(string); ok {
		return CategoryFromString(v)
	}
	return 0
}

*/

func CategoryFromId(id int32) int16 {
	return int16(id)
}

func CategoryFromString(s string) int16 {
	switch s {
	case Latency:
		return LatencyId
	case StatusCodes:
		return StatusCodesId
	case Traffic:
		return TrafficId
	case SaturationLatency:
		return SaturationLatencyId
	case SaturationUtilization:
		return SaturationUtilizationId
	case Counter:
		return CounterId
	case Profile:
		return ProfileId
	}
	return 0
}
