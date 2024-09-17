package host1

import (
	"context"
	"github.com/advanced-go/stdlib/core"
)

const (
	PkgPath = "github/advanced-go/guidance/host1"
)

// CDC

// GetHostEntries - retrieve existing HostEntry
func GetHostEntries(ctx context.Context, origin core.Origin) ([]Entry, LastCDCId, *core.Status) {
	last := LastCDCId{
		Entry:    0,
		Redirect: 0,
		Egress:   0,
	}
	return []Entry{}, last, core.StatusOK()
}

// GetNewHostEntries - retrieve new HostEntry
func GetNewHostEntries(ctx context.Context, origin core.Origin, lastId int) ([]Entry, *core.Status) {
	return []Entry{}, core.StatusOK()
}
