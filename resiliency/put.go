package resiliency

import (
	"context"
	"errors"
	"fmt"
	"github.com/advanced-go/guidance/module"
	"github.com/advanced-go/stdlib/core"
	"net/http"
)

// Put - resource PUT
func Put[T PutBodyConstraints](ctx context.Context, h http.Header, body T) *core.Status {
	if body == nil {
		return core.NewStatus(http.StatusBadRequest)
	}
	return put[core.Log, *http.Request](ctx, core.AddRequestId(h), body)
}

func put[E core.ErrorHandler](ctx context.Context, h http.Header, body any) *core.Status {
	var e E

	switch h.Get(core.XVersion) {
	case module.Ver1, "":
		entries, status := createEntries[EntryV1](h, body)
		if !status.OK() {
			e.Handle(status, core.RequestId(h))
			return status
		}
		status = addEntries[EntryV1](ctx, h, entries)
		if !status.OK() {
			e.Handle(status, core.RequestId(h))
		}
		return status
	case module.Ver2:
		entries, status := createEntries[EntryV2](h, body)
		if !status.OK() {
			return status
		}
		return addEntries[EntryV2](ctx, h, entries)
	default:
		return core.NewStatusError(http.StatusBadRequest, errors.New(fmt.Sprintf("invalid version: [%v]", h.Get(core.XVersion))))
	}
}
