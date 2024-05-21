package resiliency

import (
	"context"
	"github.com/advanced-go/stdlib/core"
	"github.com/advanced-go/stdlib/io"
	"github.com/advanced-go/stdlib/json"
	"net/http"
	"net/url"
)

type entryConstraints interface {
	entryV1 | entryV2
}

var listV1 []entryV1
var listV2 []entryV2

type entryV1 struct {
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

type entryV2 struct {
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

func getEntries[T entryConstraints](ctx context.Context, values url.Values) (entries []T, status *core.Status) {
	var buf []byte

	location := values.Get(httpx.ContentLocation)
	if location != "" {
		buf, status = io.ReadFile(location)
		if !status.OK() {
			return nil, status
		}
	}
	switch p := any(&entries).(type) {
	case *[]entryV1:
		if len(buf) > 0 {
			*p, status = json.New[[]entryV1](buf, nil)
			return
		}
		if len(listV1) == 0 {
			return entries, core.NewStatus(http.StatusNotFound)
		}
		if values == nil {
			return entries, core.StatusOK()
		}
		return filterEntries[T](ctx, values)
	case *[]entryV2:
		if len(buf) > 0 {
			*p, status = json.New[[]entryV2](buf, nil)
			return
		}
		if len(listV2) == 0 {
			return entries, core.NewStatus(http.StatusNotFound)
		}
		if values == nil {
			return entries, core.StatusOK()
		}
		return filterEntries[T](ctx, values)
	default:
		return nil, core.NewStatus(http.StatusBadRequest)
	}
}

func filterEntries[T entryConstraints](ctx context.Context, values url.Values) (entries []T, status *core.Status) {
	switch ptr := any(&entries).(type) {
	case *[]entryV1:
		filter := core.NewOrigin(values)
		for _, target := range listV1 {
			if core.OriginMatch(target.Origin, filter) {
				*ptr = append(*ptr, target)
			}
		}
		if len(*ptr) == 0 {
			return nil, core.NewStatus(http.StatusNotFound)
		}
		return entries, core.StatusOK()
	case *[]entryV2:
		filter := core.NewOrigin(values)
		for _, target := range listV2 {
			if core.OriginMatch(target.Origin, filter) {
				*ptr = append(*ptr, target)
			}
		}
		if len(*ptr) == 0 {
			return nil, core.NewStatus(http.StatusNotFound)
		}
		return entries, core.StatusOK()
	default:
		return nil, core.NewStatus(http.StatusBadRequest)
	}
}

func addEntriesV1(ctx context.Context, e []entryV1) *core.Status {
	for _, item := range e {
		//item.CreatedTS = time.Now().UTC()
		listV1 = append(listV1, item)
		//status = logActivity(ctx, item)
	}
	return core.StatusOK()
}
