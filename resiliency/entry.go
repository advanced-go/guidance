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

var (
	listV1 []EntryV1
	listV2 []EntryV2
)

func getEntries[E EntryConstraints](ctx context.Context, h http.Header, values url.Values) (entries []E, status *core.Status) {
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
	case *[]EntryV1:
		list := listV1
		if len(buf) > 0 {
			list, status = json.New[[]EntryV1](buf, nil)
			if !status.OK() {
				return nil, status
			}
			if len(list) == 0 {
				return nil, core.StatusNotFound()
			}
		}
		*p, status = filterEntries[EntryV1](ctx, list, values)
	case *[]EntryV2:
		list := listV2
		if len(buf) > 0 {
			list, status = json.New[[]EntryV2](buf, nil)
			if !status.OK() {
				return nil, status
			}
			if len(list) == 0 {
				return nil, core.StatusNotFound()
			}
		}
		*p, status = filterEntries[EntryV2](ctx, list, values)
	default:
		return nil, core.NewStatus(http.StatusBadRequest)
	}
	return
}

func filterEntries[E EntryConstraints](ctx context.Context, list any, values url.Values) (entries []E, status *core.Status) {
	switch ptr := any(&entries).(type) {
	case *[]EntryV1:
		if l1, ok := list.([]EntryV1); ok {
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
	case *[]EntryV2:
		if l2, ok := list.([]EntryV2); ok {
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

func addEntries[E EntryConstraints](ctx context.Context, _ http.Header, entries []E) *core.Status {
	if len(entries) == 0 {
		return core.StatusOK()
	}
	switch ptr := any(entries).(type) {
	case []EntryV1:
		listV1 = append(listV1, ptr...)
	case []EntryV2:
		listV2 = append(listV2, ptr...)
	default:
		return core.NewStatusError(core.StatusInvalidContent, core.NewInvalidBodyTypeError(ptr))
	}
	return core.StatusOK()
}

// createEntries - body supports []byte, io.ReadCloser, io.Reader
func createEntries[E EntryConstraints](h http.Header, body any) (entries []E, status *core.Status) {
	if body == nil {
		return nil, core.NewStatus(core.StatusInvalidContent)
	}
	switch ptr := any(&entries).(type) {
	case *[]EntryV1:
		*ptr, status = json.New[[]EntryV1](body, h)
		if !status.OK() {
			return nil, status.AddLocation()
		}
		return entries, status
	case *[]EntryV2:
		*ptr, status = json.New[[]EntryV2](body, h)
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
	var list []EntryV1

	testFilter[EntryV1](nil, list, nil)
}


*/
