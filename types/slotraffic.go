package types

const (
	GRPCTrafficProtocol   = "gRPC"
	Http10TrafficProtocol = "HTTP10"
	Http11TrafficProtocol = "HTTP11"
	Http2TrafficProtocol  = "HTTP2"
	Http3TrafficProtocol  = "HTTP3"

	IngressTrafficType = "ingress"
	EgressTrafficType  = "egress"
	PingTrafficType    = "ping"
	ProfileTrafficType = "profile"
	CounterTrafficType = "counter"
)

const (
	Http11TrafficProtocolId int16 = iota + 1
	Http10TrafficProtocolId
	Http2TrafficProtocolId
	Http3TrafficProtocolId
	GRPCTrafficProtocolId
	TrafficProtocolInvalid
)

const (
	IngressTrafficId int16 = iota + 1
	EgressTrafficTypeId
	PingTrafficTypeId
	ProfileTrafficTypeId
	CounterTrafficTypeId
	TrafficTypeInvalid
)

/*
func (c SLOTrafficProtocol) String(id SLOTrafficProtocol) string {
	if id <= 0 || id >= TrafficProtocolInvalid {
		return ""
	}
	return [...]string{GRPCTraffic, Http10Traffic, Http11Traffic, Http2Traffic, Http3Traffic}[id-1]
}

func (c SLOTrafficProtocol) Value(s string) SLOTrafficProtocol {
	switch s {
	case GRPCTraffic:
		return GRPCTrafficId
	case Http10Traffic:
		return Http10TrafficId
	case Http11Traffic:
		return Http11TrafficId
	case Http2Traffic:
		return Http2TrafficId
	case Http3Traffic:
		return Http3TrafficId
	}
	return 0
}

*/
