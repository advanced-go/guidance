package resiliency

import (
	"context"
	"fmt"
	"github.com/advanced-go/stdlib/core"
	"github.com/advanced-go/stdlib/httpx"
	"net/http"
	"net/url"
)

func patchProcess(_ *http.Request, item *[]Entry, patch *httpx.Patch) *core.Status {
	if item == nil || patch == nil {
		return core.NewStatus(http.StatusBadRequest)
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
	return core.StatusOK()
}

func _ExampleExchange_PutGet() {
	status := put[core.Output](context.Background(), nil, testEntry)
	cnt := entryRsc.Count()
	fmt.Printf("test: put() -> [status:%v] [count:%v]\n", status, cnt)

	values := make(url.Values)
	values.Add(core.ZoneKey, "zone1")
	docs, status1 := get[core.Output](context.Background(), nil, values)
	fmt.Printf("test: get() -> [status:%v] [count:%v]\n", status1, len(docs))

	//Output:
	//test: put() -> [status:OK] [count:3]
	//test: get() -> [status:OK] [count:2]

}

func ExampleExchange_Get() {
	values := make(url.Values)
	values.Add(core.RegionKey, "*")
	docs, status1 := get[core.Output](context.Background(), nil, values)
	fmt.Printf("test: get() -> [status:%v] [count:%v]\n", status1, len(docs))

	//Output:
	//test: get() -> [status:OK] [count:3]

}

func ExampleURL() {
	uri := "resiliency/controller/timeout?region=region1"
	u, _ := url.Parse(uri)

	fmt.Printf("test: url.Parse(\"%v\") -> [path:%v] [query:%v]\n", uri, u.Path, u.Query())

	//Output:
	//test: url.Parse("resiliency/controller/timeout?region=region1") -> [path:resiliency/controller/timeout] [query:map[region:[region1]]]

}
