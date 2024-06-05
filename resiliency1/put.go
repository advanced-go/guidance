package resiliency1

import (
	"bytes"
	"context"
	"github.com/advanced-go/guidance/module"
	"github.com/advanced-go/stdlib/core"
	"github.com/advanced-go/stdlib/httpx"
	json2 "github.com/advanced-go/stdlib/json"
	"github.com/advanced-go/stdlib/uri"
	"io"
	"net/http"
)

func put[E core.ErrorHandler](ctx context.Context, h http.Header, body []Entry) (http.Header, *core.Status) {
	var e E

	h2 := make(http.Header)
	if len(body) == 0 {
		return h2, core.StatusOK()
	}
	url := uri.Expansion("", module.DocumentsPath, module.DocumentsV1, nil)
	rc, _, status := createReadCloser(body)
	if !status.OK() {
		e.Handle(status, core.RequestId(h))
		h2.Add(httpx.ContentType, httpx.ContentTypeText)
		return h2, status
	}
	req, _ := http.NewRequestWithContext(ctx, http.MethodPut, url, rc)
	httpx.Forward(req.Header, h, core.XAuthority)
	_, status = httpx.DoExchange(req)
	if !status.OK() {
		h2.Add(httpx.ContentType, httpx.ContentTypeText)
		e.Handle(status, core.RequestId(h))
	}
	return h2, status
}

func createReadCloser(body any) (io.ReadCloser, int64, *core.Status) {
	switch ptr := body.(type) {
	case []Entry:
		return json2.NewReadCloser(body)
	case []byte:
		return io.NopCloser(bytes.NewReader(ptr)), int64(len(ptr)), core.StatusOK()
	default:
		return nil, 0, core.NewStatusError(core.StatusInvalidArgument, core.NewInvalidBodyTypeError(body))
	}
}
