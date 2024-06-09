package resiliency1

import (
	"context"
	"github.com/advanced-go/guidance/module"
	"github.com/advanced-go/stdlib/core"
	"github.com/advanced-go/stdlib/httpx"
	"github.com/advanced-go/stdlib/json"
	"net/http"
	"net/url"
)

// http://localhost:8081/github/advanced-go/guidance:resiliency?region=us&zone=dallas&sub-zone=dfwocp1&host=www.google.com

func get[E core.ErrorHandler](ctx context.Context, h http.Header, values url.Values) (entries []Entry, h2 http.Header, status *core.Status) {
	var e E

	if values == nil {
		return nil, nil, core.StatusNotFound()
	}
	if ctx == nil {
		ctx = context.Background()
	}
	url := resolver.Url(hostKey, module.DocumentsAuthority, module.DocumentsResourceV1, values, h)
	req, err := http.NewRequestWithContext(ctx, http.MethodGet, url, nil)
	if err != nil {
		return nil, nil, core.NewStatusError(core.StatusInvalidArgument, err)
	}
	req.Header.Set(core.XFrom, module.Authority)
	httpx.Forward(req.Header, h)
	resp, status1 := httpx.DoExchange(req)
	if !status1.OK() {
		e.Handle(status1, core.RequestId(h))
		return nil, nil, status1
	}
	entries, status = json.New[[]Entry](resp.Body, h)
	if !status.OK() {
		e.Handle(status, core.RequestId(h))
	}
	return
}
