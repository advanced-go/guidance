package action1

import (
	"context"
	"github.com/advanced-go/stdlib/core"
)

const (
	PkgPath           = "github/advanced-go/guidance/action1"
	inferenceResource = "inference"
)

// InsertEgress - insert egress actions
func InsertEgress(ctx context.Context, origin core.Origin, entries []Entry) *core.Status {
	return core.StatusOK()
}

// InsertIngress - insert ingress actions
func InsertIngress(ctx context.Context, origin core.Origin, entries []Entry) *core.Status {
	return core.StatusOK()
}
