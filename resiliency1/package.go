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

func AddRedirectStatus(ctx context.Context, origin core.Origin, status string) *core.Status {
	return core.StatusOK()
}

func IngressPercentile(ctx context.Context, origin core.Origin) (Percentile, *core.Status) {
	return Percentile{}, core.StatusOK()
}

func IngressRedirect(ctx context.Context, origin core.Origin) (Redirect, *core.Status) {
	return Redirect{}, core.StatusOK()
}

func IngressCDC(ctx context.Context, origin core.Origin) ([]CDCEntry, *core.Status) {
	return []CDCEntry{}, core.StatusOK()
}

func EgressFailover(ctx context.Context, origin core.Origin) ([]Failover, *core.Status) {
	return []Failover{}, core.StatusOK()
}

func EgressCDC(ctx context.Context, origin core.Origin) ([]CDCEntry, *core.Status) {
	return []CDCEntry{}, core.StatusOK()
}
