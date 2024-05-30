package module

import (
	"github.com/advanced-go/stdlib/uri"
	"net/url"
)

const (
	DocumentsAuthority = "github/advanced-go/documents"
	DocumentsResource  = "resiliency"

	ResiliencyResource = "resiliency"
	//ResiliencyV2Resource    = "resiliency2"
	completeFormat          = "/%v:%v/%v?%v"
	noQueryFormat           = "/%v:%v/%v"
	noVersionCompleteFormat = "/%v:%v?%v"
	noVersionQueryFormat    = "/%v:%v"
)

/*
func BuildPath(authority, version, resource string, values url.Values) string {
	return fmt.Sprintf("%v%v", authority, ":"+BuildResourcePath(version, resource, values))
}

func BuildResourcePath(version, resource string, values url.Values) string {
	return fmt.Sprintf("%v%v", formatPath(version, resource)[1:], formatValues(values))
}


*/

func BuildDocumentsPath(version string, values url.Values) string {
	return uri.BuildURLWithAuthority("", DocumentsAuthority, version, DocumentsResource, values)
}

/*
func BuildDocumentsResource(version string, values url.Values) string {
	return fmt.Sprintf("%v%v", formatPath(version, DocumentsResource)[1:], formatValues(values))

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


*/
