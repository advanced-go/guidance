package http

import (
	"context"
	"errors"
	"fmt"
	"github.com/advanced-go/guidance/module"
	resiliency1 "github.com/advanced-go/guidance/resiliency1"
	resiliency2 "github.com/advanced-go/guidance/resiliency2"
	"github.com/advanced-go/stdlib/core"
	"github.com/advanced-go/stdlib/httpx"
	"net/http"
	"net/url"
)

func resiliencyExchange(r *http.Request, path string, values url.Values) (*http.Response, *core.Status) {
	//Need to reset the URL to only the path after the ":"

	switch r.Method {
	case http.MethodGet:
		return get(r.Context(), r.Header, r.URL)
	case http.MethodDelete:
		return delete(r.Context(), r.Header, r.URL)
	case http.MethodPut:
		return put(r)
	case http.MethodPatch:
		return patch(r)
	case http.MethodPost:
		return patch(r)
	default:
		status := core.NewStatusError(http.StatusBadRequest, errors.New(fmt.Sprintf("error invalid method: [%v]", r.Method)))
		return httpx.NewResponseWithStatus(status, status.Err)
	}
}

func get(ctx context.Context, h http.Header, url *url.URL) (resp *http.Response, status *core.Status) {
	var entries any

	switch h.Get(core.XVersion) {
	case module.Ver1, "":
		entries, status = resiliency1.Get(ctx, h, url)
	case module.Ver2:
		entries, status = resiliency2.Get(ctx, h, url)
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

func delete(ctx context.Context, h http.Header, url *url.URL) (resp *http.Response, status *core.Status) {
	switch h.Get(core.XVersion) {
	case module.Ver1, "":
		status = resiliency1.Delete(ctx, h, url)
	case module.Ver2:
		status = resiliency2.Delete(ctx, h, url)
	default:
		status1 := core.NewStatusError(http.StatusBadRequest, errors.New(fmt.Sprintf("invalid version: [%v]", h.Get(core.XVersion))))
		return httpx.NewResponseWithStatus(status1, status1.Err)
	}
	return httpx.NewResponseWithStatus(status, status.Err)
}

func put(r *http.Request) (resp *http.Response, status *core.Status) {
	switch r.Header.Get(core.XVersion) {
	case module.Ver1, "":
		status = resiliency1.Put(r, nil)
	case module.Ver2:
		status = resiliency2.Put(r, nil)
	default:
		status1 := core.NewStatusError(http.StatusBadRequest, errors.New(fmt.Sprintf("invalid version: [%v]", r.Header.Get(core.XVersion))))
		return httpx.NewResponseWithStatus(status1, status1.Err)
	}
	return httpx.NewResponseWithStatus(status, status.Err)
}

func patch(r *http.Request) (resp *http.Response, status *core.Status) {
	switch r.Header.Get(core.XVersion) {
	case module.Ver1, "":
		status = resiliency1.Patch(r, nil)
	case module.Ver2:
		status = resiliency2.Patch(r, nil)
	default:
		status1 := core.NewStatusError(http.StatusBadRequest, errors.New(fmt.Sprintf("invalid version: [%v]", r.Header.Get(core.XVersion))))
		return httpx.NewResponseWithStatus(status1, status1.Err)
	}
	return httpx.NewResponseWithStatus(status, status.Err)
}

func post(r *http.Request) (resp *http.Response, status *core.Status) {
	switch r.Header.Get(core.XVersion) {
	case module.Ver1, "":
		status = resiliency1.Post(r, nil)
	case module.Ver2:
		status = resiliency2.Post(r, nil)
	default:
		status1 := core.NewStatusError(http.StatusBadRequest, errors.New(fmt.Sprintf("invalid version: [%v]", r.Header.Get(core.XVersion))))
		return httpx.NewResponseWithStatus(status1, status1.Err)
	}
	return httpx.NewResponseWithStatus(status, status.Err)
}
