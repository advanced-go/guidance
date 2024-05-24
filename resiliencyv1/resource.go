package resiliency

import (
	"errors"
	"fmt"
	"github.com/advanced-go/stdlib/core"
	"github.com/advanced-go/stdlib/httpx"
	"github.com/advanced-go/stdlib/json"
	"net/http"
	"net/url"
	"strings"
)

type Resource[T any] struct {
	List      []T
	Authority *http.Response
	MatchFn   func(item any, values url.Values) bool
	PatchFn   func(item any, patch *httpx.Patch)
}

func NewResource[T any](authority string, match func(item any, values url.Values) bool, patch func(item any, patch *httpx.Patch)) *Resource[T] {
	r := new(Resource[T])
	r.Authority = httpx.NewAuthorityResponse(authority)
	r.MatchFn = match
	r.PatchFn = patch
	return r
}

func (r *Resource[T]) Count() int {
	return len(r.List)
}

func (r *Resource[T]) Empty() {
	r.List = nil
}

func (r *Resource[T]) Get(values url.Values) (items []T, status *core.Status) {
	if r.MatchFn == nil {
		return nil, core.NewStatusError(core.StatusInvalidArgument, errors.New("MatchFunc() is nil"))
	}
	for _, target := range r.List {
		if r.MatchFn(&target, values) {
			items = append(items, target)
		}
	}
	if len(items) == 0 {
		return nil, core.StatusNotFound()
	}
	return items, core.StatusOK()
}

func (r *Resource[T]) Put(items []T) *core.Status {
	if len(items) > 0 {
		r.List = append(r.List, items...)
	}
	return core.StatusOK()
}

func (r *Resource[T]) Patch(values url.Values, patch *httpx.Patch) *core.Status {
	if r.MatchFn == nil {
		return core.NewStatusError(core.StatusInvalidArgument, errors.New("MatchFunc() is nil"))
	}
	if r.PatchFn == nil {
		return core.NewStatusError(core.StatusInvalidArgument, errors.New("PatchFunc() is nil"))
	}
	for _, target := range r.List {
		if r.MatchFn(&target, values) {
			r.PatchFn(&target, patch)
		}
	}
	return core.StatusOK()
}

func (r *Resource[T]) Delete(values url.Values) *core.Status {
	if r.MatchFn == nil {
		return core.NewStatusError(core.StatusInvalidArgument, errors.New("MatchFunc() is nil"))
	}
	for _, target := range r.List {
		if r.MatchFn(&target, values) {
			//delete
		}
	}
	return core.StatusOK()
}

func (r *Resource[T]) Do(req *http.Request) (*http.Response, *core.Status) {
	//_, _, status1 := httpx.ValidateRequestURL(req, module.Authority)
	//if !status1.OK() {
	//	return httpx.NewResponseWithStatus(status1, status1.Err)
	//}
	fmt.Printf("Do() -> [url:%v]\n", req.URL.String())
	if strings.Contains(req.URL.Path, core.AuthorityPath) {
		return r.Authority, core.StatusOK()
	}
	switch req.Method {
	case http.MethodGet:
		items, status := r.Get(req.URL.Query())
		if !status.OK() {
			return httpx.NewResponseWithStatus(status, status.Err)
		}
		resp, status1 := httpx.NewJsonResponse(items, req.Header)
		if !status1.OK() {
			return httpx.NewResponseWithStatus(status, status.Err)
		}
		return resp, core.StatusOK()
	case http.MethodPut:
		items, status := json.New[[]T](req.Body, req.Header)
		if !status.OK() {
			return httpx.NewResponseWithStatus(status, status.Err)
		}
		if len(items) == 0 {
			return httpx.NewResponseWithStatus(core.StatusNotFound(), nil)
		}
		r.Put(items)
		return httpx.NewResponseWithStatus(core.StatusOK(), nil)
	case http.MethodPatch:
		patch, status := json.New[httpx.Patch](req.Body, req.Header)
		if !status.OK() {
			return httpx.NewResponseWithStatus(status, status.Err)
		}
		status = r.Patch(req.URL.Query(), &patch)
		return httpx.NewResponseWithStatus(status, status.Err)
	case http.MethodDelete:
		status := r.Delete(req.URL.Query())
		return httpx.NewResponseWithStatus(status, status.Err)
	default:
		status := core.NewStatusError(http.StatusBadRequest, errors.New(fmt.Sprintf("unsupported method: %v", req.Method)))
		return httpx.NewResponseWithStatus(status, status.Err)
	}
}
