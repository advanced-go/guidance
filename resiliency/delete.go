package resiliency

import (
	"context"
	"github.com/advanced-go/stdlib/core"
	"github.com/advanced-go/stdlib/httpx"
	"net/http"
	"net/url"
)

// Delete - resource DELETE
func Delete(ctx context.Context, h http.Header, values url.Values) (*http.Response, *core.Status) {
	if h == nil {
		return httpx.NewResponseWithStatus(core.NewStatus(http.StatusBadRequest), nil)
	}
	if values == nil {
		return httpx.NewResponseWithStatus(core.NewStatus(http.StatusBadRequest), nil)
	}
	return httpx.NewNotFoundResponseWithStatus()
}
