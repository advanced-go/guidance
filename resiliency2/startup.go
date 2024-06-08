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
	content            = httpx.NewListContent[Entry, httpx.Patch, struct{}](false, matchEntry, nil, nil)
	resource           = httpx.NewResource[Entry, httpx.Patch, struct{}](module.DocumentsResourceV2, content, nil)
	authority, authErr = httpx.NewHost(module.DocumentsAuthorityV2, mapResource, resource.Do)
)

func init() {
	defer controller.DisableLogging(true)()
	if authErr != nil {
		fmt.Printf("error: new resource %v", authErr)
	}
	ctrl := controller.NewController("entry-resource", controller.NewPrimaryResource("", module.DocumentsAuthorityV2, time.Second*2, "", authority.Do), nil)
	err1 := controller.RegisterController(ctrl)
	if err1 != nil {
		fmt.Printf("initializeDocuments.RegisterController() [err:%v]\n", err1)
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
	return module.DocumentsResourceV2

}
