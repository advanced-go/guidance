package module

import (
	"github.com/advanced-go/stdlib/controller"
	"github.com/advanced-go/stdlib/core"
	"time"
)

const (
	Authority = "github/advanced-go/guidance"
	RouteName = "guidance"
	Version   = "1.1.1"
	Ver1      = "v1"
	Ver2      = "v2"

	ResiliencyResource = "resiliency"
)

// Configuration keys used on startup for map values
const (
	PackageNameUserKey     = "user"    // type:text, empty:false
	PackageNamePasswordKey = "pswd"    // type:text, empty:false
	PackageNameRetriesKey  = "retries" // type:int, default:-1, range:0-100
)

// Upstream authorities/resources
const (
	DocumentsAuthority = "github/advanced-go/documents"
	DocumentsResource  = "resiliency"
	DocumentsPath      = "/github/advanced-go/documents:%sresiliency"
	DocumentsV1        = "v1"
	DocumentsV2        = "v2"
	DocumentsRouteName = "documents"
)

// Routes - upstream egress traffic route configuration
var (
	Routes = []controller.Config{
		{DocumentsRouteName, "localhost:8081", DocumentsAuthority, core.HealthLivenessPath, time.Second * 2},
	}
)

// GetRoute - get the route configuration
func GetRoute(routeName string) (controller.Config, bool) {
	return controller.GetRoute(routeName, Routes)
}
