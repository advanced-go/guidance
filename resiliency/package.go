package resiliency

import (
	"github.com/advanced-go/stdlib/core"
	"net/http"
)

const (
	PkgPath = "github/advanced-go/guidance/resiliency"
)

type EntryV1 struct {
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

type EntryV2 struct {
	Origin    core.Origin
	Version   string `json:"version"`
	Status    string `json:"status"`
	CreatedTS string `json:"created-ts"`
	UpdatedTS string `json:"updated-ts"`

	// Timeout
	Timeout string `json:"timeout"`

	// Rate Limiting
	RateLimit string `json:"rate-limit"`
	RateBurst string `json:"rate-burst"`
}

type EntryConstraints interface {
	EntryV1 | EntryV2
}

type PutBodyConstraints interface {
	[]EntryV1 | []EntryV2 | []byte | *http.Request
}
