package http

import (
	"github.com/advanced-go/stdlib/core"
	"github.com/advanced-go/stdlib/uri"
	"net/http"
	"reflect"
	"testing"
)

func Test_resiliencyExchangeV2(t *testing.T) {
	type args struct {
		r *http.Request
		p *uri.Parsed
	}
	tests := []struct {
		name  string
		args  args
		want  *http.Response
		want1 *core.Status
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, got1 := resiliencyExchange[core.Output](tt.args.r, tt.args.p)
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("resiliencyExchangeV2() got = %v, want %v", got, tt.want)
			}
			if !reflect.DeepEqual(got1, tt.want1) {
				t.Errorf("resiliencyExchangeV2() got1 = %v, want %v", got1, tt.want1)
			}
		})
	}
}
