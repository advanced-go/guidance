package http

import (
	"errors"
	"fmt"
	resiliency2 "github.com/advanced-go/guidance/resiliency2"
	"github.com/advanced-go/stdlib/core"
	"github.com/advanced-go/stdlib/httpx"
	"net/http"
)

func resiliency2Exchange(r *http.Request) (resp *http.Response, status *core.Status) {
	var e core.Log
	var entries any

	switch r.Method {
	case http.MethodGet:
		entries, status = resiliency2.Get(r.Context(), r.Header, r.URL)
		if status.NotFound() || status.Timeout() {
			return httpx.NewResponseWithStatus(status, status.Err)
		}
		resp, status = httpx.NewJsonResponse(entries, r.Header)
		if !status.OK() {
			e.Handle(status, core.RequestId(r.Header))
		}
		return httpx.NewResponseWithStatus(status, status.Err)
	case http.MethodPut:
		status = resiliency2.Put(r, nil)
		return httpx.NewResponseWithStatus(status, status.Err)
	default:
		status = core.NewStatusError(http.StatusBadRequest, errors.New(fmt.Sprintf("error invalid method: [%v]", r.Method)))
		return httpx.NewResponseWithStatus(status, status.Err)
	}
}
