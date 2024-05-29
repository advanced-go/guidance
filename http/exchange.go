package http

import (
	"errors"
	"fmt"
	"github.com/advanced-go/guidance/module"
	//resiliency2 "github.com/advanced-go/guidance/resiliency2"
	"github.com/advanced-go/stdlib/controller"
	"github.com/advanced-go/stdlib/core"
	"github.com/advanced-go/stdlib/httpx"
	"net/http"
	"strings"
	"time"
)

// https://localhost:8081/github/advanced-go/guidance:v1/us-west/zone/sub-zone/app/route?q=golang

var authorityResponse = httpx.NewAuthorityResponse(module.Authority)

// Controllers - egress controllers
func Controllers() []*controller.Controller {
	return []*controller.Controller{
		controller.NewController("google-search", controller.NewPrimaryResource("www.google.com", "", time.Second*2, "", nil), nil),
	}
}

// Exchange - HTTP exchange
func Exchange(r *http.Request) (*http.Response, *core.Status) {
	if r == nil {
		return httpx.NewResponseWithStatus(core.StatusBadRequest(), nil)
	}
	p, status := httpx.ValidateURL(r.URL, module.Authority)
	if !status.OK() {
		return httpx.NewResponseWithStatus(status, status.Err)
	}
	core.AddRequestId(r.Header)
	switch strings.ToLower(p.Resource) {
	case module.ResiliencyResource:
		return resiliencyExchange(r, p)
	case core.VersionPath:
		return httpx.NewVersionResponse(module.Version), core.StatusOK()
	case core.AuthorityPath:
		return authorityResponse, core.StatusOK()
	case core.HealthReadinessPath, core.HealthLivenessPath:
		return httpx.NewHealthResponseOK(), core.StatusOK()
	default:
		status = core.NewStatusError(http.StatusNotFound, errors.New(fmt.Sprintf("error invalid URI, resource not found: [%v]", p.Path)))
		return httpx.NewResponseWithStatus(status, status.Err)
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
