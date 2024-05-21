package resiliency

import (
	"context"
	"github.com/advanced-go/stdlib/core"
	"net/http"
	"net/url"
	"reflect"
	"testing"
)

func TestGet(t *testing.T) {
	type args struct {
		ctx    context.Context
		h      http.Header
		values url.Values
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
			got, got1 := Get(tt.args.ctx, tt.args.h, tt.args.values)
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Get() got = %v, want %v", got, tt.want)
			}
			if !reflect.DeepEqual(got1, tt.want1) {
				t.Errorf("Get() got1 = %v, want %v", got1, tt.want1)
			}
		})
	}
}

func Test_get(t *testing.T) {
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
	tests := []testCase[entryV1 /* TODO: Insert concrete types here */]{
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
