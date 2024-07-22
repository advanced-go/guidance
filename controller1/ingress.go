package controller1

import "time"

// Issues, questions, and decisions.
// 1. How to do versioning??
// 2. How to manage change sets?
// 3. Can timeouts be configured locally and not in the cloud.
// 4. Hosts are immutable, any changes require a new host.
//    - Maybe mutability only applies to semantic version MAJOR-MINOR and or PATCH.
//    - System configurable to include MAJOR.MINOR and/or PATCH
//    - Dynamic routing would need to know about this decision

// IngressAuthority - ingress authorities
// Cardinality - n
// Access - Dynamic routing
// Update - User Changeset
type IngressAuthority struct {
	EntryId   int       `json:"entry-id"`   // How to refer to the main entry
	VersionId string    `json:"version-id"` // How to version this artifact
	CreatedTS time.Time `json:"created-ts"`
	AgentId   string    `json:"agent-id"` // Auditing

	Authority string `json:"authority"` // provider/account/repository
	Version   string `json:"version"`   // Semantic versioning

	// Primary or secondary. Primary would be hosted by the original author(s), and secondary would be hosted
	// by services that use the authority. This is needed when selecting a host, so that the primary is
	// the first selected. The primary should also be better able to handle higher traffic loads
	Role string `json:"role"`
}
