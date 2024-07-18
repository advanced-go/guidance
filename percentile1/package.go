package percentile1

import (
	"context"
	"github.com/advanced-go/stdlib/core"
	"time"
)

const (
	PkgPath                   = "github/advanced-go/guidance/percentile1"
	PercentilePollingDuration = time.Hour * 12
)

// Get - resource GET
func Get(ctx context.Context, origin core.Origin) (Entry, *core.Status) {
	return entryData[0], core.StatusOK()
}
