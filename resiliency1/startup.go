package resiliency

import (
	"context"
	"fmt"
	"github.com/advanced-go/guidance/module"
	"github.com/advanced-go/stdlib/controller"
	"github.com/advanced-go/stdlib/core"
	"github.com/advanced-go/stdlib/httpx"
	"net/http"
)

var (
	docsContent = httpx.NewListContent[Entry, httpx.Patch, struct{}](false, matchEntry, nil, nil)
	docsRsc     = httpx.NewResource[Entry, httpx.Patch, struct{}](module.ResiliencyResource, docsContent, nil)
	docs, err   = httpx.NewHost(module.DocumentsAuthority, mapResource, docsRsc.Do)
)

var (
	testEntry = []Entry{
		{Origin: core.Origin{Region: "region1", Zone: "Zone1", Host: "www.host1.com"}, Status: "active", Timeout: "100ms", RateLimit: "125", RateBurst: "25"},
		{Origin: core.Origin{Region: "region1", Zone: "Zone2", Host: "www.host2.com"}, Status: "inactive", Timeout: "250ms", RateLimit: "100", RateBurst: "10"},
		{Origin: core.Origin{Region: "region2", Zone: "Zone1", Host: "www.google.com"}, Status: "removed", Timeout: "3s", RateLimit: "50", RateBurst: "5"},
	}
)

func init() {
	if err != nil {
		fmt.Printf("error: new resource %v", err)
	}
	ctrl := controller.NewExchangeController("documents", docs.Do)
	controller.RegisterController(ctrl)
	status := put[core.Output](context.Background(), nil, testEntry)
	if !status.OK() {
		fmt.Printf("resiliency1 startup error: %v\n", status)
	}

}

func matchEntry(req *http.Request, item *Entry) bool {
	filter := core.NewOrigin(req.URL.Query())
	if core.OriginMatch(item.Origin, filter) {
		return true
	}
	return false
}

func mapResource(r *http.Request) string {
	return module.ResiliencyResource

}
