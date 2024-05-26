package module

import (
	"fmt"
	"net/url"
)

const (
	DocumentsAuthority = "github/advanced-go/documents"
	DocumentsResource  = "resiliency"

	ResiliencyResource      = "resiliency1"
	completeFormat          = "/%v:%v/%v?%v"
	noQueryFormat           = "/%v:%v/%v"
	noVersionCompleteFormat = "/%v:%v?%v"
	noVersionQueryFormat    = "/%v:%v"
)

func BuildPath(authority, version, resource string, values url.Values) string {
	return fmt.Sprintf("%v%v", authority, ":"+BuildResourcePath(version, resource, values))
}

func BuildResourcePath(version, resource string, values url.Values) string {
	return fmt.Sprintf("%v%v", formatPath(version, resource)[1:], formatValues(values))
}

func BuildDocumentsPath(version string, values url.Values) string {
	return BuildPath(DocumentsAuthority, version, DocumentsResource, values)
}

func BuildDocumentsResource(version string, values url.Values) string {
	return BuildResourcePath(version, DocumentsResource, values)
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
