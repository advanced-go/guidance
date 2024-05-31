package resiliency1

import (
	"context"
	"github.com/advanced-go/stdlib/core"
	"github.com/advanced-go/stdlib/httpx"
	"github.com/advanced-go/stdlib/json"
	"github.com/advanced-go/stdlib/uri"
	"net/http"
	"net/url"
)

// http://localhost:8081/github/advanced-go/guidance:resiliency?region=us&zone=dallas&sub-zone=dfwocp1&host=www.google.com

func get[E core.ErrorHandler](ctx context.Context, h http.Header, values url.Values) (entries []Entry, status *core.Status) {
	var e E
	url := uri.Expansion("", documentsPath, documentsV1, values)

	req, _ := http.NewRequestWithContext(ctx, http.MethodGet, url, nil)
	httpx.Forward(req.Header, h)
	resp, status1 := httpx.DoExchange(req)
	if !status1.OK() {
		e.Handle(status1, core.RequestId(h))
		return nil, status1
	}
	entries, status = json.New[[]Entry](resp.Body, h)
	if !status.OK() {
		e.Handle(status, core.RequestId(h))
	}
	return
}
