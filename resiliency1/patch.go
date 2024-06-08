package resiliency1

import (
	"context"
	"github.com/advanced-go/stdlib/core"
	"github.com/advanced-go/stdlib/httpx"
	"net/http"
)

func patch[E core.ErrorHandler](ctx context.Context, h http.Header, body *httpx.Patch) (http.Header, *core.Status) {
	return nil, core.StatusOK()
}

func patchProcess(_ *http.Request, item *[]Entry, patch *httpx.Patch) *core.Status {
	if item == nil || patch == nil {
		return core.NewStatus(http.StatusBadRequest)
	}
	for _, op := range patch.Updates {
		switch op.Op {
		case httpx.OpReplace:
			if op.Path == core.HostKey {
				if s, ok1 := op.Value.(string); ok1 {
					(*item)[0].Host = s
				}
			}
		default:
		}
	}
	return core.StatusOK()
}
