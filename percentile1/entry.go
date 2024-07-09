package percentile1

type Entry struct {
	Watch   int16 // Range 1 - 99
	Percent int16 // Used for latency, traffic, status codes, counter, profile
	Value   int16 // Used for latency, saturation duration or traffic
	Minimum int16 // Used for status codes to attenuate underflow, applied to the window interval
}
