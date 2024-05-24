package resiliency

import (
	"bytes"
	"fmt"
	"github.com/advanced-go/stdlib/controller"
	"github.com/advanced-go/stdlib/core"
	"github.com/advanced-go/stdlib/httpx"
	json2 "github.com/advanced-go/stdlib/json"
	"io"
	"net/http"
	"net/url"
	"time"
)

const (
	originAuthority = "github/advanced-go/origin-resource"
)

var (
	testOrigins = []core.Origin{
		{Region: "region1", Zone: "Zone1", Host: "www.host1.com"},
		{Region: "region1", Zone: "Zone2", Host: "www.host2.com"},
		{Region: "region2", Zone: "Zone1", Host: "www.google.com"},
	}
	originRsc = NewResource[core.Origin](originAuthority, originMatch, originPatch)
)

func originMatch(item any, values url.Values) bool {
	filter := core.NewOrigin(values)
	if entry, ok := item.(*Entry); ok {
		if core.OriginMatch(entry.Origin, filter) {
			return true
		}
	}
	return false
}

func originPatch(item any, patch *httpx.Patch) {
	if item == nil || patch == nil {
		return
	}
	if target, ok := item.(*core.Origin); ok {
		for _, op := range patch.Updates {
			switch op.Op {
			case httpx.OpReplace:
				if op.Path == "Host" {
					if s, ok1 := op.Value.(string); ok1 {
						target.Host = s
					}
				}
			default:
			}
		}
	}
}

func init() {
	ctrl := controller.NewController("origin-resource", controller.NewPrimaryResource("localhost", originAuthority, time.Second*2, "", originRsc.Do), nil)
	controller.RegisterController(ctrl)
}

func createOriginReadCloser(body any) (io.ReadCloser, int64, *core.Status) {
	switch ptr := body.(type) {
	case []core.Origin:
		return json2.NewReadCloser(body)
	case []byte:
		return io.NopCloser(bytes.NewReader(ptr)), int64(len(ptr)), core.StatusOK()
	default:
		return nil, 0, core.NewStatus(http.StatusBadRequest)
	}
}

func ExampleOriginResource() {
	url := originAuthority + ":resiliency"
	rc, _, status0 := createOriginReadCloser(testOrigins)
	fmt.Printf("test: createReaderCloser() -> [status:%v]\n", status0)

	req, _ := http.NewRequest(http.MethodPut, url, rc)
	resp, status := httpx.DoExchange(req)

	fmt.Printf("test: DoExchange() -> [status:%v] [resp:%v]\n", status, resp != nil)

	//Output:
	//fail

}
