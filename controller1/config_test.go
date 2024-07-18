package controller1

import (
	"fmt"
	json2 "github.com/advanced-go/stdlib/json"
)

var (
	case1 = Case{
		Desc: "client: [match] cloud: []",
		Client: Client{Match: Matcher{
			Path:      "/test/search?q=golang",
			Template:  "{path}",
			Authority: "https://www.google.com",
			Route:     "google-search",
		},
			Timeout: 0,
		},
		Cloud: Cloud{RateLimiting: false},
	}

	case2 = Case{
		Desc: "client: [match,timeout] cloud: [rate-limiting]",
		Client: Client{Match: Matcher{
			Path:      "/test/search?q=golang",
			Template:  "{path}",
			Authority: "https://www.google.com",
			Route:     "google-search",
		},
			Timeout: 3000,
		},
		Cloud: Cloud{RateLimiting: true},
	}

	case3 = Case{
		Desc: "client: [match,timeout] cloud: [rate-limiting,failover]",
		Client: Client{Match: Matcher{
			Path:      "/test/search?q=golang",
			Template:  "{path}",
			Authority: "https://www.google.com",
			Route:     "google-search",
		},
			Timeout: 3000,
		},
		Cloud: Cloud{
			RateLimiting: true,
			RegionT:      "us-central1",
			ZoneT:        "a",
			SubZoneT:     "",
			HostT:        "google.com",
		},
	}

	case4 = Case{
		Desc: "client: [match,timeout] cloud: [rate-limiting,failover,dynamic routing]",
		Client: Client{Match: Matcher{
			Path:      "/test/search?q=golang",
			Template:  "{path}",
			Authority: "https://www.google.com",
			Route:     "google-search",
		},
			Timeout: 3000,
		},
		Cloud: Cloud{
			RateLimiting:     true,
			RegionT:          "us-central1",
			ZoneT:            "a",
			SubZoneT:         "",
			HostT:            "google.com",
			Authority:        "github/advanced-go/observation",
			AuthorityVersion: "2.3.*",
		},
	}
)

func ExampleCase_1() {
	buf, status := json2.Marshal(&case1)
	fmt.Printf("test: Case_1() -> [status:%v] [%v]\n", status, string(buf))

	//Output:
	//fail
}

func ExampleCase_2() {
	buf, status := json2.Marshal(&case2)
	fmt.Printf("test: Case_2() -> [status:%v] [%v]\n", status, string(buf))

	//Output:
	//fail
}

func ExampleCase_3() {
	buf, status := json2.Marshal(&case3)
	fmt.Printf("test: Case_3() -> [status:%v] [%v]\n", status, string(buf))

	//Output:
	//fail
}

func ExampleCase_4() {
	buf, status := json2.Marshal(&case4)
	fmt.Printf("test: Case_4() -> [status:%v] [%v]\n", status, string(buf))

	//Output:
	//fail
}
