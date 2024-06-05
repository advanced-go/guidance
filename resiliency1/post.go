package resiliency1

import (
	"context"
	"github.com/advanced-go/stdlib/core"
	"net/http"
)

func post[E core.ErrorHandler](ctx context.Context, h http.Header, data *PostData) (http.Header, *core.Status) {
	return nil, core.StatusOK()
}
