package resiliency2

import (
	"fmt"
	"github.com/advanced-go/guidance/module"
	"github.com/advanced-go/stdlib/controller"
	"github.com/advanced-go/stdlib/core"
	"github.com/advanced-go/stdlib/httpx"
	"net/http"
	"time"
)

var (
	entryContent = httpx.NewListContent[Entry, httpx.Patch, struct{}](false, matchEntry, nil, nil)
	entryRsc     = httpx.NewResource[Entry, httpx.Patch, struct{}](module.DocumentsResource, entryContent, nil)
	host, err    = httpx.NewHost(module.DocumentsAuthorityV2, mapResource, entryRsc.Do)
)

func init() {
	if err != nil {
		fmt.Printf("error: new resource %v", err)
	}
	ctrl := controller.NewController("entry-resource", controller.NewPrimaryResource("", module.DocumentsAuthorityV2, time.Second*2, "", host.Do), nil)
	err = controller.RegisterController(ctrl)
	if err != nil {
		fmt.Printf("initializeDocuments.RegisterController() [err:%v]\n", err)
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
	return module.DocumentsResource

}
