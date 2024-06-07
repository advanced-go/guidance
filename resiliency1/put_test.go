package resiliency1

import (
	"fmt"
	"github.com/advanced-go/guidance/module"
	"github.com/advanced-go/stdlib/core"
	"github.com/advanced-go/stdlib/uri"
	"net/http"
)

const (
	putResp = "file://[cwd]/resiliency1test/put-resp-v1.txt"
)

func ExamplePut() {
	h := make(http.Header)
	h.Add(uri.BuildPath(module.DocumentsAuthority, module.DocumentsResourceV1, nil), putResp)

	_, status := put[core.Output](nil, h, nil)
	fmt.Printf("test: put(nil,h,nil) -> [status:%v]\n", status)

	_, status = put[core.Output](nil, h, []Entry{{Region: "us-west"}})
	fmt.Printf("test: put(nil,h,[]Entry) -> [status:%v]\n", status)

	//Output:
	//test: put(nil,h,nil) -> [status:OK]
	//test: put(nil,h,[]Entry) -> [status:I'm A Teapot]

}
