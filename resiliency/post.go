package resiliency

import (
	"errors"
	"github.com/advanced-go/stdlib/core"
	"github.com/advanced-go/stdlib/httpx"
	"github.com/advanced-go/stdlib/json"
	"io"
	"net/http"
	"net/url"
	"strings"
)

// PostConstraints - Post constraints
type PostEntryConstraints interface {
	[]entryV1 | []byte | *http.Request
}

func Post[E core.ErrorHandler](h http.Header, method string, values url.Values, body *http.Request) (*http.Response, *core.Status) {
	var e E

	if r == nil {
		return httpx.NewErrorResponseWithStatus(core.NewStatus(http.StatusBadRequest))
	}
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
}

func createEntries[T entryConstraints](h http.Header, body io.ReadCloser) (entries []T, status *core.Status) {
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
