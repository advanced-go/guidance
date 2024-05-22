package resiliency

import (
	"context"
	"errors"
	"fmt"
	"github.com/advanced-go/stdlib/core"
	"github.com/advanced-go/stdlib/httpx"
	"net/http"
	"net/url"
)

const (
	version1 = "v1"
	version2 = "v2"
)

// http://localhost:8081/github/advanced-go/guidance:resiliency?reg=us&az=dallas&sz=dfwocp1&host=www.google.com

// Get - resource GET
func Get(ctx context.Context, h http.Header, values url.Values) (resp *http.Response, status *core.Status) {
	var entries any

	if h == nil {
		return httpx.NewResponseWithStatus(core.NewStatus(http.StatusBadRequest), nil)
	}
	if values == nil {
		return httpx.NewResponseWithStatus(core.NewStatus(http.StatusBadRequest), nil)
	}

	switch h.Get(core.XVersion) {
	case version1, "":
		entries, status = get[core.Log, entryV1](ctx, h, values)
	case version2:
		entries, status = get[core.Log, entryV2](ctx, h, values)
	default:
		status = core.NewStatusError(http.StatusBadRequest, errors.New(fmt.Sprintf("invalid version: [%v]", h.Get(core.XVersion))))
		return httpx.NewResponseWithStatus(status, status.Err)
	}
	if !status.OK() {
		return httpx.NewResponseWithStatus(status, status.Err)
	}
	resp, status = httpx.NewJsonResponse(entries, nil)
	if !status.OK() {
		var e core.Log
		e.Handle(status, core.RequestId(h))
		return httpx.NewResponseWithStatus(status, status.Err)
	}
	return
}

func get[E core.ErrorHandler, T entryConstraints](ctx context.Context, h http.Header, values url.Values) (entries []T, status *core.Status) {
	entries, status = getEntries[T](ctx, values)
	if status.OK() || status.NotFound() || status.Timeout() {
		return
	}
	var e E
	e.Handle(status, core.RequestId(h))
	return
}
