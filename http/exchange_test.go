package http

import (
	"fmt"
	resiliency1 "github.com/advanced-go/guidance/resiliency1"
	"github.com/advanced-go/stdlib/httpx"
	"github.com/advanced-go/stdlib/json"
	"net/http"
)

func ExampleExchange_Invalid() {
	resp, status := Exchange(nil)
	fmt.Printf("test: Exchange(nil) -> [status:%v] [content-type:%v] [content-length:%v]\n", status, resp.Header.Get(httpx.ContentType), resp.ContentLength)

	//Output:
	//test: Exchange(nil) -> [status:Bad Request] [content-type:text/plain charset=utf-8] [content-length:0]

}

func ExampleExchange_Resiliency() {
	uri := "http://localhost:8081/github/advanced-go/guidance:v1/resiliency?region=region1"
	req, _ := http.NewRequest(http.MethodGet, uri, nil)

	resp, status := Exchange(req)
	if !status.OK() {
		fmt.Printf("test: Exchange() -> [status:%v]\n", status)
	} else {
		entries, status1 := json.New[[]resiliency1.Entry](resp.Body, resp.Header)
		fmt.Printf("test: Exchange() -> [status:%v] [status-code:%v] [header:%v] [bytes:%v] [content:%v]\n", status1, resp.StatusCode, resp.Header, resp.ContentLength, entries)
	}

	//Output:
	//test: Exchange() -> [status:OK] [status-code:200] [header:map[Content-Type:[application/json]]] [bytes:343] [content:[{region1 Zone1  www.host1.com active   100 125 25} {region1 Zone2  www.host2.com inactive   250 100 10}]]

}
