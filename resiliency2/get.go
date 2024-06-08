package resiliency2

import (
	"context"
	"errors"
	"github.com/advanced-go/guidance/module"
	"github.com/advanced-go/stdlib/core"
	"github.com/advanced-go/stdlib/httpx"
	"github.com/advanced-go/stdlib/json"
	"github.com/advanced-go/stdlib/uri"
	"net/http"
	"net/url"
)

// http://localhost:8081/github/advanced-go/guidance:resiliency?reg=us&az=dallas&sz=dfwocp1&host=www.google.com

func get[E core.ErrorHandler](ctx context.Context, h http.Header, url *url.URL) (entries []Entry, status *core.Status) {
	var e E
	if url == nil {
		return nil, core.NewStatusError(core.StatusInvalidArgument, errors.New("invalid argument: URL is nil"))
	}
	url2 := uri.Expansion("", module.DocumentsAuthorityV2, module.DocumentsResourceV2, url.Query())
	req, err := http.NewRequestWithContext(ctx, http.MethodGet, url2, nil)
	if err != nil {
		return nil, core.NewStatusError(core.StatusInvalidArgument, err)
	}
	req.Header.Set(core.XFrom, module.Authority)
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
