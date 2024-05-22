package resiliency

import (
	"fmt"
	"github.com/advanced-go/stdlib/core"
	"github.com/advanced-go/stdlib/httpx"
	"net/http"
	"net/url"
)

func ExampleGet() {
	resp, status := Get(nil, nil, nil)
	fmt.Printf("test: Get(nil,nil,nil) -> [status:%v] [content-type:%v] [content-length:%v]\n", status, resp.Header.Get(httpx.ContentType), resp.ContentLength)

	// Empty version defaults to v1
	values := make(url.Values)
	values.Add(httpx.ContentLocation, emptyEntryPath)
	resp, status = Get(nil, nil, values)
	fmt.Printf("test: Get(nil,nil,values) -> [status:%v] [content-type:%v] [content-length:%v]\n", status, resp.Header.Get(httpx.ContentType), resp.ContentLength)

	h := make(http.Header)
	h.Add(core.XVersion, "v5")
	resp, status = Get(nil, h, values)
	fmt.Printf("test: Get(nil,nil,values) -> [status:%v] [content-type:%v] [content-length:%v]\n", status, resp.Header.Get(httpx.ContentType), resp.ContentLength)

	h.Set(core.XVersion, "v1")
	values.Set(httpx.ContentLocation, entryV1Path)
	resp, status = Get(nil, h, values)
	fmt.Printf("test: Get(nil,nil,values) -> [status:%v] [content-type:%v] [content-length:%v]\n", status, resp.Header.Get(httpx.ContentType), resp.ContentLength)

	h.Set(core.XVersion, "v2")
	values.Set(httpx.ContentLocation, entryV2Path)
	resp, status = Get(nil, h, values)
	fmt.Printf("test: Get(nil,nil,values) -> [status:%v] [content-type:%v] [content-length:%v]\n", status, resp.Header.Get(httpx.ContentType), resp.ContentLength)

	//Output:
	//test: Get(nil,nil,nil) -> [status:Bad Request] [content-type:text/plain charset=utf-8] [content-length:32]
	//test: Get(nil,nil,values) -> [status:Not Found] [content-type:] [content-length:0]
	//test: Get(nil,nil,values) -> [status:Bad Request [invalid version: [v5]]] [content-type:text/plain charset=utf-8] [content-length:21]
	//test: Get(nil,nil,values) -> [status:OK] [content-type:application/json] [content-length:618]
	//test: Get(nil,nil,values) -> [status:OK] [content-type:application/json] [content-length:881]

}
