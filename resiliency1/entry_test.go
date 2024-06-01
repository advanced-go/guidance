package resiliency1

import (
	"encoding/json"
	"fmt"
)

var list = []Entry{
	{Region: "region1", Zone: "Zone1", SubZone: "s-zone", Host: "www.host1.com", Status: "active", Timeout: "100ms", RateLimit: "125", RateBurst: "25"},
	{Region: "region1", Zone: "Zone2", SubZone: "s-zone", Host: "www.host2.com", Status: "inactive", Timeout: "250ms", RateLimit: "100", RateBurst: "10"},
	{Region: "region2", Zone: "Zone1", SubZone: "s-zone", Host: "www.google.com", Status: "removed", Timeout: "3s", RateLimit: "50", RateBurst: "5"},
}

func ExampleEntry() {
	buf, err := json.Marshal(list)
	if err != nil {
		fmt.Printf("test: Entry{} -> [err:%v]\n", err)
	} else {
		fmt.Printf("test: Entry{} -> %v\n", string(buf))
	}

	//Output:
	//fail

}
