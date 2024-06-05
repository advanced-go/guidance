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
	"github.com/advanced-go/stdlib/uri"
	"net/http"
	"net/url"
)

func resiliencyExchange[E core.ErrorHandler](r *http.Request, p *uri.Parsed) (*http.Response, *core.Status) {
	h2 := make(http.Header)
	h2.Add(httpx.ContentType, httpx.ContentTypeText)

	if p == nil {
		p1, status := httpx.ValidateURL(r.URL, module.Authority)
		if !status.OK() {
			return httpx.NewResponse[E](status.HttpCode(), h2, status.Err)
		}
		p = p1
	}
	switch r.Method {
	case http.MethodGet:
		return get[E](r.Context(), r.Header, r.URL, p.Version)
	case http.MethodDelete:
		return delete[E](r.Context(), r.Header, r.URL, p.Version)
	case http.MethodPut:
		return put[E](r, p.Version)
	case http.MethodPatch:
		return patch[E](r, p.Version)
	case http.MethodPost:
		return post[E](r, p.Version)
	default:
		status := core.NewStatusError(http.StatusBadRequest, errors.New(fmt.Sprintf("error invalid method: [%v]", r.Method)))
		return httpx.NewResponse[core.Log](status.HttpCode(), h2, status.Err)
	}
}

func get[E core.ErrorHandler](ctx context.Context, h http.Header, url *url.URL, version string) (resp *http.Response, status *core.Status) {
	var entries any
	var h2 http.Header

	switch version {
	case module.Ver1, "":
		entries, h2, status = resiliency1.Get(ctx, h, url.Query())
	case module.Ver2:
		entries, status = resiliency2.Get(ctx, h, url)
	default:
		status = core.NewStatusError(http.StatusBadRequest, errors.New(fmt.Sprintf("invalid version: [%v]", h.Get(core.XVersion))))
		h2 = make(http.Header)
		h2.Add(httpx.ContentType, httpx.ContentTypeText)
	}
	if !status.OK() {
		return httpx.NewResponse[E](status.HttpCode(), h2, status.Err)
	}
	return httpx.NewResponse[E](status.HttpCode(), h2, entries)
}

func delete[E core.ErrorHandler](ctx context.Context, h http.Header, url *url.URL, version string) (resp *http.Response, status *core.Status) {
	var h2 http.Header

	switch version {
	case module.Ver1, "":
		h2, status = resiliency1.Delete(ctx, h, url.Query())
	case module.Ver2:
		status = resiliency2.Delete(ctx, h, url)
	default:
		status = core.NewStatusError(http.StatusBadRequest, errors.New(fmt.Sprintf("invalid version: [%v]", h.Get(core.XVersion))))
		h2 = make(http.Header)
		h2.Add(httpx.ContentType, httpx.ContentTypeText)
	}
	return httpx.NewResponse[E](status.HttpCode(), h2, status.Err)
}

func put[E core.ErrorHandler](r *http.Request, version string) (resp *http.Response, status *core.Status) {
	var h2 http.Header

	switch version {
	case module.Ver1, "":
		h2, status = resiliency1.Put(r, nil)
	case module.Ver2:
		status = resiliency2.Put(r, nil)
	default:
		status = core.NewStatusError(http.StatusBadRequest, errors.New(fmt.Sprintf("invalid version: [%v]", r.Header.Get(core.XVersion))))
		h2 = make(http.Header)
		h2.Add(httpx.ContentType, httpx.ContentTypeText)
	}
	return httpx.NewResponse[E](status.HttpCode(), h2, status.Err)
}

func patch[E core.ErrorHandler](r *http.Request, version string) (resp *http.Response, status *core.Status) {
	var h2 http.Header

	switch version {
	case module.Ver1, "":
		h2, status = resiliency1.Patch(r, nil)
	case module.Ver2:
		status = resiliency2.Patch(r, nil)
	default:
		status = core.NewStatusError(http.StatusBadRequest, errors.New(fmt.Sprintf("invalid version: [%v]", r.Header.Get(core.XVersion))))
		h2 = make(http.Header)
		h2.Add(httpx.ContentType, httpx.ContentTypeText)
	}
	return httpx.NewResponse[E](status.HttpCode(), h2, status.Err)
}

func post[E core.ErrorHandler](r *http.Request, version string) (resp *http.Response, status *core.Status) {
	var h2 http.Header

	switch version {
	case module.Ver1, "":
		h2, status = resiliency1.Post(r, nil)
	case module.Ver2:
		status = resiliency2.Post(r, nil)
	default:
		status = core.NewStatusError(http.StatusBadRequest, errors.New(fmt.Sprintf("invalid version: [%v]", r.Header.Get(core.XVersion))))
		h2 = make(http.Header)
		h2.Add(httpx.ContentType, httpx.ContentTypeText)
	}
	return httpx.NewResponse[E](status.HttpCode(), h2, status.Err)
}
