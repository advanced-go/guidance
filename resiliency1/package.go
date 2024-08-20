package resiliency1

import (
	"context"
	"errors"
	"github.com/advanced-go/stdlib/core"
	json2 "github.com/advanced-go/stdlib/json"
	"net/http"
	"time"
)

const (
	PkgPath                   = "github/advanced-go/guidance/resiliency1"
	PercentilePollingDuration = time.Hour * 12
	RedirectStatusScheduled   = "scheduled"
	RedirectStatusInProgress  = "in-progress"
	RedirectStatusSucceeded   = "succeeded"
	RedirectStatusFailed      = "failed"
)

// Really only need a put

// Put - resource PUT, with optional content override
func Put(r *http.Request, body []Entry) (http.Header, *core.Status) {
	if r == nil {
		return nil, core.NewStatusError(core.StatusInvalidArgument, errors.New("error: request is nil"))
	}
	if body == nil {
		content, status := json2.New[[]Entry](r.Body, r.Header)
		if !status.OK() {
			var e core.Log
			e.Handle(status, core.RequestId(r.Header))
			return nil, status
		}
		body = content
	}
	return nil, core.StatusOK() //put[core.Log](r.Context(), core.AddRequestId(r.Header), body)
}

// Actions - append only to maintain audit trail with inference

// GetActions - retrieve the latest ingress and egress actions for all of an origin's routes. This is called on
// host startup.
func GetActions(ctx context.Context, origin core.Origin) (Actions, *core.Status) {
	return Actions{}, core.StatusOK()
}

// AddRateLimitingAction - add a rate limiting action
func AddRateLimitingAction(ctx context.Context, origin core.Origin, action *RateLimitingAction) *core.Status {
	return core.StatusOK()
}

// AddRoutingAction - add a routing action
func AddRoutingAction(ctx context.Context, origin core.Origin, action *RoutingAction) *core.Status {
	return core.StatusOK()
}

// AddRedirectAction - add a redirect action
func AddRedirectAction(ctx context.Context, origin core.Origin, action *RedirectAction) *core.Status {
	return core.StatusOK()
}

// SLO
// Percentile SLOs are system generated and only need to be retrieved on a scheduled basis.

// GetPercentileSLO - retrieve the percentile SLO for an origin
func GetPercentileSLO(ctx context.Context, origin core.Origin) (PercentileSLO, *core.Status) {
	return PercentileSLO{}, core.StatusOK()
}

// Plans
// Redirect update only as the previous redirect is not needed??
// Failover update only

// GetRedirectPlan - retrieve the ingress redirect plan
func GetRedirectPlan(ctx context.Context, origin core.Origin) (RedirectPlan, *core.Status) {
	return RedirectPlan{}, core.StatusOK()
}

// UpdateRedirectPlan - update the ingress redirect plan
func UpdateRedirectPlan(ctx context.Context, origin core.Origin, status string) *core.Status {
	return core.StatusOK()
}

// GetFailoverPlan - retrieve the route egress plan
func GetFailoverPlan(ctx context.Context, origin core.Origin) ([]FailoverPlan, *core.Status) {
	return []FailoverPlan{}, core.StatusOK()
}

// State

// NewIngressRedirectState - initialize
func NewIngressRedirectState() *IngressRedirectState {
	s := new(IngressRedirectState)
	s.Percent = DefaultPercentileSLO.Percent
	s.Latency = DefaultPercentileSLO.Latency
	s.Minimum = DefaultPercentileSLO.Minimum
	s.Percentage = -1
	return s
}

// GetIngressRedirectState - retrieve the state needed to start the ingress redirect agent
func GetIngressRedirectState(ctx context.Context, origin core.Origin) (*IngressRedirectState, *core.Status) {
	return NewIngressRedirectState(), core.StatusOK()
}

// NewIngressResiliencyState - initialize
func NewIngressResiliencyState() *IngressResiliencyState {
	s := new(IngressResiliencyState)
	s.Percent = DefaultPercentileSLO.Percent
	s.Latency = DefaultPercentileSLO.Latency
	s.Minimum = DefaultPercentileSLO.Minimum
	return s
}

// GetIngressResiliencyState - retrieve the state needed to start the ingress resiliency agent
func GetIngressResiliencyState(ctx context.Context, origin core.Origin) (*IngressResiliencyState, *core.Status) {
	return NewIngressResiliencyState(), core.StatusOK()
}

// GetEgressState - retrieve the state needed to start the egress resiliency agent
func GetEgressState(ctx context.Context, origin core.Origin) ([]EgressState, *core.Status) {
	return []EgressState{}, core.StatusOK()
}

// CDC

// GetAssignments - retrieve existing Entry
func GetAssignments(ctx context.Context, origin core.Origin) ([]Entry, int, *core.Status) {
	lastId := 0
	return []Entry{}, lastId, core.StatusOK()
}

// GetNewAssignments - retrieve new Entry
func GetNewAssignments(ctx context.Context, origin core.Origin, lastId int) ([]Entry, int, *core.Status) {
	return []Entry{}, lastId, core.StatusOK()
}

// GetUpdatedRedirectPlans - retrieve updated RedirectPlan
func GetUpdatedRedirectPlans(ctx context.Context, origin core.Origin, lastId int) ([]RedirectPlan, int, *core.Status) {
	return []RedirectPlan{}, lastId, core.StatusOK()
}

// GetUpdatedFailoverPlans - retrieve updated FailoverPlan
func GetUpdatedFailoverPlans(ctx context.Context, origin core.Origin, lastId int) ([]FailoverPlan, int, *core.Status) {
	return []FailoverPlan{}, lastId, core.StatusOK()
}
