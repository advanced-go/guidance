package resiliency1

import (
	"context"
	"errors"
	"fmt"
	"github.com/advanced-go/stdlib/core"
	"github.com/advanced-go/stdlib/httpx"
	json2 "github.com/advanced-go/stdlib/json"
	"net/http"
	"net/url"
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
func Delete(ctx context.Context, h http.Header, values url.Values) (http.Header, *core.Status) {
	return delete[core.Log](ctx, core.AddRequestId(h), values)
}

// Put - resource PUT, with optional content override
func Put(r *http.Request, body []Entry) (http.Header, *core.Status) {
	h2 := make(http.Header)
	h2.Add(httpx.ContentType, httpx.ContentTypeText)
	if r == nil {
		return h2, core.NewStatusError(core.StatusInvalidArgument, errors.New("error: request is nil"))
	}
	if body == nil {
		content, status := json2.New[[]Entry](r.Body, r.Header)
		if !status.OK() {
			var e core.Log
			e.Handle(status, core.RequestId(r.Header))
			h2.Add(httpx.ContentType, httpx.ContentTypeText)
			return h2, status
		}
		body = content
	}
	return put[core.Log](r.Context(), core.AddRequestId(r.Header), body)
}

// Post - resource POST, with optional content override
func Post(r *http.Request, body *PostData) (http.Header, *core.Status) {
	h2 := make(http.Header)
	if r == nil {
		h2.Add(httpx.ContentType, httpx.ContentTypeText)
		return h2, core.NewStatusError(core.StatusInvalidArgument, errors.New("error: request is nil"))
	}
	if body == nil {
		content, status := json2.New[PostData](r.Body, r.Header)
		if !status.OK() {
			var e core.Log
			e.Handle(status, core.RequestId(r.Header))
			h2.Add(httpx.ContentType, httpx.ContentTypeText)
			return h2, status
		}
		body = &content
	}
	return post[core.Log](r.Context(), core.AddRequestId(r.Header), body)
}

// Patch - resource PATCH, with optional content override
func Patch(r *http.Request, body *httpx.Patch) (http.Header, *core.Status) {
	h2 := make(http.Header)
	if r == nil {
		h2.Add(httpx.ContentType, httpx.ContentTypeText)
		return h2, core.NewStatusError(core.StatusInvalidArgument, errors.New("error: request is nil"))
	}
	if body == nil {
		content, status := json2.New[httpx.Patch](r.Body, r.Header)
		if !status.OK() {
			var e core.Log
			e.Handle(status, core.RequestId(r.Header))
			h2.Add(httpx.ContentType, httpx.ContentTypeText)
			return h2, status
		}
		body = &content
	}
	return patch[core.Log](r.Context(), core.AddRequestId(r.Header), body)
}
