package resiliency

import (
	"fmt"
	"github.com/advanced-go/guidance/module"
	"github.com/advanced-go/stdlib/core"
	"github.com/advanced-go/stdlib/httpx"
	"net/http"
)

const (
	PkgPath        = "github/advanced-go/guidance/resiliency1"
	resiliencyName = "resiliency1"
)

type Entry struct {
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

type PutBodyConstraints interface {
	[]Entry | []byte | *http.Request
}

var (
	entryContent = httpx.NewListContent[Entry, httpx.Patch, struct{}](matchEntry, nil, nil)
	entryRsc     = httpx.NewResource2[Entry, httpx.Patch, struct{}](resiliencyName, entryContent, nil)
	host, err    = httpx.NewHost(module.DocumentsAuthority, mapResource, entryRsc.Do)
)

func init() {
	if err != nil {
		fmt.Printf("error: new resource %v", err)
	}
	//ctrl := controller.NewController("entry-resource", controller.NewPrimaryResource("", module.DocumentsAuthority, time.Second*2, "", host.Do), nil)
	//controller.RegisterController(ctrl)
}

func matchEntry(req *http.Request, item *Entry) bool {
	filter := core.NewOrigin(req.URL.Query())
	if core.OriginMatch(item.Origin, filter) {
		return true
	}
	return false
}

func mapResource(r *http.Request) string {
	return resiliencyName

}
