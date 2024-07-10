package schedule1

import (
	"context"
	"github.com/advanced-go/stdlib/core"
	"net/url"
)

const (
	PkgPath = "github/advanced-go/guidance/schedule1"
)

// Get - resource GET
func Get(ctx context.Context, values url.Values) (Entry, *core.Status) {
	return entryData[0], core.StatusOK()
}
