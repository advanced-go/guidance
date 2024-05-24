package resiliency

import (
	"context"
	"github.com/advanced-go/stdlib/core"
	"github.com/advanced-go/stdlib/httpx"
	"github.com/advanced-go/stdlib/json"
	"net/http"
	"net/url"
)

// http://localhost:8081/github/advanced-go/guidance:resiliency?reg=us&az=dallas&sz=dfwocp1&host=www.google.com

// Get - resource GET
func Get(ctx context.Context, h http.Header, values url.Values) (entries []Entry, status *core.Status) {
	return get[core.Log](ctx, core.AddRequestId(h), values)
}

func get[E core.ErrorHandler](ctx context.Context, h http.Header, values url.Values) (entries []Entry, status *core.Status) {
	var e E
	url := "github/advanced-go/documents:resiliency" + "?" + values.Encode()

	req, _ := http.NewRequestWithContext(ctx, http.MethodGet, url, nil)
	// Set headers
	resp, status1 := httpx.DoExchange(req)
	if !status1.OK() && !status1.NotFound() && !status1.Timeout() {
		e.Handle(status1, core.RequestId(h))
		return nil, status1
	}
	entries, status = json.New[[]Entry](resp.Body, h)
	if !status.OK() && !status1.NotFound() && !status1.Timeout() {
		e.Handle(status, core.RequestId(h))
	}
	return
}
