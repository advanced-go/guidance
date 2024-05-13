package types

const (
	LocalityScope = "localityScope"
	None          = "none"
	Region        = "region"
	Zone          = "zone"
	Default       = "all"
)

const (
	NoneId          = 0
	DefaultId int16 = iota + 1
	RegionId        // AWS region, Data Center
	ZoneId          // AWS Zone, Cluster
	LocalityIdInvalid
)

/*
func (id SLOLocalityScope) String() string {
	if id < 0 || id >= LocalityIdInvalid {
		return ""
	}
	return [...]string{None, Region, Zone, Default}[id]
}

func (c SLOLocalityScope) Value(s string) int16 {
	switch s {
	case None:
		return NoneId
	case Region:
		return RegionId
	case Zone:
		return ZoneId
	case Default:
		return DefaultId
	}
	return 0
}

*/
