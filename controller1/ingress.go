package controller1

import "time"

// Issues, questions, and decisions.

type Ingress struct {
	EntryId   int       `json:"entry-id"`   // How to refer to the main entry
	VersionId string    `json:"version-id"` // How to version this artifact
	CreatedTS time.Time `json:"created-ts"`
	AgentId   string    `json:"agent-id"` // Auditing

	// How to do redirects?? Send a 308 with a location header
	Location string `json:"location"` // Redirect location
	// Where to redirect to. How to determine origin. Should be same origin
	// Step percentage determined by cloud.
	// Need this status to determine on client startup whether all traffic should be redirected.
	Status string `json:"status"` // Completed, failed, may be processing??

}

func (i Ingress) IsEmpty() bool {
	return i.EntryId <= 0
}

func (i Ingress) IsRedirect() bool {
	return i.Location != ""
}

func (i Ingress) InProcess() bool {
	return i.IsRedirect() && i.Status == ""
}
