package resiliency

import (
	"context"
	"fmt"
	"github.com/advanced-go/stdlib/core"
	"github.com/advanced-go/stdlib/httpx"
	"net/http"
	"net/url"
	"reflect"
	"testing"
)

func ExampleGet() {
	resp, status := Get(nil, nil, nil)
	fmt.Printf("test: Get(nil,nil,nil) -> [status:%v] [content-type:%v] [content-length:%v]\n", status, resp.Header.Get(httpx.ContentType), resp.ContentLength)

	// Empty version defaults to v1
	values := make(url.Values)
	resp, status = Get(nil, nil, values)
	fmt.Printf("test: Get(nil,nil,values) -> [status:%v] [content-type:%v] [content-length:%v]\n", status, resp.Header.Get(httpx.ContentType), resp.ContentLength)

	h := make(http.Header)
	h.Add(core.XVersion, "v5")
	resp, status = Get(nil, h, values)
	fmt.Printf("test: Get(nil,nil,values) -> [status:%v] [content-type:%v] [content-length:%v]\n", status, resp.Header.Get(httpx.ContentType), resp.ContentLength)

	h.Set(core.XVersion, "v1")
	values.Add(httpx.ContentLocation, entryV1Path)
	resp, status = Get(nil, h, values)
	fmt.Printf("test: Get(nil,nil,values) -> [status:%v] [content-type:%v] [content-length:%v]\n", status, resp.Header.Get(httpx.ContentType), resp.ContentLength)

	h.Set(core.XVersion, "v2")
	values.Add(httpx.ContentLocation, entryV2Path)
	resp, status = Get(nil, h, values)
	fmt.Printf("test: Get(nil,nil,values) -> [status:%v] [content-type:%v] [content-length:%v]\n", status, resp.Header.Get(httpx.ContentType), resp.ContentLength)

	//Output:
	//test: Get(nil,nil,nil) -> [status:Bad Request] [content-type:text/plain charset=utf-8] [content-length:32]
	//test: Get(nil,nil,values) -> [status:Not Found] [content-length:0]
	//test: Get(nil,nil,values) -> [status:Bad Request [invalid version: [v5]]] [content-type:text/plain charset=utf-8] [content-length:21]
	//test: Get(nil,nil,values) -> [status:OK] [content-type:application/json] [content-length:618]
	//test: Get(nil,nil,values) -> [status:OK] [content-type:application/json] [content-length:657]

}

func _Test_get(t *testing.T) {
	type args struct {
		ctx    context.Context
		h      http.Header
		values url.Values
	}
	type testCase[T entryConstraints] struct {
		name  string
		args  args
		want  []T
		want1 *core.Status
	}
	tests := []testCase[entryV1 /* TODO: Insert concrete types here */ ]{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, got1 := get[core.Output, entryV1](tt.args.ctx, tt.args.h, tt.args.values)
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("get() got = %v, want %v", got, tt.want)
			}
			if !reflect.DeepEqual(got1, tt.want1) {
				t.Errorf("get() got1 = %v, want %v", got1, tt.want1)
			}
		})
	}
}
