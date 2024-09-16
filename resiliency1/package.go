package resiliency1

import (
	"context"
	"errors"
	"github.com/advanced-go/stdlib/core"
	json2 "github.com/advanced-go/stdlib/json"
	"net/http"
	"net/url"
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

// Get - resiliency1 GET
func Get(ctx context.Context, h http.Header, values url.Values) (entries []RedirectConfig, h2 http.Header, status *core.Status) {
	return entries, h, core.StatusOK()
}

func Delete(ctx context.Context, h http.Header, values url.Values) (h2 http.Header, status *core.Status) {
	return h, core.StatusOK()
}

// Put - resource PUT, with optional content override
func Put(r *http.Request, body []HostEntry) (http.Header, *core.Status) {
	if r == nil {
		return nil, core.NewStatusError(core.StatusInvalidArgument, errors.New("error: request is nil"))
	}
	if body == nil {
		content, status := json2.New[[]HostEntry](r.Body, r.Header)
		if !status.OK() {
			var e core.Log
			e.Handle(status.WithRequestId(r.Header))
			return nil, status
		}
		body = content
	}
	return nil, core.StatusOK() //put[core.Log](r.Context(), core.AddRequestId(r.Header), body)
}

// SLO
// Percentile SLOs are system generated and only need to be retrieved on a scheduled basis.

// GetPercentileSLO - retrieve the percentile SLO for an origin
func GetPercentileSLO(ctx context.Context, origin core.Origin) (PercentileSLO, *core.Status) {
	var state PercentileSLO
	NewPercentileSLO(&state)
	return state, core.StatusOK()
}

// Configuration

// UpdateRedirectConfig - update the ingress redirect configuration
func UpdateRedirectConfig(ctx context.Context, origin core.Origin, status string) *core.Status {
	return core.StatusOK()
}

// DeleteEgressConfig - delete the route egress config
func DeleteEgressConfig(ctx context.Context, origin core.Origin) *core.Status {
	return core.StatusOK()
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

// GetEgressState - retrieve the state needed to start the egress resiliency agents for a host
func GetEgressState(ctx context.Context, origin core.Origin) ([]EgressState, *core.Status) {
	return []EgressState{}, core.StatusOK()
}

// CDC

// GetHostEntries - retrieve existing HostEntry
func GetHostEntries(ctx context.Context, origin core.Origin) ([]HostEntry, LastCDCId, *core.Status) {
	last := LastCDCId{
		Entry:    0,
		Redirect: 0,
		Egress:   0,
	}
	return []HostEntry{}, last, core.StatusOK()
}

// GetNewHostEntries - retrieve new HostEntry
func GetNewHostEntries(ctx context.Context, origin core.Origin, lastId int) ([]HostEntry, *core.Status) {
	return []HostEntry{}, core.StatusOK()
}

// GetUpdatedRedirectConfigs - retrieve updated Redirect configurations
func GetUpdatedRedirectConfigs(ctx context.Context, origin core.Origin, lastId int) ([]RedirectConfig, *core.Status) {
	return []RedirectConfig{}, core.StatusOK()
}

// GetUpdatedEgressConfigs - retrieve updated Egress configurations
func GetUpdatedEgressConfigs(ctx context.Context, origin core.Origin, lastId int) ([]EgressConfig, *core.Status) {
	return []EgressConfig{}, core.StatusOK()
}
