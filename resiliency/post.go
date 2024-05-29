package resiliency

import (
	"context"
	"github.com/advanced-go/stdlib/core"
	"net/http"
	"net/url"
)

func post[E core.ErrorHandler, T PutBodyConstraints](ctx context.Context, h http.Header, values url.Values, body T) *core.Status {

	return core.StatusOK()
}
