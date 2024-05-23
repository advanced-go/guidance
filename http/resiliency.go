package http

import (
	"context"
	"errors"
	"fmt"
	"github.com/advanced-go/guidance/module"
	"github.com/advanced-go/guidance/resiliency"
	"github.com/advanced-go/stdlib/core"
	"github.com/advanced-go/stdlib/httpx"
	"net/http"
	"net/url"
)

func resiliencyExchange(r *http.Request) (*http.Response, *core.Status) {
	switch r.Method {
	case http.MethodGet:
		return get(r.Context(), r.Header, r.URL.Query())
	case http.MethodPut:
		return put(r)
	default:
		status := core.NewStatusError(http.StatusBadRequest, errors.New(fmt.Sprintf("error invalid method: [%v]", r.Method)))
		return httpx.NewResponseWithStatus(status, status.Err)
	}
}

func get(ctx context.Context, h http.Header, values url.Values) (resp *http.Response, status *core.Status) {
	var entries any

	switch h.Get(core.XVersion) {
	case module.Ver1, "":
		entries, status = resiliency.Get[resiliency.EntryV1](ctx, h, values)
	case module.Ver2:
		entries, status = resiliency.Get[resiliency.EntryV2](ctx, h, values)
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

func put(r *http.Request) (resp *http.Response, status *core.Status) {
	switch r.Header.Get(core.XVersion) {
	case module.Ver1, module.Ver2, "":
		status = resiliency.Put[*http.Request](r.Context(), r.Header, r)
	default:
		status1 := core.NewStatusError(http.StatusBadRequest, errors.New(fmt.Sprintf("invalid version: [%v]", r.Header.Get(core.XVersion))))
		return httpx.NewResponseWithStatus(status1, status1.Err)
	}
	return httpx.NewResponseWithStatus(status, status.Err)
}
