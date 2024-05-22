package http

import (
	"fmt"
	"github.com/advanced-go/stdlib/httpx"
)

func ExampleExchange_Invalid() {
	resp, status := Exchange(nil)
	fmt.Printf("test: Exchange(nil) -> [status:%v] [content-type:%v] [content-length:%v]\n", status, resp.Header.Get(httpx.ContentType), resp.ContentLength)

	//Output:
	//test: Exchange(nil) -> [status:Invalid Argument [error: request is nil]] [content-type:text/plain charset=utf-8] [content-length:21]

}
