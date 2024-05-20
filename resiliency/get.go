package resiliency

import (
	"context"
	"errors"
	"fmt"
	"github.com/advanced-go/stdlib/core"
	"github.com/advanced-go/stdlib/httpx"
	"net/http"
	"net/url"
)

const (
	version1 = "v1"
	version2 = "v2"
)

// http://localhost:8081/github/advanced-go/guidance:resiliency?reg=us&az=dallas&sz=dfwocp1&host=www.google.com

func Get[E core.ErrorHandler](ctx context.Context, version string, h http.Header, values url.Values) (*http.Response, *core.Status) {
	var e E

	/*
		if r == nil {
			return httpx.NewErrorResponseWithStatus(core.NewStatus(http.StatusBadRequest))
		}
		ver, _, status := httpx.ValidateRequestURL(r, module.Authority)
		if !status.OK() {
			return httpx.NewErrorResponseWithStatus(status)
		}
		if ver == "" {
			ver = version1
		}

	*/

	switch version {
	case version1:
		_, status1 := getEntries[entryV1](ctx, values)
		if !status1.OK() {
			e.Handle(status1, core.RequestId(h))
		}
		return nil, core.StatusOK()
	default:
		return httpx.NewErrorResponseWithStatus(core.NewStatusError(http.StatusBadRequest, errors.New(fmt.Sprintf("invalid version: [%v]", version))))
	}
}

/*
func getV1[E core.ErrorHandler](r *http.Request) (*http.Response, *core.Status) {
	if r == nil {
		return httpx.NewErrorResponseWithStatus(core.NewStatus(http.StatusBadRequest))
	}
	entries, status := getEntriesV1(r.Context(), r.URL.Query())
	switch status.Code {
	case http.StatusNotFound, http.StatusGatewayTimeout, core.StatusDeadlineExceeded:
		return httpx.NewResponseWithStatus(status, "")
	case http.StatusOK:
		var resp *http.Response
		if entries != nil {

		}
		return resp, status
	default:
		var e E
		e.Handle(status, core.RequestId(r))
		return httpx.NewErrorResponseWithStatus(status)
	}
}


*/
