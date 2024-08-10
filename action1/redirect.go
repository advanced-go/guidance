package action1

import "time"

// Redirect - ingress and egress
type Redirect struct {
	EntryId     int       `json:"entry-id"`
	RouteName   string    `json:"route"`
	CreatedTS   time.Time `json:"created-ts"`
	InferenceId int       `json:"inference-id"`
	Location    string    `json:"location"`
	StatusCode  string    `json:"status-code"` // Only for ingress
}
