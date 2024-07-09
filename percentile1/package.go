package percentile1

import (
	"context"
	"github.com/advanced-go/stdlib/core"
	"net/url"
)

const (
	PkgPath   = "github/advanced-go/guidance/percentile1"
	routeName = "documents-resiliency"
	hostKey   = "docs-host"
)

// Get - resource GET
func Get(ctx context.Context, values url.Values) (Entry, *core.Status) {
	return Entry{}, core.StatusOK()
}
