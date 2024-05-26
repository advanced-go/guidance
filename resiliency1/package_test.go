package resiliency

import (
	"context"
	"fmt"
	"github.com/advanced-go/guidance/module"
	"github.com/advanced-go/stdlib/controller"
	"github.com/advanced-go/stdlib/core"
	"github.com/advanced-go/stdlib/httpx"
	"net/http"
	"net/url"
	"time"
)

const (
	rscName = "test-rsc"
)

var (
	testEntry = []Entry{
		{Origin: core.Origin{Region: "region1", Zone: "Zone1", Host: "www.host1.com"}, Status: "active", Timeout: "100ms", RateLimit: "125", RateBurst: "25"},
		{Origin: core.Origin{Region: "region1", Zone: "Zone2", Host: "www.host2.com"}, Status: "inactive", Timeout: "250ms", RateLimit: "100", RateBurst: "10"},
		{Origin: core.Origin{Region: "region2", Zone: "Zone1", Host: "www.google.com"}, Status: "removed", Timeout: "3s", RateLimit: "50", RateBurst: "5"},
	}
	rsc = httpx.NewResource[Entry, httpx.Patch, struct{}](rscName, match, nil, patchProcess, nil)
)

func match(item *Entry, req *http.Request) bool {
	filter := core.NewOrigin(req.URL.Query())
	if core.OriginMatch(item.Origin, filter) {
		return true
	}
	return false
}

func patchProcess(item *[]Entry, patch *httpx.Patch) *http.Response {
	if item == nil || patch == nil {
		return httpx.NewResponse(core.NewStatus(http.StatusBadRequest), nil)
	}
	for _, op := range patch.Updates {
		switch op.Op {
		case httpx.OpReplace:
			if op.Path == core.HostKey {
				if s, ok1 := op.Value.(string); ok1 {
					(*item)[0].Origin.Host = s
				}
			}
		default:
		}
	}
	return httpx.NewResponse(core.StatusOK(), nil)
}

func init() {
	ctrl := controller.NewController("entry-resource", controller.NewPrimaryResource("localhost:8082", module.DocumentsAuthority, time.Second*2, "", rsc.Do), nil)
	controller.RegisterController(ctrl)
}

func ExampleExchange_PutGet() {
	status := put[core.Output](context.Background(), nil, testEntry)
	cnt := rsc.Count()
	fmt.Printf("test: put() -> [status:%v] [count:%v]\n", status, cnt)

	values := make(url.Values)
	values.Add(core.ZoneKey, "zone1")
	docs, status1 := get[core.Output](context.Background(), nil, values)
	fmt.Printf("test: get() -> [status:%v] [count:%v]\n", status1, len(docs))

	//Output:
	//test: put() -> [status:OK] [count:3]
	//test: get() -> [status:OK] [count:2]

}
