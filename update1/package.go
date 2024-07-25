package update1

import (
	"context"
	"github.com/advanced-go/stdlib/core"
)

const (
	PkgPath = "github/advanced-go/guidance/update1"
)

func IngressRedirect(ctx context.Context, origin core.Origin, status string) *core.Status {
	return core.StatusOK()
}

func EgressRedirect(ctx context.Context, origin core.Origin, location string) *core.Status {
	return core.StatusOK()
}
