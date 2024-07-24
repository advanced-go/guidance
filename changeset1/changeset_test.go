package changeset1

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
			Update: []IngressChange{
				{
					RouteName: "google-search",
					Location:  "location",
				},
			},
		},
	}
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
