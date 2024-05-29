package http

import (
	"github.com/advanced-go/stdlib/core"
	"github.com/advanced-go/stdlib/uri"
	"net/http"
)

func resiliencyExchangeV2(r *http.Request, p *uri.Parsed) (*http.Response, *core.Status) {
	return resiliencyExchange(r, p)
}
