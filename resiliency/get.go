package resiliency

import (
	"context"
	"github.com/advanced-go/stdlib/core"
	"net/http"
	"net/url"
)

// http://localhost:8081/github/advanced-go/guidance:resiliency?reg=us&az=dallas&sz=dfwocp1&host=www.google.com

// Get - resource GET
func Get[T EntryConstraints](ctx context.Context, h http.Header, values url.Values) (entries []T, status *core.Status) {
	return get[core.Log, T](ctx, core.AddRequestId(h), values)
}

func get[E core.ErrorHandler, T EntryConstraints](ctx context.Context, h http.Header, values url.Values) (entries []T, status *core.Status) {
	entries, status = getEntries[T](ctx, h, values)
	if status.OK() || status.NotFound() || status.Timeout() {
		return
	}
	var e E
	e.Handle(status, core.RequestId(h))
	return
}
