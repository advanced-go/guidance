package resiliency1

import (
	"context"
	"errors"
	"fmt"
	"github.com/advanced-go/guidance/module"
	"github.com/advanced-go/stdlib/core"
	"github.com/advanced-go/stdlib/httpx"
	json2 "github.com/advanced-go/stdlib/json"
	"github.com/advanced-go/stdlib/uri"
	"net/http"
	"net/url"
	"strings"
)

const (
	PkgPath = "github/advanced-go/guidance/resiliency1"
)

func errorInvalidURL(path string) *core.Status {
	return core.NewStatusError(core.StatusInvalidArgument, errors.New(fmt.Sprintf("invalid argument: URL path is invalid %v", path)))
}

// Get - resource GET
func Get(ctx context.Context, h http.Header, values url.Values) ([]Entry, http.Header, *core.Status) {
	return get[core.Log](ctx, core.AddRequestId(h), values)
	/*
		if url == nil || !strings.HasPrefix(url.Path, "/"+module.Authority) {
			return nil, core.NewStatusError(core.StatusInvalidArgument, errors.New(fmt.Sprintf("invalid or nil URL")))
		}
		if url.Query() == nil {
			return nil, core.NewStatusError(core.StatusInvalidContent, errors.New(fmt.Sprintf("query arguments are nil")))
		}
		p := uri.Uproot(url.Path)
		switch p.Resource {
		case module.ResiliencyResource:
			return get[core.Log](ctx, core.AddRequestId(h), url)
		default:
			return nil, errorInvalidURL(url.Path)
		}

	*/
}

// Delete - resource DELETE
func Delete(ctx context.Context, h http.Header, url *url.URL) *core.Status {
	if url == nil || !strings.HasPrefix(url.Path, "/"+module.Authority) {
		return core.NewStatusError(core.StatusInvalidArgument, errors.New(fmt.Sprintf("invalid URL")))
	}
	p := uri.Uproot(url.Path)
	switch p.Resource {
	case module.ResiliencyResource:
		return delete[core.Log](ctx, core.AddRequestId(h), url)
	default:
		return errorInvalidURL(url.Path)
	}
}

// Put - resource PUT, with optional content override
func Put(r *http.Request, body []Entry) *core.Status {
	if r == nil || r.URL == nil || !strings.HasPrefix(r.URL.Path, "/"+module.Authority) {
		return core.NewStatusError(core.StatusInvalidArgument, errors.New("invalid URL"))
	}
	if body == nil {
		content, status := json2.New[[]Entry](r.Body, r.Header)
		if !status.OK() {
			var e core.Log
			e.Handle(status, core.RequestId(r.Header))
			return status
		}
		body = content
	}
	p := uri.Uproot(r.URL.Path)
	switch p.Resource {
	case module.ResiliencyResource:
		return put[core.Log](r.Context(), core.AddRequestId(r.Header), body)
	default:
		return errorInvalidURL(r.URL.Path)
	}
}

// Post - resource POST, with optional content override
func Post(r *http.Request, body *PostData) *core.Status {
	if r == nil || r.URL == nil || !strings.HasPrefix(r.URL.Path, "/"+module.Authority) {
		return core.NewStatusError(core.StatusInvalidArgument, errors.New("invalid URL"))
	}
	if body == nil {
		content, status := json2.New[PostData](r.Body, r.Header)
		if !status.OK() {
			var e core.Log
			e.Handle(status, core.RequestId(r.Header))
			return status
		}
		body = &content
	}
	p := uri.Uproot(r.URL.Path)
	switch p.Resource {
	case module.ResiliencyResource:
		return post[core.Log](r.Context(), core.AddRequestId(r.Header), body)
	default:
		return errorInvalidURL(r.URL.Path)
	}
}

// Patch - resource PATCH, with optional content override
func Patch(r *http.Request, body *httpx.Patch) *core.Status {
	if r == nil || r.URL == nil || !strings.HasPrefix(r.URL.Path, "/"+module.Authority) {
		return core.NewStatusError(core.StatusInvalidArgument, errors.New("invalid URL"))
	}
	if body == nil {
		content, status := json2.New[httpx.Patch](r.Body, r.Header)
		if !status.OK() {
			var e core.Log
			e.Handle(status, core.RequestId(r.Header))
			return status
		}
		body = &content
	}
	p := uri.Uproot(r.URL.Path)
	switch p.Resource {
	case module.ResiliencyResource:
		return patch[core.Log](r.Context(), core.AddRequestId(r.Header), body)
	default:
		return errorInvalidURL(r.URL.Path)
	}
}
