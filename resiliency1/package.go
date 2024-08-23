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
func Put(r *http.Request, body []HostEntry) (http.Header, *core.Status) {
	if r == nil {
		return nil, core.NewStatusError(core.StatusInvalidArgument, errors.New("error: request is nil"))
	}
	if body == nil {
		content, status := json2.New[[]HostEntry](r.Body, r.Header)
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
func AddRateLimitingAction(ctx context.Context, origin core.Origin, action RateLimitingAction) *core.Status {
	return core.StatusOK()
}

// AddRoutingAction - add a routing action
func AddRoutingAction(ctx context.Context, origin core.Origin, action RoutingAction) *core.Status {
	return core.StatusOK()
}

// AddRedirectAction - add a redirect action
func AddRedirectAction(ctx context.Context, origin core.Origin, action RedirectAction) *core.Status {
	return core.StatusOK()
}

// SLO
// Percentile SLOs are system generated and only need to be retrieved on a scheduled basis.

// GetPercentileSLO - retrieve the percentile SLO for an origin
func GetPercentileSLO(ctx context.Context, origin core.Origin) (PercentileSLO, *core.Status) {
	var state PercentileSLO
	NewPercentileSLO(&state)
	return state, core.StatusOK()
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

// GetIngressRedirectState - retrieve the state needed to start the ingress redirect agent
func GetIngressRedirectState(ctx context.Context, origin core.Origin) (IngressRedirectState, *core.Status) {
	var state IngressRedirectState
	NewIngressRedirectState(&state)
	return state, core.StatusOK()
}

// GetIngressResiliencyState - retrieve the state needed to start the ingress resiliency agent
func GetIngressResiliencyState(ctx context.Context, origin core.Origin) (IngressResiliencyState, *core.Status) {
	var state IngressResiliencyState
	NewIngressResiliencyState(&state)
	return state, core.StatusOK()
}

// GetEgressState - retrieve the state needed to start the egress resiliency agent
func GetEgressState(ctx context.Context, origin core.Origin) ([]EgressState, *core.Status) {
	return []EgressState{}, core.StatusOK()
}

// CDC

// LastCDCId -
type LastCDCId struct {
	Entry    int
	Redirect int
	Failover int
}

// GetHostEntries - retrieve existing HostEntry
func GetHostEntries(ctx context.Context, origin core.Origin) ([]HostEntry, LastCDCId, *core.Status) {
	last := LastCDCId{
		Entry:    0,
		Redirect: 0,
		Failover: 0,
	}
	return []HostEntry{}, last, core.StatusOK()
}

// GetNewHostEntries - retrieve new HostEntry
func GetNewHostEntries(ctx context.Context, origin core.Origin, lastId int) ([]HostEntry, *core.Status) {
	return []HostEntry{}, core.StatusOK()
}

// GetUpdatedRedirectPlans - retrieve updated RedirectPlan
func GetUpdatedRedirectPlans(ctx context.Context, origin core.Origin, lastId int) ([]RedirectPlan, *core.Status) {
	return []RedirectPlan{}, core.StatusOK()
}

// GetUpdatedFailoverPlans - retrieve updated FailoverPlan
func GetUpdatedFailoverPlans(ctx context.Context, origin core.Origin, lastId int) ([]FailoverPlan, *core.Status) {
	return []FailoverPlan{}, core.StatusOK()
}
