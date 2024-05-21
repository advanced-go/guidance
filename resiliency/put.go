package resiliency

import (
	"context"
	"github.com/advanced-go/stdlib/core"
	"github.com/advanced-go/stdlib/httpx"
	"github.com/advanced-go/stdlib/json"
	"net/http"
	"net/url"
)

// Put - resource PUT
func Put(r *http.Request) (*http.Response, *core.Status) {
	if r == nil {
		return httpx.NewResponseWithStatus(core.NewStatus(http.StatusBadRequest), nil)
	}
	return put[core.Log, *http.Request](r.Context(), r.Header, r.URL.Query(), r)
}

type putBodyConstraints interface {
	[]entryV1 | []entryV2 | []byte | *http.Request
}

func put[E core.ErrorHandler, T putBodyConstraints](ctx context.Context, h http.Header, values url.Values, body T) (*http.Response, *core.Status) {
	//var e E

	/*
		switch strings.ToUpper(r.Method) {
		case http.MethodPut:
			entries, status := createEntries(r.Header, r.Body)
			if !status.OK() {
				e.Handle(status, core.RequestId(r.Header))
				return httpx.NewErrorResponseWithStatus(status)
			}
			if len(entries) == 0 {
				status = core.NewStatusError(core.StatusInvalidContent, errors.New("error: no entries found"))
				e.Handle(status, core.RequestId(r.Header))
				return httpx.NewErrorResponseWithStatus(status)
			}
			//status = addEntriesV1(r.Context(), entries)
			//if !status.OK() {
			//	e.Handle(status, core.RequestId(r.Header))
			//}
			return httpx.NewResponseWithStatus(core.StatusOK(), "")
		default:
			return httpx.NewErrorResponseWithStatus(core.NewStatus(http.StatusMethodNotAllowed))
		}

	*/
	return nil, nil
}

func createEntries[T putBodyConstraints, E entryConstraints](h http.Header, body T) (entries []E, status *core.Status) {
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
	default:
		return nil, core.NewStatusError(core.StatusInvalidContent, core.NewInvalidBodyTypeError(body))
	}
	return entries, core.StatusOK()
}
