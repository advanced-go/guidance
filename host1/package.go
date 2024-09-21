package host1

import (
	"context"
	"github.com/advanced-go/stdlib/core"
)

const (
	PkgPath = "github/advanced-go/guidance/host1"
)

const (
	EntryResourceName   = "entry"
	IngressResourceName = "ingress"
	EgressResourceName  = "egress"
)

// CDC
// Functionality:
//  Startup - read all hosts and determine if there is an active redirect for ingress and egress
//  Real time - query to determine status changes to hosts and redirects, then notify the
//              appropriate agent

// HostQuery - on startup, retrieve existing hosts and redirect status
func HostQuery(ctx context.Context, origin core.Origin, state *CDCState) ([]EntryStatus, []CDCState, *core.Status) {
	var cdc []CDCState
	return []EntryStatus{}, cdc, core.StatusOK()
}

// RedirectStateChanges - retrieve redirects where the active/inactive state has changed, so that the
//
//	appropriate agent can be notified.
func RedirectStateChanges(ctx context.Context, origin core.Origin, state *CDCState, ingress bool) ([]core.Origin, *core.Status) {
	return []core.Origin{}, core.StatusOK()
}
