package controller1

import "time"

// Issues, questions, and decisions.
// If the hosts are immutable, then we do not need changesets? YES, still can be changes to host configuration
// How to manage/work changesets?
// Is a change set just new configuration entries?
// How to notify of new changesets??
// How to rollback changesets, and notifications of changes.
// Can rate limiting be a single configuration, possibly JSON?
// YES, as it would only be configured to turn rate limiting off.

// Control - host, utilize semantic versioning
type Control struct {
	EntryId   int       `json:"entry-id"`
	Region    string    `json:"region"`
	Zone      string    `json:"zone"`
	SubZone   string    `json:"sub-zone"`
	Host      string    `json:"host"`
	Version   string    `json:"version"` // Used to determine current changeset
	CreatedTS time.Time `json:"created-ts"`
	UpdatedTS time.Time `json:"updated-ts"`

	// Notifications
	Email string
	Slack string
}

// Changeset1 - really is the latest version table
type Changeset1 struct {
	EntryId   int
	VersionId string

	IngressRouteVersion     string
	IngressAuthorityVersion string
	EgressRouteVersion      string
}

// RateLimiting - rate limiting by Route
type RateLimiting struct {
	EntryId   int       `json:"entry-id"`   // How to refer to the main entry
	VersionId string    `json:"version-id"` // How to version this artifact
	CreatedTS time.Time `json:"created-ts"`
	AgentId   string    `json:"agent-id"` // Auditing

	Config string `json:"config"` // Should be JSON as NameValue pairs

}
