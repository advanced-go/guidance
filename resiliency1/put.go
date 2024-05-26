package resiliency

import (
	"bytes"
	"context"
	"github.com/advanced-go/guidance/module"
	"github.com/advanced-go/stdlib/core"
	"github.com/advanced-go/stdlib/httpx"
	json2 "github.com/advanced-go/stdlib/json"
	"io"
	"net/http"
)

// Put - resource PUT
func Put[T PutBodyConstraints](ctx context.Context, h http.Header, body T) *core.Status {
	if body == nil {
		return core.NewStatus(http.StatusBadRequest)
	}
	return put[core.Log](ctx, core.AddRequestId(h), body)
}

func put[E core.ErrorHandler](ctx context.Context, h http.Header, body any) *core.Status {
	var e E

	url := module.BuildDocumentsPath(module.Ver1, nil)
	rc, _, status := createReadCloser(body)
	if !status.OK() {
		e.Handle(status, core.RequestId(h))
		return status
	}
	req, _ := http.NewRequestWithContext(ctx, http.MethodPut, url, rc)
	httpx.Forward(req.Header, h)
	_, status = httpx.DoExchange(req)
	if !status.OK() {
		e.Handle(status, core.RequestId(h))
	}
	return status
}

func createReadCloser(body any) (io.ReadCloser, int64, *core.Status) {
	switch ptr := body.(type) {
	case []Entry:
		return json2.NewReadCloser(body)
	case []byte:
		return io.NopCloser(bytes.NewReader(ptr)), int64(len(ptr)), core.StatusOK()
	default:
		return nil, 0, core.NewStatus(http.StatusBadRequest)
	}
}
