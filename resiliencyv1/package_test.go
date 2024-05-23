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

var testEntry = []Entry{
	{Origin: core.Origin{Region: "region1", Zone: "Zone1", Host: "www.host1.com"}, Status: "active", Timeout: "100ms", RateLimit: "125", RateBurst: "25"},
	{Origin: core.Origin{Region: "region1", Zone: "Zone2", Host: "www.host2.com"}, Status: "inactive", Timeout: "250ms", RateLimit: "100", RateBurst: "10"},
	{Origin: core.Origin{Region: "region2", Zone: "Zone1", Host: "www.google.com"}, Status: "removed", Timeout: "3s", RateLimit: "50", RateBurst: "5"},
}

type documents struct {
	list []Entry
}

var (
	authorityResponse = httpx.NewAuthorityResponse("github/advanced-go/documents")
	store             = new(documents)
)

func (d *documents) get(values url.Values) (docs []Entry, status *core.Status) {
	filter := core.NewOrigin(values)
	for _, target := range d.list {
		if core.OriginMatch(target.Origin, filter) {
			docs = append(docs, target)
		}
	}
	if len(docs) == 0 {
		return nil, core.StatusNotFound()
	}
	return docs, core.StatusOK()
}

func (d *documents) len() int {
	return len(d.list)
}

func (d *documents) add(docs []Entry) *core.Status {
	if len(docs) > 0 {
		d.list = append(d.list, docs...)
	}
	return core.StatusOK()
}

func init() {
	ctrl := controller.NewController("resiliencyv1", controller.NewPrimaryResource("", authority, time.Second*2, "", do), nil)
	controller.RegisterController(ctrl)
}

func do(r *http.Request) (*http.Response, *core.Status) {
	//_, path, status1 := httpx.ValidateRequestURL(r, module.Authority)
	//if !status1.OK() {
	//	return httpx.NewResponseWithStatus(status1, status1.Err)
	//}
	if strings.Contains(r.URL.Path, core.AuthorityPath) {
		return authorityResponse, core.StatusOK()
	}
	switch r.Method {
	case http.MethodGet:
		// Need to add authority
		docs, status := store.get(r.URL.Query())
		if !status.OK() {
			return httpx.NewResponseWithStatus(status, nil)
		}
		resp, status1 := httpx.NewJsonResponse(docs, r.Header)
		if !status1.OK() {
			return httpx.NewResponseWithStatus(status, status.Err)
		}
		return resp, core.StatusOK()
	case http.MethodPut:
		docs, status := json.New[[]Entry](r.Body, r.Header)
		if !status.OK() {
			return httpx.NewResponseWithStatus(status, nil)
		}
		store.add(docs)
		return httpx.NewResponseWithStatus(core.StatusOK(), nil)
	default:
		return httpx.NewResponseWithStatus(core.StatusNotFound(), nil)
	}
}

func ExampleExchange_Simple() {
	status := put[core.Output](context.Background(), nil, testEntry)
	fmt.Printf("test: put() -> [status:%v] [count:%v]\n", status, store.len())

	values := make(url.Values)
	values.Add(core.ZoneKey, "zone1")
	docs, status1 := get[core.Output](context.Background(), nil, values)
	fmt.Printf("test: get() -> [status:%v] [count:%v]\n", status1, len(docs))

	//Output:
	//test: put() -> [status:OK] [count:3]
	//test: get() -> [status:OK] [count:2]

}
