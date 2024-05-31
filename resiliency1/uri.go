package resiliency

import (
	"fmt"
	"net/url"
	"strings"
)

const (
	documentsAuthority = "github/advanced-go/documents"
	//DocumentsResource  = "resiliency"

	documentsPath = "/github/advanced-go/documents:%sresiliency"
	documentsV1   = "v1"
	Localhost     = "localhost"
	HttpsScheme   = "https"
	HttpScheme    = "http"
)

func Expansion(host, path, version string, values url.Values) string {
	newUrl := strings.Builder{}
	if host != "" {
		scheme := HttpsScheme
		if strings.Contains(host, Localhost) {
			scheme = HttpScheme
		}
		newUrl.WriteString(scheme)
		newUrl.WriteString("://")
		newUrl.WriteString(host)
	}
	newUrl.WriteString(fmt.Sprintf(path, formatVersion(version)))
	newUrl.WriteString(formatValues(values))
	return newUrl.String()
}

func formatValues(values url.Values) string {
	if values == nil {
		return ""
	}
	return "?" + values.Encode()
}

func formatVersion(version string) string {
	if version == "" {
		return ""
	}
	return version + "/"
}
