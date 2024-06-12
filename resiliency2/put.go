package resiliency2

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

func put[E core.ErrorHandler](ctx context.Context, h http.Header, body []Entry) *core.Status {
	var e E

	url := uri.Expansion("", module.DocumentsAuthorityV2, module.DocumentsResourceV2, nil)
	rc, _, status := createReadCloser(body)
	if !status.OK() {
		e.Handle(status, core.RequestId(h))
		return status
	}
	req, err := http.NewRequestWithContext(ctx, http.MethodPut, url, rc)
	if err != nil {
		return core.NewStatusError(core.StatusInvalidArgument, err)
	}
	req.Header.Set(core.XFrom, module.Authority)
	httpx.Forward(req.Header, h, core.XAuthority)
	_, status = httpx.Exchange(req)
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
