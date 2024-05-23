package resiliency

import (
	"context"
	"github.com/advanced-go/stdlib/core"
	"github.com/advanced-go/stdlib/httpx"
	"net/http"
	"net/url"
)

// Post - resource POST
func Post(r *http.Request) (*http.Response, *core.Status) {
	if r == nil {
		return httpx.NewResponseWithStatus(core.NewStatus(http.StatusBadRequest), nil)
	}
	return post[core.Log, *http.Request](r.Context(), r.Header, r.URL.Query(), r)
}

type postBodyConstraints interface {
	*http.Request
}

func post[E core.ErrorHandler, T PutBodyConstraints](ctx context.Context, h http.Header, values url.Values, body T) (*http.Response, *core.Status) {

	return nil, nil
}
