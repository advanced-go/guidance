package http

import (
	"github.com/advanced-go/stdlib/core"
	"net/http"
)

func ValidateRequest(modulePath string, r *http.Request) (string, *core.Status) {
	//if r == nil {
	//	return "", core.NewStatusError(http.StatusBadRequest, errors.New("error: request is nil"))
	//}
	//nid, nss, ok := uri.UprootUrn(r.URL.String())
	//if !ok {
	//	return "", core.NewStatusError(http.StatusBadRequest, errors.New("error: request is nil"))
	//}
	return "", nil
}
