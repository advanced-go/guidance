package action1

import (
	"context"
	"github.com/advanced-go/stdlib/core"
)

const (
	PkgPath           = "github/advanced-go/guidance/action1"
	inferenceResource = "inference"
)

// AddRateLimiting - insert rate limiting action
func AddRateLimiting(ctx context.Context, origin core.Origin, inferenceId int, limit float64, burst int) *core.Status {
	return core.StatusOK()
}

// AddRouting - insert routing action
func AddRouting(ctx context.Context, origin core.Origin, inferenceId int, location string, percent int) *core.Status {
	return core.StatusOK()
}

// AddRedirect - insert redirect action
func AddRedirect(ctx context.Context, origin core.Origin, inferenceId int, location string, status string) *core.Status {
	return core.StatusOK()
}
