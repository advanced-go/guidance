package state1

import (
	"context"
	"github.com/advanced-go/stdlib/core"
)

const (
	PkgPath = "github/advanced-go/guidance/state1"
)

// Ingress

// GetIngressRedirectState - retrieve the state needed to start the ingress redirect agent
func GetIngressRedirectState(ctx context.Context, origin core.Origin) (IngressRedirectState, *core.Status) {
	var state IngressRedirectState
	NewIngressRedirectState(&state)
	return state, core.StatusOK()
}

// GetIngressRateLimitingState - retrieve the state needed to start the ingress resiliency agent
func GetIngressRateLimitingState(ctx context.Context, origin core.Origin) (IngressResiliencyState, *core.Status) {
	var state IngressResiliencyState
	NewIngressResiliencyState(&state)
	return state, core.StatusOK()
}

// Egress

// GetEgressRedirectState - retrieve the state needed to start the ingress redirect controller
func GetEgressRedirectState(ctx context.Context, origin core.Origin) (IngressRedirectState, *core.Status) {
	var state IngressRedirectState
	NewIngressRedirectState(&state)
	return state, core.StatusOK()
}

// GetEgressRateLimitingState - retrieve the state needed to start the egress resiliency agent
func GetEgressRateLimitingState(ctx context.Context, origin core.Origin) ([]EgressState, *core.Status) {
	return []EgressState{}, core.StatusOK()
}
