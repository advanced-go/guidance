package resiliency

import (
	"context"
	"github.com/advanced-go/stdlib/core"
	"github.com/advanced-go/stdlib/httpx"
	"github.com/advanced-go/stdlib/io"
	"github.com/advanced-go/stdlib/json"
	"net/http"
	"net/url"
)

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

type entryConstraints interface {
	entryV1 | entryV2
}

var (
	listV1 []entryV1
	listV2 []entryV2
)

func getEntries[E entryConstraints](ctx context.Context, h http.Header, values url.Values) (entries []E, status *core.Status) {
	var buf []byte

	if h != nil {
		location := h.Get(httpx.ContentLocation)
		if location != "" {
			buf, status = io.ReadFile(location)
			if !status.OK() {
				return nil, status
			}
			if len(buf) == 0 {
				return nil, core.StatusNotFound()
			}
		}
	}
	switch p := any(&entries).(type) {
	case *[]entryV1:
		list := listV1
		if len(buf) > 0 {
			list, status = json.New[[]entryV1](buf, nil)
			if !status.OK() {
				return nil, status
			}
			if len(list) == 0 {
				return nil, core.StatusNotFound()
			}
		}
		*p, status = filterEntries[entryV1](ctx, list, values)
	case *[]entryV2:
		list := listV2
		if len(buf) > 0 {
			list, status = json.New[[]entryV2](buf, nil)
			if !status.OK() {
				return nil, status
			}
			if len(list) == 0 {
				return nil, core.StatusNotFound()
			}
		}
		*p, status = filterEntries[entryV2](ctx, list, values)
	default:
		return nil, core.NewStatus(http.StatusBadRequest)
	}
	return
}

func filterEntries[E entryConstraints](ctx context.Context, list any, values url.Values) (entries []E, status *core.Status) {
	switch ptr := any(&entries).(type) {
	case *[]entryV1:
		if l1, ok := list.([]entryV1); ok {
			if len(l1) == 0 {
				return nil, core.NewStatus(http.StatusNotFound)
			}
			filter := core.NewOrigin(values)
			for _, target := range l1 {
				if core.OriginMatch(target.Origin, filter) {
					*ptr = append(*ptr, target)
				}
			}
		} else {
			return nil, core.NewStatusError(core.StatusInvalidContent, core.NewInvalidBodyTypeError(list))
		}
		if len(*ptr) == 0 {
			return nil, core.NewStatus(http.StatusNotFound)
		}
	case *[]entryV2:
		if l2, ok := list.([]entryV2); ok {
			if len(l2) == 0 {
				return nil, core.NewStatus(http.StatusNotFound)
			}
			filter := core.NewOrigin(values)
			for _, target := range l2 {
				if core.OriginMatch(target.Origin, filter) {
					*ptr = append(*ptr, target)
				}
			}
		} else {
			return nil, core.NewStatusError(core.StatusInvalidContent, core.NewInvalidBodyTypeError(list))
		}
		if len(*ptr) == 0 {
			return nil, core.NewStatus(http.StatusNotFound)
		}
	default:
		return nil, core.NewStatus(http.StatusBadRequest)
	}
	return entries, core.StatusOK()
}

func addEntries[E entryConstraints](ctx context.Context, entries []E) *core.Status {
	if len(entries) == 0 {
		return core.StatusOK()
	}
	switch ptr := any(entries).(type) {
	case []entryV1:
		listV1 = append(listV1, ptr...)
	case []entryV2:
		listV2 = append(listV2, ptr...)
	default:
		return core.NewStatusError(core.StatusInvalidContent, core.NewInvalidBodyTypeError(ptr))
	}
	return core.StatusOK()
}

// createEntries - body supports []byte, io.ReadCloser, io.Reader
func createEntries[E entryConstraints](h http.Header, body any) (entries []E, status *core.Status) {
	if body == nil {
		return nil, core.NewStatus(core.StatusInvalidContent)
	}
	switch ptr := any(&entries).(type) {
	case *[]entryV1:
		*ptr, status = json.New[[]entryV1](body, h)
		if !status.OK() {
			return nil, status.AddLocation()
		}
		return entries, status
	case *[]entryV2:
		*ptr, status = json.New[[]entryV2](body, h)
		if !status.OK() {
			return nil, status.AddLocation()
		}
		return entries, status
	default:
		return nil, core.NewStatusError(core.StatusInvalidContent, core.NewInvalidBodyTypeError(body))
	}
}

/*
func testFilter[E entryConstraints](ctx context.Context, list []E, values url.Values) (entries []E, status *core.Status) {
	return nil, nil
}

func testCall[E entryConstraints]() {
	var list []entryV1

	testFilter[entryV1](nil, list, nil)
}


*/
