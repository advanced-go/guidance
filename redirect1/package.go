package redirect1

import (
	"context"
	"github.com/advanced-go/stdlib/core"
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

// Ingress

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

/*
// UpdateRedirectConfig - update the ingress redirect configuration
func UpdateRedirectConfig(ctx context.Context, origin core.Origin, status string) *core.Status {
	return core.StatusOK()
}

// DeleteEgressConfig - delete the route egress config
func DeleteEgressConfig(ctx context.Context, origin core.Origin) *core.Status {
	return core.StatusOK()
}


*/
