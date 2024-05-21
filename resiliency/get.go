package resiliency

import (
	"context"
	"errors"
	"fmt"
	"github.com/advanced-go/stdlib/core"
	"github.com/advanced-go/stdlib/httpx"
	"github.com/advanced-go/stdlib/json"
	"net/http"
	"net/url"
)

const (
	version1 = "v1"
	version2 = "v2"
)

// http://localhost:8081/github/advanced-go/guidance:resiliency?reg=us&az=dallas&sz=dfwocp1&host=www.google.com

// Get - resource GET
func Get(ctx context.Context, h http.Header, values url.Values) (*http.Response, *core.Status) {
	if h == nil {
		return httpx.NewResponseWithStatus(core.NewStatus(http.StatusBadRequest), nil)
	}
	if values == nil {
		return httpx.NewResponseWithStatus(core.NewStatus(http.StatusBadRequest), nil)
	}
	switch h.Get(core.XVersion) {
	case version1, "":
		entries, status := get[core.Log, entryV1](ctx, h, values)
		if status.NotFound() || status.Timeout() {
			return httpx.NewResponseWithStatus(status, nil)
		}
		if !status.OK() {
			return httpx.NewResponseWithStatus(status, status.Err)
		}
		if entries != nil {
			var e core.Log
			rc, status1 := json.NewReadCloser(entries)
			if !status1.OK() {
				e.Handle(status1, core.RequestId(h))
				return httpx.NewResponseWithStatus(status, status.Err)
			}
			rh := make(http.Header)
			rh.Add(httpx.ContentType, httpx.ContentTypeJson)
			return &http.Response{StatusCode: status1.HttpCode(), Status: status1.String(), Header: rh, Body: rc}, core.StatusOK()
		}
		return nil, core.StatusOK()
	default:
		status := core.NewStatusError(http.StatusBadRequest, errors.New(fmt.Sprintf("invalid version: [%v]", h.Get(core.XVersion))))
		return httpx.NewResponseWithStatus(status, status.Err)
	}
}

func get[E core.ErrorHandler, T entryConstraints](ctx context.Context, h http.Header, values url.Values) ([]T, *core.Status) {
	entries, status := getEntries[T](ctx, values)
	if status.NotFound() || status.Timeout() {
		return nil, status
	}
	if !status.OK() {
		var e E
		e.Handle(status, core.RequestId(h))
		return nil, status
	}
	return entries, status
}
