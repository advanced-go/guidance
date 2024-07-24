package authority1

import (
	"context"
	"github.com/advanced-go/stdlib/core"
	"net/url"
)

const (
	PkgPath = "github/advanced-go/observation/authority1"
)

func Query_NOT_USED(ctx context.Context, origin core.Origin, values url.Values) ([]Entry, *core.Status) {
	return []Entry{}, core.StatusOK()
}
