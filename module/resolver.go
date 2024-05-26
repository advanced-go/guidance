package module

import (
	"fmt"
	"net/url"
)

const (
	DocumentsAuthority      = "github/advanced-go/documents"
	ResiliencyResource      = "resiliency"
	completeFormat          = "/%v:%v/%v?%v"
	noQueryFormat           = "/%v:%v/%v"
	noVersionCompleteFormat = "/%v:%v?%v"
	noVersionQueryFormat    = "/%v:%v"
)

func BuildDocumentsPath(version string, values url.Values) string {
	return fmt.Sprintf("%v%v%v", DocumentsAuthority, formatPath(version, ResiliencyResource), formatValues(values))
}

func formatValues(values url.Values) string {
	if values == nil {
		return ""
	}
	return "?" + values.Encode()
}

func formatPath(version, path string) string {
	if version == "" {
		return ":" + path
	}
	return ":" + version + "/" + path
}
