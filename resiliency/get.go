package resiliency

import (
	"errors"
	"fmt"
	"github.com/advanced-go/guidance/module"
	"github.com/advanced-go/stdlib/core"
	"github.com/advanced-go/stdlib/httpx"
	"net/http"
)

const (
	version1 = "v1"
	version2 = "v2"
)

// http://localhost:8081/github/advanced-go/guidance:resiliency?reg=us&az=dallas&sz=dfwocp1&host=www.google.com

func Get[E core.ErrorHandler](r *http.Request) (*http.Response, *core.Status) {
	var e E

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
	switch ver {
	case version1:
		entries, status1 := getEntries[entryV1](r.Context(), r.URL.Query())
		if !status1.OK() {

		}
	default:
		return httpx.NewErrorResponseWithStatus(core.NewStatusError(http.StatusBadRequest, errors.New(fmt.Sprintf("invalid version: [%v]", ver))))
	}
}

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
