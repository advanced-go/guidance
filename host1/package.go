package host1

import (
	"context"
	"github.com/advanced-go/stdlib/core"
)

const (
	PkgPath = "github/advanced-go/guidance/host1"
)

// CDC

// QueryIngressHosts - on startup, retrieve existing hosts and redirect status
func QueryIngressHosts(ctx context.Context, origin core.Origin) ([]EntryQuery, LastCDCId, *core.Status) {
	last := LastCDCId{
		Entry:    0,
		Redirect: 0,
		Egress:   0,
	}
	return []EntryQuery{}, last, core.StatusOK()
}

// QueryNewIngressHosts - retrieve new host entries
func QueryNewIngressHosts(ctx context.Context, origin core.Origin, lastId int) ([]Entry, *core.Status) {
	return []Entry{}, core.StatusOK()
}

// QueryEgressHosts - on startup, retrieve existing hosts and redirect status
func QueryEgressHosts(ctx context.Context, origin core.Origin) ([]EntryQuery, LastCDCId, *core.Status) {
	last := LastCDCId{
		Entry:    0,
		Redirect: 0,
		Egress:   0,
	}
	return []EntryQuery{}, last, core.StatusOK()
}

// QueryNewEgressHosts - retrieve new host entries
func QueryNewEgressHosts(ctx context.Context, origin core.Origin, lastId int) ([]Entry, *core.Status) {
	return []Entry{}, core.StatusOK()
}
