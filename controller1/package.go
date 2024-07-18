package controller1

import (
	"context"
	"errors"
	"github.com/advanced-go/stdlib/core"
	json2 "github.com/advanced-go/stdlib/json"
	"net/http"
	"net/url"
)

const (
	PkgPath = "github/advanced-go/observation/controller1"
)

// Get - resource GET
func Get(ctx context.Context, h http.Header, values url.Values) (entries []Entry, status *core.Status) {
	return []Entry{lastEntry()}, core.StatusOK()
}

// Put - resource PUT, with optional content override
func Put(r *http.Request, body []Entry) *core.Status {
	if r == nil {
		return core.NewStatusError(core.StatusInvalidArgument, errors.New("error: request is nil"))
	}
	if body == nil {
		content, status := json2.New[[]Entry](r.Body, r.Header)
		if !status.OK() {
			var e core.Log
			e.Handle(status, core.RequestId(r.Header))
			return status
		}
		body = content
	}
	return core.StatusOK() //put[core.Log](r.Context(), core.AddRequestId(r.Header), inferenceResource, "", body, nil)
}
