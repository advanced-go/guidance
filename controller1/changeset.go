package controller1

import "github.com/advanced-go/stdlib/core"

type DependencyUpdateChange struct {
	Enable bool `json:"enable"`
}

type DependencyUpdateChangeset struct {
	Update []DependencyUpdateChange `json:"update"`
}

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

type IngressChange struct {
	RouteName    string `json:"route"`
	RateLimiting bool   `json:"rate-limiting"`
}

type IngressChangeset struct {
	//Insert []IngressChange `json:"insert""`
	Update []IngressChange `json:"update"`
	//Delete []IngressChange `json:"delete"`
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
	Version          string                    `json:"version"`
	Origin           core.Origin               `json:"origin"`
	DependencyUpdate DependencyUpdateChangeset `json:"dependency-update-changeset"`
	Authority        AuthorityChangeset        `json:"authority-changeset"`
	Ingress          IngressChangeset          `json:"ingress-changeset"`
	Egress           EgressChangeset           `json:"egress-changeset"`
}
