package resiliency2

import (
	"context"
	"errors"
	"fmt"
	"github.com/advanced-go/guidance/module"
	"github.com/advanced-go/stdlib/core"
	"github.com/advanced-go/stdlib/httpx"
	json2 "github.com/advanced-go/stdlib/json"
	"net/http"
	"net/url"
	"strings"
)

const (
	PkgPath = "github/advanced-go/guidance/resiliency2"
)

type Entry struct {
	Origin    core.Origin
	Status    string `json:"status"`
	CreatedTS string `json:"created-ts"`
	UpdatedTS string `json:"updated-ts"`

	// Timeout
	Timeout string `json:"timeout"`

	// Rate Limiting
	RateLimit string `json:"rate-limit"`
	RateBurst string `json:"rate-burst"`
}

type PostData struct{}

// Get - resource GET
func Get(ctx context.Context, h http.Header, url *url.URL) ([]Entry, *core.Status) {
	if url == nil {
		return nil, core.StatusBadRequest()
	}
	switch url.Path {
	case module.DocumentsResourceV2:
		return get[core.Log](ctx, core.AddRequestId(h), url)
	default:
		return nil, core.StatusBadRequest()
	}
}

// Delete - resource DELETE
func Delete(ctx context.Context, h http.Header, url *url.URL) *core.Status {
	if url == nil || !strings.HasPrefix(url.Path, module.ResiliencyResource) {
		return core.NewStatusError(http.StatusBadRequest, errors.New(fmt.Sprintf("invalid URL")))
	}
	switch url.Path {
	case module.ResiliencyResource:
		return delete[core.Log](ctx, core.AddRequestId(h), url.Query())
	default:
		return core.StatusBadRequest()
	}
}

// Put - resource PUT
func Put(r *http.Request, body []Entry) *core.Status {
	if r == nil || r.URL == nil {
		return core.NewStatus(http.StatusBadRequest)
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
	switch r.URL.Path {
	case module.DocumentsResourceV2:
		return put[core.Log](r.Context(), core.AddRequestId(r.Header), body)
	default:
		return core.StatusBadRequest()
	}
}

// Post - resource POST
func Post(r *http.Request, body *PostData) *core.Status {
	if r == nil || r.URL == nil {
		return core.NewStatus(http.StatusBadRequest)
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
	switch r.URL.Path {
	case module.DocumentsResourceV2:
		return post[core.Log](r.Context(), core.AddRequestId(r.Header), body)
	default:
		return core.StatusBadRequest()
	}
}

// Patch - resource PATCH
func Patch(r *http.Request, body *httpx.Patch) *core.Status {
	if r == nil || r.URL == nil {
		return core.NewStatus(http.StatusBadRequest)
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
	switch r.URL.Path {
	case module.DocumentsResourceV2:
		return patch[core.Log](r.Context(), core.AddRequestId(r.Header), body)
	default:
		return core.StatusBadRequest()
	}
}
