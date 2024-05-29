package resiliency

import (
	"context"
	"github.com/advanced-go/stdlib/core"
	"net/http"
)

func post[E core.ErrorHandler](ctx context.Context, h http.Header, data *PostData) *core.Status {
	return core.StatusOK()
}
