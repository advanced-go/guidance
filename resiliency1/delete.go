package resiliency

import (
	"context"
	"github.com/advanced-go/stdlib/core"
	"net/http"
	"net/url"
)

func delete[E core.ErrorHandler](ctx context.Context, h http.Header, values url.Values) *core.Status {
	//var e E
	//url := module.BuildDocumentsPath(module.Ver1, values)

	return core.StatusOK()
}

