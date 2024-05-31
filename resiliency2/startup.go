package resiliency

import (
	"fmt"
	"github.com/advanced-go/stdlib/controller"
	"github.com/advanced-go/stdlib/core"
	"github.com/advanced-go/stdlib/httpx"
	"net/http"
	"time"
)

var (
	entryContent = httpx.NewListContent[Entry, httpx.Patch, struct{}](false, matchEntry, nil, nil)
	entryRsc     = httpx.NewResource[Entry, httpx.Patch, struct{}](documentsResource, entryContent, nil)
	host, err    = httpx.NewHost(DocumentsAuthority, mapResource, entryRsc.Do)
)

func init() {
	if err != nil {
		fmt.Printf("error: new resource %v", err)
	}
	ctrl := controller.NewController("entry-resource", controller.NewPrimaryResource("", DocumentsAuthority, time.Second*2, "", host.Do), nil)
	controller.RegisterController(ctrl)
}

func matchEntry(req *http.Request, item *Entry) bool {
	filter := core.NewOrigin(req.URL.Query())
	if core.OriginMatch(item.Origin, filter) {
		return true
	}
	return false
}

func mapResource(r *http.Request) string {
	return documentsResource

}
