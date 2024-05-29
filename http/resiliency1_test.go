package http

import (
	"fmt"
	"github.com/advanced-go/guidance/module"
	resiliency1 "github.com/advanced-go/guidance/resiliency1"
	resiliency2 "github.com/advanced-go/guidance/resiliency2"
	"github.com/advanced-go/stdlib/core"
	"github.com/advanced-go/stdlib/httpx/httpxtest"
	"io"
	"net/http"
	"reflect"
	"strings"
	"testing"
)

const (
	documentsV1Name = module.Ver1 + "/" + module.DocumentsResource
	actionTest      = "test"
	actionEmpty     = "empty"
	actionInit      = "init"
	actionAdd       = "add"
)

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
	//if err != nil {
	//	fmt.Printf("error: new resource %v", err)
	//}
	//ctrl := controller.NewController("entry-resource", controller.NewPrimaryResource("", module.DocumentsAuthority, time.Second*2, "", host.Do), nil)
	//controller.RegisterController(ctrl)
}

func Test_resiliency1Exchange(t *testing.T) {
	basePath := "file://[cwd]/httptest/resiliency1/"

	type args struct {
		req  string
		resp string
	}
	workflow := []struct {
		action string
		name   string
		args   args
	}{
		{actionInit, "rsc-one", args{req: "document-init.json", resp: ""}},
		{actionAdd, "rsc-one", args{req: "document-add.json", resp: ""}},
		{actionEmpty, "rsc-one", args{req: "", resp: ""}},

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
			switch tt.action {
			case actionEmpty:

			}
			continue
		}
		failures, req, resp := httpxtest.ReadHttp(basePath, tt.args.req, tt.args.resp)
		if failures != nil {
			t.Errorf("ReadHttp() failures = %v", failures)
			continue
		}
		t.Run(tt.name, func(t *testing.T) {
			got, status := resiliencyExchange(req, nil)
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
			switch req.Header.Get(core.XVersion) {
			case module.Ver1, "":
				failures, gotT, wantT = httpxtest.Unmarshal[resiliency1.Entry](gotBuf, wantBuf)
			case module.Ver2:
				failures, gotT, wantT = httpxtest.Unmarshal[resiliency2.Entry](gotBuf, wantBuf)
			default:
			}
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
