package http

import (
	"fmt"
	"github.com/advanced-go/stdlib/httpx"
)

func ExampleExchange_Invalid() {
	resp, status := Exchange(nil)
	fmt.Printf("test: Exchange(nil) -> [status:%v] [content-type:%v] [content-length:%v]\n", status, resp.Header.Get(httpx.ContentType), resp.ContentLength)

	//Output:
	//test: Exchange(nil) -> [status:Bad Request] [content-type:] [content-length:0]
	
}
