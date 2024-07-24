package controller1

import "github.com/advanced-go/stdlib/core"

type IngressChange struct {
	RouteName        string `json:"route"`
	RedirectLocation any    `json:"redirect-location"` // github/advanced-go/observation: provider/account/repository
}

type IngressChangeset struct {
	Update []IngressChange `json:"update"`
}

type EgressChange struct {
	RouteName         string `json:"route"`
	FailoverScope     any    `json:"failover-scope"`
	FailoverThreshold any    `json:"failover-threshold"`
}

type EgressChangeset struct {
	Insert []EgressChange `json:"insert"`
	Update []EgressChange `json:"update"`
	Delete []EgressChange `json:"delete"`
}

type Changeset struct {
	Version string `json:"version"`
	// Do we need both? Can we only use a changeset id??
	ChangesetId string      `json:"changeset-id"`
	Origin      core.Origin `json:"origin"`
	//Processing  ProcessingUpdate   `json:"processing-update"`
	//Authority   AuthorityChangeset `json:"authority-changeset"`
	Ingress IngressChangeset `json:"ingress-changeset"`
	Egress  EgressChangeset  `json:"egress-changeset"`
}
