package resiliency

import (
	"context"
	"github.com/advanced-go/guidance/module"
	"github.com/advanced-go/stdlib/core"
	"github.com/advanced-go/stdlib/httpx"
	"net/http"
	"net/url"
)

const (
	PkgPath = "github/advanced-go/guidance/resiliency1"
)

type EntryV1 struct {
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

type EntryV2 struct {
	Origin    core.Origin
	Version   string `json:"version"`
	Status    string `json:"status"`
	CreatedTS string `json:"created-ts"`
	UpdatedTS string `json:"updated-ts"`

	// Timeout
	Timeout string `json:"timeout"`

	// Rate Limiting
	RateLimit string `json:"rate-limit"`
	RateBurst string `json:"rate-burst"`
}

type EntryConstraints interface {
	EntryV1 | EntryV2
}

type PutBodyConstraints interface {
	[]EntryV1 | []EntryV2 | []byte | *http.Request
}

// Get - resource GET
func Get[T EntryConstraints](ctx context.Context, h http.Header, u *url.URL) (entries []T, status *core.Status) {
	var e core.Log
	// Validate the path
	//ver, path, status := httpx.ValidateURL(u)
	url := module.BuildDocumentsPath(module.Ver1, u.Query())

	// Make upstream calls, with correct upstream version based on requested resiliency version
	req, _ := http.NewRequestWithContext(ctx, http.MethodGet, url, nil)
	httpx.Forward(req.Header, h)
	_, status1 := httpx.DoExchange(req)
	if status1.NotFound() || status1.Timeout() {
		return nil, core.StatusOK()
	}
	if !status1.OK() {
		e.Handle(status1, core.RequestId(h))
		return nil, status1
	}

	return get[core.Log, T](ctx, core.AddRequestId(h), u.Query())
}

// Delete - resource DELETE
func Delete(ctx context.Context, h http.Header, u *url.URL) *core.Status {
	if h == nil {
		return core.NewStatus(http.StatusBadRequest)
	}
	if u.Query() == nil {
		return core.NewStatus(http.StatusBadRequest)
	}
	return core.StatusNotFound()
}

// Put - resource PUT
func Put[T PutBodyConstraints](r *http.Request) *core.Status {
	//if body == nil {
	//	return core.NewStatus(http.StatusBadRequest)
	//}
	return put[core.Log, *http.Request](r.Context(), core.AddRequestId(r.Header), r.Body)
}

// Post - resource POST
func Post(r *http.Request) *core.Status {
	if r == nil {
		return core.NewStatus(http.StatusBadRequest)
	}
	return post[core.Log, *http.Request](r.Context(), r.Header, r.URL.Query(), r)
}
