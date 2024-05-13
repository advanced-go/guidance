package http

import (
	"github.com/advanced-go/guidance/types"
	"github.com/advanced-go/stdlib/core"
	"net/http"
)

const (
	ModulePath = "github/advanced-go/guidance"
	PkgPath    = ModulePath + "/http"
)

// https://localhost:8081/github/advanced-go/guidance:v1/search?q=golang

func Exchange(r *http.Request) (*http.Response, *core.Status) {
	var e *types.EntryV1
	if e == nil {
	}
	return nil, nil
}
