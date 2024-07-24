package controller1

import "time"

// Issues, questions, and decisions.

type IngressController struct {
	EntryId   int       `json:"entry-id"`   // How to refer to the main entry
	VersionId string    `json:"version-id"` // How to version this artifact
	CreatedTS time.Time `json:"created-ts"`
	AgentId   string    `json:"agent-id"` // Auditing

	// How to do redirects?? Send a 308 with a location header
	RedirectLocation string `json:"redirect-location"`
	// Where to redirect to. How to determine origin. Should be same origin
	// Step percentage determined by cloud.

}
