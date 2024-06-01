package resiliency

import (
	"fmt"
	"github.com/advanced-go/guidance/module"
	"github.com/advanced-go/stdlib/core"
	"github.com/advanced-go/stdlib/httpx"
	"net/http"
)

var (
	entryContent = httpx.NewListContent[Entry, httpx.Patch, struct{}](false, matchEntry2, nil, nil)
	entryRsc     = httpx.NewResource[Entry, httpx.Patch, struct{}](module.DocumentsResource, entryContent, nil)
	host, err    = httpx.NewHost(module.DocumentsAuthority, mapResource2, entryRsc.Do)
)

func init() {
	if err != nil {
		fmt.Printf("error: new resource %v", err)
	}
	//ctrl := controller.NewController("entry-resource", controller.NewPrimaryResource("", module.DocumentsAuthority, time.Second*2, "", host.Do), nil)
	//controller.RegisterController(ctrl)
}

func matchEntry2(req *http.Request, item *Entry) bool {
	filter := core.NewOrigin(req.URL.Query())
	if core.OriginMatch(item.Origin, filter) {
		return true
	}
	return false
}

func mapResource2(r *http.Request) string {
	return module.DocumentsResource

}
