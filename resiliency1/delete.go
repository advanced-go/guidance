package resiliency1

import (
	"context"
	"github.com/advanced-go/stdlib/core"
	"net/http"
	"net/url"
)

func delete[E core.ErrorHandler](ctx context.Context, h http.Header, values url.Values) (http.Header, *core.Status) {
	h2 := make(http.Header)
	if values == nil {
		return h2, core.StatusNotFound()
	}

	return h2, core.StatusOK()
}
