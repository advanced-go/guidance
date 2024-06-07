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
	DocumentsAuthority  = "github/advanced-go/documents"
	DocumentsResourceV1 = "v1/resiliency"

	DocumentsAuthorityV2 = "github/advanced-go/documents/2"
	DocumentsResourceV2  = "v2/resiliency"

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
