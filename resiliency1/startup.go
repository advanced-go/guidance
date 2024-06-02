package resiliency1

import (
	"context"
	"fmt"
	"github.com/advanced-go/guidance/module"
	"github.com/advanced-go/stdlib/controller"
	"github.com/advanced-go/stdlib/core"
	"github.com/advanced-go/stdlib/host"
	"github.com/advanced-go/stdlib/httpx"
	"github.com/advanced-go/stdlib/json"
	"github.com/advanced-go/stdlib/messaging"
	"net/http"
	"time"
)

const (
	entriesJson = "file:///c:/Users/markb/GitHub/guidance/resiliency1/documents-v1.json"
)

func init() {
	a, err1 := host.RegisterControlAgent(PkgPath, messageHandler)
	if err1 != nil {
		fmt.Printf("init(\"%v\") failure: [%v]\n", PkgPath, err1)
	}
	a.Run()
	initializeDocuments()
}

func messageHandler(msg *messaging.Message) {
	start := time.Now()
	switch msg.Event() {
	case messaging.StartupEvent:
		// Any processing for a Startup event would be here
		messaging.SendReply(msg, core.NewStatusDuration(http.StatusOK, time.Since(start)))
	case messaging.ShutdownEvent:
	case messaging.PingEvent:
		// Any processing for a Shutdown/Ping event would be here
		messaging.SendReply(msg, core.NewStatusDuration(http.StatusOK, time.Since(start)))
	}
}

var (
	docsContent = httpx.NewListContent[Entry, httpx.Patch, struct{}](false, matchEntry, nil, nil)
	docsRsc     = httpx.NewResource[Entry, httpx.Patch, struct{}](module.DocumentsResource, docsContent, nil)
	docs, err   = httpx.NewHost(module.DocumentsAuthority, mapResource, docsRsc.Do)
)

func initializeDocuments() {
	//defer controller.DisableLogging(true)()
	if err != nil {
		fmt.Printf("error: new resource %v", err)
	}
	entries, status := json.New[[]Entry](entriesJson, nil)
	if !status.OK() {
		fmt.Printf("initializeDocuments.New() -> [status:%v]\n", status)
		return
	}
	cfg, ok := module.GetRoute(module.DocumentsRouteName)
	if !ok {
		fmt.Printf("initializeDocuments.GetRoute() [ok:%v]\n", ok)
	}
	ctrl := controller.New(cfg, docs.Do)
	err = controller.RegisterController(ctrl)
	if err != nil {
		fmt.Printf("initializeDocuments.RegisterController() [err:%v]\n", err)
	}
	status = put[core.Output](context.Background(), nil, entries)
	if !status.OK() {
		fmt.Printf("initializeDocuments.put() [status:%v]\n", status)
	}
}

func matchEntry(req *http.Request, item *Entry) bool {
	filter := core.NewOrigin(req.URL.Query())
	target := core.Origin{Region: item.Region, Zone: item.Zone, SubZone: item.SubZone, Host: item.Host}
	if core.OriginMatch(target, filter) {
		return true
	}
	return false
}

func mapResource(r *http.Request) string {
	return module.DocumentsResource

}
