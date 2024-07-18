package controller1

// Issues, questions, and decisions.
// Could there be a way to run in a test environment without cloud configuration:
// Rate Limiting would be not enabled
// Host would have to be configured
// How to not know to contact cloud for configuration
// Only thing needed to download is the startup hosts.

type ClientIngressRoute struct {
	Timeout int    `json:"timeout"`
	Host    string `json:"host"`
}

// ClientIngressAuthority -
// Optional - for logging
type ClientIngressAuthority struct {
	Authority string `json:"authority"` // provider/account/repository
	RouteName string `json:"route"`
}

type ClientEgressRoute struct {
	// Used by client
	Timeout int `json:"timeout"`

	// Route matching rules for routes.
	// For static routing, then need the full URL including Host name
	RouteName string
}

type ClientIngressLogging struct {
}

type ClientEgressLogging struct {
}
