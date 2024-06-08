package http

import (
	"fmt"
	"github.com/advanced-go/guidance/module"
	resiliency1 "github.com/advanced-go/guidance/resiliency1"
	"github.com/advanced-go/stdlib/controller"
	"github.com/advanced-go/stdlib/core"
	"github.com/advanced-go/stdlib/httpx"
	"github.com/advanced-go/stdlib/httpx/httpxtest"
	"io"
	"net/http"
	"reflect"
	"strings"
	"testing"
)

const (
	documentsV1Name = module.Ver1 + "/" + "resiliency"
	actionTest      = "test"
	//actionEmpty     = "empty"
	actionInit  = "init"
	actionAdd   = "add"
	entriesJson = "file:///c:/Users/markb/GitHub/guidance/resiliency1/resiliency1test/documents-v1.json"
)

var (
	content            = httpx.NewListContent[resiliency1.Entry, httpx.Patch, struct{}](false, matchEntry, nil, nil)
	resource           = httpx.NewResource[resiliency1.Entry, httpx.Patch, struct{}](module.DocumentsResourceV1, content, nil)
	authority, hostErr = httpx.NewHost(module.DocumentsAuthority, mapResource, resource.Do)
)

func matchEntry(req *http.Request, item *resiliency1.Entry) bool {
	filter := core.NewOrigin(req.URL.Query())
	target := core.Origin{Region: item.Region, Zone: item.Zone, SubZone: item.SubZone, Host: item.Host}
	if core.OriginMatch(target, filter) {
		return true
	}
	return false
}

func mapResource(r *http.Request) string {
	return module.DocumentsResourceV1

}

func matchOrigin(item *core.Origin, req *http.Request) bool {
	filter := core.NewOrigin(req.URL.Query())
	if core.OriginMatch(*item, filter) {
		return true
	}
	return false
}

func mapResiliency(r *http.Request) string {
	if strings.Contains(r.URL.Path, documentsV1Name) {
		return documentsV1Name
	}
	return documentsV1Name
}

var (
// rsc       = httpx.NewResource2[core.Origin, struct{}, struct{}](documentsV1Name, matchOrigin, nil, nil, nil)
// host, err = httpx.NewHost(module.DocumentsAuthority, mapResiliency, rsc.Do)
)

func init() {
	//initializeDocuments
}

func initializeDocuments() {
	defer controller.DisableLogging(true)()
	if hostErr != nil {
		fmt.Printf("error: new resource %v", hostErr)
	}
	//entries, status := json.New[[]resiliency1.Entry](entriesJson, nil)
	//if !status.OK() {
	//	fmt.Printf("initializeDocuments.New() -> [status:%v]\n", status)
	//	return
	//}
	cfg, ok := module.GetRoute(module.DocumentsRouteName)
	if !ok {
		fmt.Printf("initializeDocuments.GetRoute() [ok:%v]\n", ok)
	}
	ctrl := controller.New(cfg, authority.Do)
	err0 := controller.RegisterController(ctrl)
	if err0 != nil {
		fmt.Printf("initializeDocuments.RegisterController() [err:%v]\n", err0)
	}
	//_, status = put[core.Output](context.Background(), nil, entries)
	///if !status.OK() {
	//	fmt.Printf("initializeDocuments.put() [status:%v]\n", status)
	//}
}

func Test_resiliencyExchangeV1(t *testing.T) {
	defer controller.DisableLogging(true)()
	basePath := "file://[cwd]/httptest/resiliency1/"

	type args struct {
		req  string
		resp string
	}
	workflow := []struct {
		suite  string
		test   string
		action string
		args   args
	}{
		{"empty-all", "first", "test", args{req: "get-empty-req.txt", resp: "get-empty-resp.txt"}},
		{"empty-all", "next", "test", args{req: "get-all-req.txt", resp: "get-all-resp.txt"}},

		/*
			{actionTest,"get-empty", args{req: "get-empty-req.txt", resp: "get-empty-resp.txt"}},
			{actionTest,"get-entry-empty", args{req: "get-entry-empty-req.txt", resp: "get-empty-resp.txt"}},
			{actionTest,"get-v1", args{req: "get-req-v1.txt", resp: "get-resp-v1.txt"}},
			{actionTest,"get-v2", args{req: "get-req-v2.txt", resp: "get-resp-v2.txt"}},
			{actionTest,"get-query-v2", args{req: "get-query-req-v2.txt", resp: "get-query-resp-v2.txt"}},

		*/
	}
	for _, tt := range workflow {
		if tt.action != actionTest {
			fmt.Printf("action: %v\n", tt.action)
			continue
		}
		failures, req, resp := httpxtest.ReadHttp(basePath, tt.args.req, tt.args.resp)
		if failures != nil {
			t.Errorf("ReadHttp() failures = %v", failures)
			continue
		}
		t.Run(tt.suite+"/"+tt.test, func(t *testing.T) {
			got, status := resiliencyExchange[core.Output](req, nil)
			// test status code
			if got.StatusCode != resp.StatusCode {
				var buf []byte
				if got.Body != nil {
					buf, _ = io.ReadAll(got.Body)
				}
				t.Errorf("StatusCode got = %v, want = %v content = %v", got.StatusCode, resp.StatusCode, string(buf))
				return
			}
			if !status.OK() {
				return
			}
			// test headers if needed - test2.Headers(w.Result(),resp,names... string) (failures []Args)

			// test content type, body IO, and optionally, content length
			var gotBuf []byte
			var wantBuf []byte
			failures, gotBuf, wantBuf = httpxtest.Content(got, resp)
			if failures != nil {
				httpxtest.Errorf(t, failures)
				return
			}

			// test content
			var gotT any
			var wantT any
			failures, gotT, wantT = httpxtest.Unmarshal[resiliency1.Entry](gotBuf, wantBuf)
			if failures != nil {
				httpxtest.Errorf(t, failures)
				return
			}
			if !reflect.DeepEqual(gotT, wantT) {
				t.Errorf("DeepEqual() got = %v, want %v", gotT, wantT)
			}
		})
	}
}
