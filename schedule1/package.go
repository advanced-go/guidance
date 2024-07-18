package schedule1

import (
	"context"
	"github.com/advanced-go/stdlib/core"
)

const (
	PkgPath = "github/advanced-go/guidance/schedule1"
)

// GetGlobal - resource GET
func GetGlobal(ctx context.Context, origin core.Origin) (Entry, *core.Status) {
	return entryData[0], core.StatusOK()
}

/*
func GetHost(ctx context.Context, origin core.Origin) (Entry, *core.Status) {
	return entryData[0], core.StatusOK()
}
*/

func GetGroup(ctx context.Context, groupId string) (Entry, *core.Status) {
	return entryData[0], core.StatusOK()
}

func IsScheduled(scheduleId string) bool {
	return true
}

func IsIngressControllerScheduled() bool {
	return IsScheduled("ingress-controller")
}

func IsEgressControllerScheduled() bool {
	return IsScheduled("egress-controller")
}

func IsDependencyUpdateScheduled(scheduleId string) bool {
	return IsScheduled("dependency-update")
}
