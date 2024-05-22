package http

import (
	"errors"
	"fmt"
	"github.com/advanced-go/guidance/resiliency"
	"github.com/advanced-go/stdlib/core"
	"github.com/advanced-go/stdlib/httpx"
	"net/http"
)

func resiliencyExchange(r *http.Request) (*http.Response, *core.Status) {
	switch r.Method {
	case http.MethodGet:
		return resiliency.Get(r.Context(), r.Header, r.URL.Query())
	case http.MethodPut:
		return resiliency.Put(r)
	default:
		status := core.NewStatusError(http.StatusBadRequest, errors.New(fmt.Sprintf("error invalid method: [%v]", r.Method)))
		return httpx.NewResponseWithStatus(status, status.Err)
	}
}
