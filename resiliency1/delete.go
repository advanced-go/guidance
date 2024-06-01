package resiliency1

import (
	"context"
	"errors"
	"github.com/advanced-go/stdlib/core"
	"net/http"
	"net/url"
)

func delete[E core.ErrorHandler](ctx context.Context, h http.Header, url *url.URL) *core.Status {
	if url == nil {
		return core.NewStatusError(core.StatusInvalidArgument, errors.New("invalid argument: URL is nil"))
	}

	return core.StatusOK()
}
