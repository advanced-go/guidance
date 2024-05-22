package http

import (
	"github.com/advanced-go/guidance/module"
	"github.com/advanced-go/stdlib/core"
	"github.com/advanced-go/stdlib/httpx/httpxtest"
	"reflect"
	"testing"
)

type resiliencyEntryV1 struct {
	Origin    core.Origin
	Status    string `json:"status"`
	CreatedTS string `json:"created-ts"`
	UpdatedTS string `json:"updated-ts"`

	// Timeout
	Timeout string `json:"timeout"`

	// Rate Limiting
	RateLimit string `json:"rate-limit"`
	RateBurst string `json:"rate-burst"`
}

type resiliencyEntryV2 struct {
	Origin    core.Origin
	Version   string `json:"version"`
	Status    string `json:"status"`
	CreatedTS string `json:"created-ts"`
	UpdatedTS string `json:"updated-ts"`

	// Timeout
	Timeout string `json:"timeout"`

	// Rate Limiting
	RateLimit string `json:"rate-limit"`
	RateBurst string `json:"rate-burst"`
}

func Test_resiliencyExchange(t *testing.T) {
	basePath := "file://[cwd]/httptest/"

	type args struct {
		req  string
		resp string
	}
	tests := []struct {
		name string
		args args
		//want  *http.Response
		//want1 *core.Status
	}{
		{"get-v1", args{req: "get-req-v1.txt", resp: "get-resp-v1.txt"}},
	}
	for _, tt := range tests {
		failures, req, resp := httpxtest.ReadHttp(basePath, tt.args.req, tt.args.resp)
		if failures != nil {
			t.Errorf("ReadHttp() failures = %v", failures)
			continue
		}
		t.Run(tt.name, func(t *testing.T) {
			got, _ := resiliencyExchange(req)
			// test status code
			if got.StatusCode != resp.StatusCode {
				t.Errorf("StatusCode got = %v, want %v", got.StatusCode, resp.StatusCode)
			} else {
				// test headers if needed - test2.Headers(w.Result(),resp,names... string) (failures []Args)

				// test content type and length
				var gotBuf []byte
				var wantBuf []byte
				failures, gotBuf, wantBuf = httpxtest.Content(got, resp)
				if failures != nil {
					httpxtest.Errorf(t, failures)
				} else {
					// test content
					switch req.Header.Get(core.XVersion) {
					case module.Ver1, "":
						failures, gotT, wantT = httpxtest.Unmarshal[resiliencyEntryV1](gotBuf, wantBuf)
					case module.Ver2:
						failures, gotT, wantT = httpxtest.Unmarshal[resiliencyEntryV2](gotBuf, wantBuf)
					default:
					}
					if failures != nil {
						httpxtest.Errorf(t, failures)
					} else {
						if !reflect.DeepEqual(gotT, wantT) {
							t.Errorf("DeepEqual() got = %v, want %v", gotT, wantT)
						}
					}
				}
			}
		})
	}
}
