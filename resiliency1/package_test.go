package resiliency1

import (
	"context"
	"fmt"
	"github.com/advanced-go/guidance/module"
	"github.com/advanced-go/stdlib/core"
	"github.com/advanced-go/stdlib/httpx"
	"github.com/advanced-go/stdlib/uri"
	"net/http"
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
					(*item)[0].Host = s
				}
			}
		default:
		}
	}
	return core.StatusOK()
}

func ExampleExchange_GetAll() {
	values := uri.BuildValues("region=*")
	h := make(http.Header)
	h.Add(core.XAuthority, module.Authority)
	docs1, h, status1 := Get(context.Background(), h, values)
	fmt.Printf("test: get() -> [status:%v] [header:%v] [count:%v]\n", status1, h, len(docs1))

	//Output:
	//test: get() -> [status:OK] [header:map[]] [count:3]

}
