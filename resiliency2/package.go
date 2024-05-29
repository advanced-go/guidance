package resiliency

import (
	"context"
	"github.com/advanced-go/stdlib/core"
	"github.com/advanced-go/stdlib/httpx"
	"net/http"
	"net/url"
)

const (
	PkgPath            = "github/advanced-go/guidance/resiliency2"
	documentsAuthority = "github/advanced-go/documentsv2"
	documentsResource  = "resiliency2"
	resiliencyRoot     = "resiliency2"
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
	if url == nil || url.Path != resiliencyRoot {
		return nil, core.StatusBadRequest()
	}
	switch url.Path {
	case resiliencyRoot:
		return get[core.Log](ctx, core.AddRequestId(h), url.Query())
	default:
		return nil, core.StatusBadRequest()
	}
}

// Delete - resource DELETE
func Delete(ctx context.Context, h http.Header, url *url.URL) *core.Status {
	if url == nil {
		return core.StatusBadRequest()
	}
	switch url.Path {
	case resiliencyRoot:
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
	switch r.URL.Path {
	case resiliencyRoot:
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
	switch r.URL.Path {
	case resiliencyRoot:
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
	switch r.URL.Path {
	case resiliencyRoot:
		return patch[core.Log](r.Context(), core.AddRequestId(r.Header), body)
	default:
		return core.StatusBadRequest()
	}
}
