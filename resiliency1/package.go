package resiliency1

import (
	"context"
	"errors"
	"github.com/advanced-go/guidance/module"
	"github.com/advanced-go/stdlib/controller"
	"github.com/advanced-go/stdlib/core"
	"github.com/advanced-go/stdlib/httpx"
	json2 "github.com/advanced-go/stdlib/json"
	"github.com/advanced-go/stdlib/uri"
	"net/http"
	"net/url"
	"time"
)

const (
	PkgPath   = "github/advanced-go/guidance/resiliency1"
	routeName = "documents-resiliency"
	hostKey   = "docs-host"
)

var resolver = uri.NewResolver([]uri.HostEntry{{Key: hostKey, Host: "www.documents.com", Proxy: false}})

// EgressRoute - upstream egress traffic route configuration
func EgressRoute() *controller.Config {
	return &controller.Config{RouteName: routeName, Host: resolver.Host(hostKey), Authority: module.DocumentsAuthority, LivenessPath: core.HealthLivenessPath, Duration: time.Second * 2}
}

// Get - resource GET
func Get(ctx context.Context, h http.Header, values url.Values) ([]Entry, http.Header, *core.Status) {
	return get[core.Log](ctx, core.AddRequestId(h), values)
}

// Delete - resource DELETE
func Delete(ctx context.Context, h http.Header, values url.Values) (http.Header, *core.Status) {
	return delete[core.Log](ctx, core.AddRequestId(h), values)
}

// Put - resource PUT, with optional content override
func Put(r *http.Request, body []Entry) (http.Header, *core.Status) {
	if r == nil {
		return nil, core.NewStatusError(core.StatusInvalidArgument, errors.New("error: request is nil"))
	}
	if body == nil {
		content, status := json2.New[[]Entry](r.Body, r.Header)
		if !status.OK() {
			var e core.Log
			e.Handle(status, core.RequestId(r.Header))
			return nil, status
		}
		body = content
	}
	return put[core.Log](r.Context(), core.AddRequestId(r.Header), body)
}

// Post - resource POST, with optional content override
func Post(r *http.Request, body *PostData) (http.Header, *core.Status) {
	if r == nil {
		return nil, core.NewStatusError(core.StatusInvalidArgument, errors.New("error: request is nil"))
	}
	if body == nil {
		content, status := json2.New[PostData](r.Body, r.Header)
		if !status.OK() {
			var e core.Log
			e.Handle(status, core.RequestId(r.Header))
			return nil, status
		}
		body = &content
	}
	return post[core.Log](r.Context(), core.AddRequestId(r.Header), body)
}

// Patch - resource PATCH, with optional content override
func Patch(r *http.Request, body *httpx.Patch) (http.Header, *core.Status) {
	if r == nil {
		return nil, core.NewStatusError(core.StatusInvalidArgument, errors.New("error: request is nil"))
	}
	if body == nil {
		content, status := json2.New[httpx.Patch](r.Body, r.Header)
		if !status.OK() {
			var e core.Log
			e.Handle(status, core.RequestId(r.Header))
			return nil, status
		}
		body = &content
	}
	return patch[core.Log](r.Context(), core.AddRequestId(r.Header), body)
}
