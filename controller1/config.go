package controller1

type Matcher struct {
	Path      string `json:"path"`
	Template  string `json:"template"`
	Authority string `json:"authority"`
	RouteName string `json:"route-name"`
}

type Client struct {
	Match   Matcher `json:"match"`
	Timeout int     `json:"timeout"`
}

type CloudIngress struct {
	RedirectAuthority string `json:"redirect-authority"` // github/advanced-go/observation: provider/account/repository
}

type CloudEgress struct {
	RouteName         string `json:"route-name"`
	RedirectAuthority string `json:"redirect-authority"` // github/advanced-go/observation: provider/account/repository
	FailoverScope     string `json:"failover-scope"`
	FailoverThreshold int    `json:"failover-threshold"`
}

type Config struct {
	Client  Client
	Ingress CloudIngress
	Egress  CloudEgress
}

type Case struct {
	Desc    string       `json:"desc"`
	Client  Client       `json:"client"`
	Ingress CloudIngress `json:"ingress"`
	Egress  CloudEgress  `json:"egress"`
}
