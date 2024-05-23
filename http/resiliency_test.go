package http

import (
	"github.com/advanced-go/guidance/module"
	"github.com/advanced-go/stdlib/core"
	"github.com/advanced-go/stdlib/httpx/httpxtest"
	"io"
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
	basePath := "file://[cwd]/httptest/resiliency/"

	type args struct {
		req  string
		resp string
	}
	tests := []struct {
		name string
		args args
	}{
		{"get-empty", args{req: "get-empty-req.txt", resp: "get-empty-resp.txt"}},
		{"get-entry-empty", args{req: "get-entry-empty-req.txt", resp: "get-empty-resp.txt"}},
		{"get-v1", args{req: "get-req-v1.txt", resp: "get-resp-v1.txt"}},
		{"get-v2", args{req: "get-req-v2.txt", resp: "get-resp-v2.txt"}},
		{"get-query-v2", args{req: "get-query-req-v2.txt", resp: "get-query-resp-v2.txt"}},
	}
	for _, tt := range tests {
		failures, req, resp := httpxtest.ReadHttp(basePath, tt.args.req, tt.args.resp)
		if failures != nil {
			t.Errorf("ReadHttp() failures = %v", failures)
			continue
		}
		t.Run(tt.name, func(t *testing.T) {
			got, status := resiliencyExchange(req)
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
				failures, gotT, wantT = httpxtest.Unmarshal[resiliencyEntryV1](gotBuf, wantBuf)
			case module.Ver2:
				failures, gotT, wantT = httpxtest.Unmarshal[resiliencyEntryV2](gotBuf, wantBuf)
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
