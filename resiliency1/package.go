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

// Actions

// GetActions - retrieve the actions for an origin, containing a route name
func GetActions(ctx context.Context, origin core.Origin) (Actions, *core.Status) {
	return Actions{}, core.StatusOK()
}

// UpdateRateLimitingAction - insert/update a rate limiting action
func UpdateRateLimitingAction(ctx context.Context, origin core.Origin, action *RateLimitingAction) *core.Status {
	return core.StatusOK()
}

// UpdateRoutingAction - insert/update a routing action
func UpdateRoutingAction(ctx context.Context, origin core.Origin, action *RoutingAction) *core.Status {
	return core.StatusOK()
}

// UpdateRedirectAction - insert/update a redirect action
func UpdateRedirectAction(ctx context.Context, origin core.Origin, action *RedirectAction) *core.Status {
	return core.StatusOK()
}

// Plans

// GetPercentileSLO - retrieve the percentile SLO for an origin
func GetPercentileSLO(ctx context.Context, origin core.Origin) (PercentileSLO, *core.Status) {
	return PercentileSLO{}, core.StatusOK()
}

func GetRedirectPlan(ctx context.Context, origin core.Origin) (RedirectPlan, *core.Status) {
	return RedirectPlan{}, core.StatusOK()
}

func UpdateRedirectPlan(ctx context.Context, origin core.Origin, status string) *core.Status {
	return core.StatusOK()
}

func GetFailoverPlan(ctx context.Context, origin core.Origin) ([]FailoverPlan, *core.Status) {
	return []FailoverPlan{}, core.StatusOK()
}

// CDC

// GetEntryCDC - retrieve Entry CDC
func GetEntryCDC(ctx context.Context, origin core.Origin) ([]CDCEntry, *core.Status) {
	return []CDCEntry{}, core.StatusOK()
}

// GetRedirectCDC - retrieve RedirectPlan CDC
func GetRedirectCDC(ctx context.Context, origin core.Origin) ([]CDCRedirect, *core.Status) {
	return []CDCRedirect{}, core.StatusOK()
}

// GetFailoverCDC - retrieve FailoverPlan CDC
func GetFailoverCDC(ctx context.Context, origin core.Origin) ([]CDCFailover, *core.Status) {
	return []CDCFailover{}, core.StatusOK()
}

/*

func IngressCDC(ctx context.Context, origin core.Origin) ([]CDCEntry, *core.Status) {
	return []CDCEntry{}, core.StatusOK()
}

func IngressAssignment(ctx context.Context, origin core.Origin, ts time.Time, changesOnly bool) ([]Assignment, *core.Status) {
	return []Assignment{}, core.StatusOK()
}



func EgressCDC(ctx context.Context, origin core.Origin) ([]CDCEntry, *core.Status) {
	return []CDCEntry{}, core.StatusOK()
}

func EgressAssignment(ctx context.Context, origin core.Origin, ts time.Time, changesOnly bool) ([]Assignment, *core.Status) {
	return []Assignment{}, core.StatusOK()
}

func EntryQuery(ctx context.Context, origin core.Origin) ([]Entry, *core.Status) {
	return []Entry{}, core.StatusOK()
}


*/
