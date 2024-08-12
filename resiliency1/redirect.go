package resiliency1

import "time"

// Redirect - routing config
// TODO : need to add CDC to this table for Agent data change notifications.
type Redirect struct {
	EntryId   int       `json:"entry-id"`
	RouteName string    `json:"route"`
	AgentId   string    `json:"agent-id"`
	CreatedTS time.Time `json:"created-ts"`
	Location  string    `json:"location"`
	Status    string    `json:"status"` // Scheduled,In-Progress,Completed,Failed
}
