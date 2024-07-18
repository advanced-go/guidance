package controller1

import (
	"fmt"
	json2 "github.com/advanced-go/stdlib/json"
)

var (
	set = Changeset{
		Authority: AuthorityChangeset{
			Insert: []AuthorityChange{
				{
					Name:    "github/advanced-go/observation",
					Version: "2.3.*",
					Role:    "primary",
				},
			},
		},
		Egress: EgressChangeset{
			Insert: []EgressChange{
				{
					RouteName:    "google-search",
					RateLimiting: false,
					RegionT:      "us-central1",
					ZoneT:        "a",
					SubZoneT:     "",
					HostT:        "google.com",
					Authority:    "github/advanced-go/observation",
					Version:      "2.3.*",
				},
			},
		},
	}
	/*
		set = Changset{
			Authority: AuthorityChangeset{
				Insert{
				Name: "github/advanced-go/observation",
				Version: "2.3.2",
				Role: "primary",
				},
			},
			Ingress:   []IngressChangset{
				{
					RouteName: "host",
					RateLimiting: true,
				},
			},
			Egress:    []EgressChangset {
				{
					RouteName: "google-search",
					RateLimiting: true,
				}
			},
		}

	*/
)

func ExampleChangeset_Nil() {
	buf, status := json2.Marshal(&set)
	fmt.Printf("test: Case_4() -> [status:%v] [%v]\n", status, string(buf))

	//Output:
	//fail

}
