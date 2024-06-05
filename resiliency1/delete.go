package resiliency1

import (
	"context"
	"github.com/advanced-go/stdlib/core"
	"net/http"
	"net/url"
)

func delete[E core.ErrorHandler](ctx context.Context, h http.Header, values url.Values) (http.Header, *core.Status) {
	if values == nil {
		return nil, core.StatusNotFound()
	}
	return nil, core.StatusOK()
}
