package http

import (
	"errors"
	"fmt"
	"github.com/advanced-go/guidance/module"
	"github.com/advanced-go/stdlib/controller"
	"github.com/advanced-go/stdlib/core"
	"github.com/advanced-go/stdlib/httpx"
	"net/http"
	"strings"
	"time"
)

// https://localhost:8081/github/advanced-go/guidance:v1/search?q=golang

const (
	googleProvider = "google"
	yahooProvider  = "yahoo"
)

var (
	versionResponse = httpx.NewResponse(core.StatusOK(), core.VersionContent(module.Version))
)

// Controllers - authority controllers
func Controllers() []*controller.Controller {
	return []*controller.Controller{
		controller.NewController("google-search", controller.NewPrimaryResource("www.google.com", time.Second*2, "", nil), nil),
	}
}

// Exchange - HTTP exchange function
func Exchange(r *http.Request) (*http.Response, *core.Status) {
	_, path, status := httpx.ValidateRequestURL(r, module.Authority)
	if !status.OK() {
		return httpx.NewErrorResponse(status), status
	}
	switch strings.ToLower(path) {
	case core.VersionPath:
		return versionResponse, core.StatusOK()
	case core.HealthReadinessPath, core.HealthLivenessPath:
		return httpx.HealthResponseOK, core.StatusOK()
	default:
		status = core.NewStatusError(http.StatusNotFound, errors.New(fmt.Sprintf("error invalid URI, resource not found: [%v]", path)))
		return httpx.NewErrorResponse(status), status
	}
}
