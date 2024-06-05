package resiliency1

import (
	"context"
	"github.com/advanced-go/guidance/module"
	"github.com/advanced-go/stdlib/core"
	"github.com/advanced-go/stdlib/httpx"
	"github.com/advanced-go/stdlib/json"
	"github.com/advanced-go/stdlib/uri"
	"net/http"
	"net/url"
)

// http://localhost:8081/github/advanced-go/guidance:resiliency?region=us&zone=dallas&sub-zone=dfwocp1&host=www.google.com

func get[E core.ErrorHandler](ctx context.Context, h http.Header, values url.Values) (entries []Entry, h2 http.Header, status *core.Status) {
	var e E

	h2 = make(http.Header)
	if values == nil {
		return nil, h2, core.StatusNotFound()
	}
	url2 := uri.Expansion("", module.DocumentsPath, module.DocumentsV1, values)
	req, _ := http.NewRequestWithContext(ctx, http.MethodGet, url2, nil)
	httpx.Forward(req.Header, h, core.XAuthority)
	resp, status1 := httpx.DoExchange(req)
	if !status1.OK() {
		e.Handle(status1, core.RequestId(h))
		h2.Add(httpx.ContentType, httpx.ContentTypeText)
		return nil, h2, status1
	}
	entries, status = json.New[[]Entry](resp.Body, h)
	if !status.OK() {
		h2.Add(httpx.ContentType, httpx.ContentTypeText)
		e.Handle(status, core.RequestId(h))
	} else {
		h2.Add(httpx.ContentType, httpx.ContentTypeJson)
	}
	return
}
