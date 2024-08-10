package routing1

import "time"

// Redirect - routing config
// TODO : need to add CDC to this table for Agent data change notifications.
type Redirect struct {
	EntryId   int       `json:"entry-id"`
	RouteName string    `json:"route"`
	CreatedTS time.Time `json:"created-ts"`
	AgentId   string    `json:"agent-id"`
	Location  string    `json:"location"`
	Status    string    `json:"status"` // Scheduled,In-Progress,Completed,Failed
}
