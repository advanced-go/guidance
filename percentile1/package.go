package percentile1

import (
	"context"
	"github.com/advanced-go/stdlib/core"
)

const (
	PkgPath = "github/advanced-go/guidance/percentile1"
)

// Get - resource GET
func Get(ctx context.Context, origin core.Origin) (Entry, *core.Status) {
	return entryData[0], core.StatusOK()
}
