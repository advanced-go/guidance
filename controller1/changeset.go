package controller1

import "github.com/advanced-go/stdlib/core"

type AuthorityChange struct {
	Name    string `json:"name"`
	Version string `json:"version"`
	Role    string `json:"role"`
}

type AuthorityChangeset struct {
	Insert []AuthorityChange `json:"insert"`
	Update []AuthorityChange `json:"update"`
	Delete []AuthorityChange `json:"delete"`
}

type EgressChange struct {
	RouteName    string `json:"route"`
	RateLimiting bool   `json:"rate-limiting"`
	RegionT      string `json:"region-t"`
	ZoneT        string `json:"zone-t"`
	SubZoneT     string `json:"sub-zone-t"`
	HostT        string `json:"host-t"`
	Authority    string `json:"authority"` // github/advanced-go/observation: provider/account/repository
	Version      string `json:"version"`
}

type EgressChangeset struct {
	Insert []EgressChange `json:"insert"`
	Update []EgressChange `json:"update"`
	Delete []EgressChange `json:"delete"`
}

type Changeset struct {
	Version   string             `json:"version"`
	Origin    core.Origin        `json:"origin"`
	Authority AuthorityChangeset `json:"authority-changeset"`
	Egress    EgressChangeset    `json:"egress-changeset"`
}
