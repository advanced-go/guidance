package resiliency1

import (
	"fmt"
	"github.com/advanced-go/guidance/module"
	"github.com/advanced-go/stdlib/core"
	"github.com/advanced-go/stdlib/uri"
	"net/http"
	"net/url"
)

const (
	getAllResp = "file://[cwd]/resiliency1test/get-all-resp-v1.txt"
)

func ExampleGet() {
	values := make(url.Values)
	h := make(http.Header)
	h.Add(uri.BuildPath(module.DocumentsAuthority, module.DocumentsResourceV1, nil), getAllResp)
	entries, _, status := get[core.Output](nil, h, values)

	fmt.Printf("test: get() -> [status:%v] [entries:%v]\n", status, len(entries))

	//Output:
	//test: get() -> [status:OK] [entries:3]

}
