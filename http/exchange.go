package http

import (
	"errors"
	"fmt"
	"github.com/advanced-go/guidance/module"
	"github.com/advanced-go/stdlib/core"
	"github.com/advanced-go/stdlib/httpx"
	"net/http"
	"strings"
)

// https://localhost:8081/github/advanced-go/guidance:v1/search?q=golang

var (
	versionResponse = httpx.NewResponse(core.StatusOK(), core.VersionContent(module.Version))
)

func Exchange(r *http.Request) (*http.Response, *core.Status) {
	_, path, status0 := httpx.ValidateRequestURL(r, module.Authority)
	if !status0.OK() {
		return httpx.NewErrorResponse(status0), status0
	}
	switch strings.ToLower(path) {
	case core.VersionPath:
		return versionResponse, core.StatusOK()
	case core.HealthReadinessPath, core.HealthLivenessPath:
		return httpx.HealthResponseOK, core.StatusOK()
	default:
		status := core.NewStatusError(http.StatusNotFound, errors.New(fmt.Sprintf("error invalid URI, resource not found: [%v]", path)))
		return httpx.NewErrorResponse(status), status
	}
}
