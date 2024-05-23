package resiliency

import (
	"context"
	"fmt"
	"github.com/advanced-go/stdlib/controller"
	"github.com/advanced-go/stdlib/core"
	"github.com/advanced-go/stdlib/httpx"
	"github.com/advanced-go/stdlib/json"
	"net/http"
	"net/url"
	"strings"
	"time"
)

const (
	authority = "github/advanced-go/documents"
)

var (
	testEntry = []Entry{
		{Origin: core.Origin{Region: "region1", Zone: "Zone1", Host: "www.host1.com"}, Status: "active", Timeout: "100ms", RateLimit: "125", RateBurst: "25"},
		{Origin: core.Origin{Region: "region1", Zone: "Zone2", Host: "www.host2.com"}, Status: "inactive", Timeout: "250ms", RateLimit: "100", RateBurst: "10"},
		{Origin: core.Origin{Region: "region2", Zone: "Zone1", Host: "www.google.com"}, Status: "removed", Timeout: "3s", RateLimit: "50", RateBurst: "5"},
	}

	rsc = NewResource[Entry]("github/advanced-go/documents",
		func(item any, values url.Values) bool {
			filter := core.NewOrigin(values)
			if entry, ok := item.(*Entry); ok {
				if core.OriginMatch(entry.Origin, filter) {
					return true
				}
			}
			return false
		},
		func(item any, r *http.Request) {

		})
)

type Resource[T any] struct {
	List      []T
	Authority *http.Response
	Match     func(item any, values url.Values) bool
	Patch     func(item any, r *http.Request)
}

func NewResource[T any](authority string, match func(item any, values url.Values) bool, patch func(item any, r *http.Request)) *Resource[T] {
	r := new(Resource[T])
	r.Authority = httpx.NewAuthorityResponse(authority)
	r.Match = match
	r.Patch = patch
	return r
}

func (r *Resource[T]) Length() int {
	return len(r.List)
}

func (r *Resource[T]) Get(values url.Values) (items []T, status *core.Status) {
	if r.Match == nil {
		return nil, core.NewStatus(core.StatusInvalidArgument)
	}
	for _, target := range r.List {
		if r.Match(&target, values) {
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

func (r *Resource[T]) Do(req *http.Request) (*http.Response, *core.Status) {
	//_, path, status1 := httpx.ValidateRequestURL(r, module.Authority)
	//if !status1.OK() {
	//	return httpx.NewResponseWithStatus(status1, status1.Err)
	//}
	if strings.Contains(req.URL.Path, core.AuthorityPath) {
		return r.Authority, core.StatusOK()
	}
	switch req.Method {
	case http.MethodGet:
		items, status := r.Get(req.URL.Query())
		if !status.OK() {
			return httpx.NewResponseWithStatus(status, nil)
		}
		resp, status1 := httpx.NewJsonResponse(items, req.Header)
		if !status1.OK() {
			return httpx.NewResponseWithStatus(status, status.Err)
		}
		return resp, core.StatusOK()
	case http.MethodPut:
		items, status := json.New[[]T](req.Body, req.Header)
		if !status.OK() {
			return httpx.NewResponseWithStatus(status, nil)
		}
		if len(items) == 0 {
			return httpx.NewResponseWithStatus(core.StatusNotFound(), nil)
		}
		r.Put(items)
		return httpx.NewResponseWithStatus(core.StatusOK(), nil)
	case http.MethodPatch:
		if r.Patch == nil {
			return httpx.NewResponseWithStatus(core.NewStatus(core.StatusInvalidArgument), nil)
		}
		if r.Match == nil {
			return httpx.NewResponseWithStatus(core.NewStatus(core.StatusInvalidArgument), nil)
		}
		for _, target := range r.List {
			if r.Match(&target, req.URL.Query()) {
				r.Patch(&target, req)
			}
		}
		return httpx.NewResponseWithStatus(core.StatusOK(), nil)
	default:
		return httpx.NewResponseWithStatus(core.StatusNotFound(), nil)
	}
}

func init() {
	ctrl := controller.NewController("resiliencyv1", controller.NewPrimaryResource("", authority, time.Second*2, "", rsc.Do), nil)
	controller.RegisterController(ctrl)
}

func ExampleExchange_Simple() {
	status := put[core.Output](context.Background(), nil, testEntry)
	fmt.Printf("test: put() -> [status:%v] [count:%v]\n", status, rsc.Length())

	values := make(url.Values)
	values.Add(core.ZoneKey, "zone1")
	docs, status1 := get[core.Output](context.Background(), nil, values)
	fmt.Printf("test: get() -> [status:%v] [count:%v]\n", status1, len(docs))

	//Output:
	//test: put() -> [status:OK] [count:3]
	//test: get() -> [status:OK] [count:2]

}
