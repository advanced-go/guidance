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

func Get(r *http.Request) (*http.Response, *core.Status) {
	if r == nil {
		return httpx.NewResponseWithStatus(core.NewStatus(http.StatusBadRequest), nil)
	}
	switch r.Header.Get(core.XVersion) {
	case version1, "":
		entries, status := get[core.Log, entryV1](r.Context(), r.Header, r.URL.Query())
		if !status.OK() {
			return httpx.NewErrorResponseWithStatus(status)
		}
		if entries != nil {
			//buf,err := json
		}
		return nil, core.StatusOK()
	default:
		return httpx.NewErrorResponseWithStatus(core.NewStatusError(http.StatusBadRequest, errors.New(fmt.Sprintf("invalid version: [%v]", r.Header.Get(core.XVersion)))))
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
	return entries, core.StatusOK()
}

/*
func getV1[E core.ErrorHandler](r *http.Request) (*http.Response, *core.Status) {
	if r == nil {
		return httpx.NewErrorResponseWithStatus(core.NewStatus(http.StatusBadRequest))
	}
	entries, status := getEntriesV1(r.Context(), r.URL.Query())
	switch status.Code {
	case http.StatusNotFound, http.StatusGatewayTimeout, core.StatusDeadlineExceeded:
		return httpx.NewResponseWithStatus(status, "")
	case http.StatusOK:
		var resp *http.Response
		if entries != nil {

		}
		return resp, status
	default:
		var e E
		e.Handle(status, core.RequestId(r))
		return httpx.NewErrorResponseWithStatus(status)
	}
}


*/
