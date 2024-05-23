package resiliency

import (
	"github.com/advanced-go/stdlib/core"
	"net/http"
)

const (
	PkgPath = "github/advanced-go/guidance/resiliency"
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
