package resiliency1

import (
	"context"
	"github.com/advanced-go/stdlib/core"
	"github.com/advanced-go/stdlib/httpx"
	"net/http"
)

func patch[E core.ErrorHandler](ctx context.Context, h http.Header, body *httpx.Patch) (http.Header, *core.Status) {
	h2 := make(http.Header)
	return h2, core.StatusOK()
}
