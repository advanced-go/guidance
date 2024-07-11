package controller1

import (
	"context"
	"github.com/advanced-go/stdlib/core"
	"net/http"
)

const (
	PkgPath = "github/advanced-go/guidance/controller1"
)

// Get - resource GET
func Get(ctx context.Context, h http.Header, origin core.Origin) ([]Entry, *core.Status) {
	return []Entry{lastEntry()}, core.StatusOK()
}
