package http

import (
	"errors"
	"fmt"
	"github.com/advanced-go/guidance/module"
	"github.com/advanced-go/guidance/resiliency"
	"github.com/advanced-go/stdlib/controller"
	"github.com/advanced-go/stdlib/core"
	"github.com/advanced-go/stdlib/httpx"
	"net/http"
	"strings"
	"time"
)

// https://localhost:8081/github/advanced-go/guidance:v1/us-west/zone/sub-zone/app/route?q=golang

const (
	resiliencyPath = "resiliency"
)

var authorityResponse = httpx.NewAuthorityResponse(module.Authority)

// Controllers - egress controllers
func Controllers() []*controller.Controller {
	return []*controller.Controller{
		controller.NewController("google-search", controller.NewPrimaryResource("www.google.com", "", time.Second*2, "", nil), nil),
	}
}

// Exchange - HTTP exchange function
func Exchange(r *http.Request) (*http.Response, *core.Status) {
	_, path, status := httpx.ValidateRequestURL(r, module.Authority)
	if !status.OK() {
		return httpx.NewErrorResponse(status), status
	}
	switch strings.ToLower(path) {
	case resiliencyPath:
		return resiliencySwitch(r)
	case core.VersionPath:
		return httpx.NewVersionResponse(module.Version), core.StatusOK()
	case core.AuthorityPath:
		return authorityResponse, core.StatusOK()
	case core.HealthReadinessPath, core.HealthLivenessPath:
		return httpx.NewHealthResponseOK(), core.StatusOK()
	default:
		status = core.NewStatusError(http.StatusNotFound, errors.New(fmt.Sprintf("error invalid URI, resource not found: [%v]", path)))
		return httpx.NewErrorResponse(status), status
	}
}

func resiliencySwitch(r *http.Request) (*http.Response, *core.Status) {
	switch r.Method {
	case http.MethodGet:
		return resiliency.Get(r)
	case http.MethodPut:
		return resiliency.Post(r)
	default:
		status := core.NewStatusError(http.StatusBadRequest, errors.New(fmt.Sprintf("error invalid method: [%v]", r.Method)))
		return httpx.NewErrorResponse(status), status
	}
}
