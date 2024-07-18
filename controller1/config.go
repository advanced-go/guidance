package controller1

type Matcher struct {
	Path      string `json:"path"`
	Template  string `json:"template"`
	Authority string `json:"authority"`
	Route     string `json:"route"`
}

type Client struct {
	Match   Matcher `json:"match"`
	Timeout int     `json:"timeout"`
}

type Cloud struct {
	RateLimiting     bool   `json:"rate-limiting"`
	RegionT          string `json:"region-t"`
	ZoneT            string `json:"zone-t"`
	SubZoneT         string `json:"sub-zone-t"`
	HostT            string `json:"host-t"`
	Authority        string `json:"authority"` // github/advanced-go/observation: provider/account/repository
	AuthorityVersion string `json:"authority-version"`
}

type Config struct {
	Client Client
	Cloud  Cloud
}

type Case struct {
	Desc   string `json:"desc"`
	Client Client `json:"client"`
	Cloud  Cloud  `json:"cloud"`
}
