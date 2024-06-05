package http

import (
	"errors"
	"fmt"
	"github.com/advanced-go/guidance/module"

	"github.com/advanced-go/stdlib/core"
	"github.com/advanced-go/stdlib/httpx"
	"net/http"
)

// https://localhost:8081/github/advanced-go/guidance:v1/us-west/zone/sub-zone/app/route?q=golang

var authorityResponse = httpx.NewAuthorityResponse(module.Authority)

// Exchange - HTTP exchange
func Exchange(r *http.Request) (*http.Response, *core.Status) {
	h2 := make(http.Header)
	h2.Add(httpx.ContentType, httpx.ContentTypeText)

	if r == nil {
		return httpx.NewResponse[core.Log](http.StatusBadRequest, h2, nil)
	}
	p, status := httpx.ValidateURL(r.URL, module.Authority)
	if !status.OK() {
		return httpx.NewResponse[core.Log](status.HttpCode(), h2, status.Err)
	}
	core.AddRequestId(r.Header)
	r.Header.Set(core.XAuthority, module.Authority)
	switch p.Resource {
	case module.ResiliencyResource:
		return resiliencyExchange[core.Log](r, p)
	case core.VersionPath:
		return httpx.NewVersionResponse(module.Version), core.StatusOK()
	case core.AuthorityPath:
		return authorityResponse, core.StatusOK()
	case core.HealthReadinessPath, core.HealthLivenessPath:
		return httpx.NewHealthResponseOK(), core.StatusOK()
	default:
		status = core.NewStatusError(http.StatusNotFound, errors.New(fmt.Sprintf("error invalid URI, resource not found: [%v]", p.Path)))
		return httpx.NewResponse[core.Log](status.HttpCode(), h2, status.Err)
	}
}

/*
func resiliencyMux(r *http.Request, p *uri.Parsed) (*http.Response, *core.Status) {
	switch p.Version {
	case module.Ver1, "":
		return resiliencyExchangeV1(r, p)
	case module.Ver2:
		return resiliencyExchangeV2(r, p)
	default:
		status := core.NewStatusError(http.StatusBadRequest, errors.New(fmt.Sprintf("invalid version: [%v]", r.Header.Get(core.XVersion))))
		return httpx.NewResponseWithStatus(status, status.Err)
	}

}


*/
