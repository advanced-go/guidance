package controller1

import (
	"fmt"
	json2 "github.com/advanced-go/stdlib/json"
)

const (
	test2Path = "file://[cwd]/changeset-test2.json"
)

var (
	set = Changeset{
		Egress: EgressChangeset{
			Insert: []EgressChange{
				{
					RouteName:         "google-search",
					FailoverScope:     "region",
					FailoverThreshold: 10,
				},
			},
		},
		Ingress: IngressChangeset{
			Insert: []IngressChange{
				{
					RouteName:        "google-search",
					RedirectLocation: "location",
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

func ExampleChangeset_Marshal() {
	buf, status := json2.Marshal(&set)
	fmt.Printf("test: Marshal() -> [status:%v] [%v]\n", status, string(buf))

	//Output:
	//fail

}

func _ExampleChangeset_Unmarshal() {
	//change, status := json2.New[Changeset](test2Path, nil)

	//fmt.Printf("test: Unmarshal() -> [status:%v] [%v]\n", status, change.Processing)

	//Output:
	//fail

}
