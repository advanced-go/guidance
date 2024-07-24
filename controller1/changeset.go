package controller1

import "github.com/advanced-go/stdlib/core"

type ProcessingUpdate struct {
	ProcessingScheduleId any `json:"processing-schedule-id"`
	DependencyScheduleId any `json:"dependency-schedule-id"`
	DependencyUpdates    any `json:"dependency-updates"`
	Email                any `json:"email"`
	Slack                any `json:"slack"`
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
	RouteName        string `json:"route"`
	RedirectLocation string `json:"redirect-location"` // github/advanced-go/observation: provider/account/repository
}

type IngressChangeset struct {
	Insert []IngressChange `json:"insert"`
	Update []IngressChange `json:"update"`
	Delete []IngressChange `json:"delete"`
}

type EgressChange struct {
	RouteName         string `json:"route"`
	FailoverScope     string `json:"failover-scope"`
	FailoverThreshold int    `json:"failover-threshold"`
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
