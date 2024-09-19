package redirect1

import (
	"context"
	"github.com/advanced-go/stdlib/core"
	"time"
)

const (
	PkgPath                   = "github/advanced-go/guidance/redirect1"
	PercentilePollingDuration = time.Hour * 12
	RedirectStatusScheduled   = "scheduled"
	RedirectStatusInProgress  = "in-progress"
	RedirectStatusSucceeded   = "succeeded"
	RedirectStatusFailed      = "failed"
	RedirectStatusTerminated  = "terminated"
	RedirectStatusActive      = "active"
	RedirectStatusInactive    = "inactive"
)

// Ingress CaseOfficer functions for the following:
//   Startup - create and run Redirect agents
//   New - check for new Redirects after startup
//   Terminated Status - notify a Redirect agent if a redirect has been terminated by a user

// QueryIngressOpen - find all open redirects, based on status. Used on case officer startup to create
// the necessary Redirect agents
//func QueryIngressOpen(ctx context.Context, origin core.Origin) ([]Entry, *core.Status) {
//	return []Entry{}, core.StatusOK()
//}

// QueryIngressNew - find new redirects
func QueryIngressNew(ctx context.Context, origin core.Origin, lastCDCId int) ([]core.Origin, *core.Status) {
	return []core.Origin{}, core.StatusOK()
}

// QueryIngressInactive - find terminated redirects
func QueryIngressInactive(ctx context.Context, origin core.Origin, lastCDCId int) ([]core.Origin, *core.Status) {
	return []core.Origin{}, core.StatusOK()
}

// GetIngress - retrieve an ingress redirect
func GetIngress(ctx context.Context, origin core.Origin) (Entry, *core.Status) {
	return Entry{}, core.StatusOK()
}

// AddIngressStatus - add a status
func AddIngressStatus(ctx context.Context, origin core.Origin, status string) *core.Status {
	return core.StatusOK()
}

// Egress CaseOfficer functions for the following:
//   New - check for new Redirects after startup, notify Egress agent
//   Status - notify an Egress agent if a redirect has been activated/de-activated

// QueryEgressNew - find new redirects
func QueryEgressNew(ctx context.Context, origin core.Origin, lastCDCId int) ([]core.Origin, *core.Status) {
	return []core.Origin{}, core.StatusOK()
}

// QueryEgressInactive - find inactive redirects
func QueryEgressInactive(ctx context.Context, origin core.Origin, lastCDCId int) ([]core.Origin, *core.Status) {
	return []core.Origin{}, core.StatusOK()
}

// GetEgress - retrieve an egress redirect
func GetEgress(ctx context.Context, origin core.Origin) (Entry, *core.Status) {
	return Entry{}, core.StatusOK()
}

// GetHostEgress - retrieve all egress redirects for a host, selecting on the active parameter
func GetHostEgress(ctx context.Context, origin core.Origin) ([]Entry, *core.Status) {
	return []Entry{}, core.StatusOK()
}

// AddEgressStatus - add a status
func AddEgressStatus(ctx context.Context, origin core.Origin, status string) *core.Status {
	return core.StatusOK()
}

/*
// GetIngressRedirect - get ingress redirect
func GetIngressRedirect(ctx context.Context, origin core.Origin) ([]RedirectConfig, *core.Status) {
	return []RedirectConfig{}, core.StatusOK()
}

// GetUpdatedIngressRedirect - retrieve updated Ingress redirect configurations
func GetUpdatedIngressRedirect(ctx context.Context, origin core.Origin, lastId int) ([]RedirectConfig, *core.Status) {
	return []RedirectConfig{}, core.StatusOK()
}

// AddIngressRedirectStatus - add an Ingress redirect status
func AddIngressRedirectStatus(ctx context.Context, origin core.Origin, status string) *core.Status {
	return core.StatusOK()
}

// Egress

// GetEgressRedirect - retrieve egress redirect
func GetEgressRedirect(ctx context.Context, origin core.Origin) ([]RedirectConfig, *core.Status) {
	return []RedirectConfig{}, core.StatusOK()
}

// GetUpdatedEgressRedirect - retrieve updated Egress redirect configurations
func GetUpdatedEgressRedirect(ctx context.Context, origin core.Origin, lastId int) ([]RedirectConfig, *core.Status) {
	return []RedirectConfig{}, core.StatusOK()
}

// UpdateRedirectConfig - update the ingress redirect configuration
func UpdateRedirectConfig(ctx context.Context, origin core.Origin, status string) *core.Status {
	return core.StatusOK()
}

// DeleteEgressConfig - delete the route egress config
func DeleteEgressConfig(ctx context.Context, origin core.Origin) *core.Status {
	return core.StatusOK()
}


*/
