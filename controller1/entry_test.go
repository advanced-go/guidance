package controller1

import (
	"encoding/json"
	"fmt"
)

var list = []Entry{
	{Region: "region1", Zone: "Zone1", SubZone: "s-zone", Host: "www.host1.com"},
	{Region: "region1", Zone: "Zone2", SubZone: "s-zone", Host: "www.host2.com"},
	{Region: "region2", Zone: "Zone1", SubZone: "s-zone", Host: "www.google.com"},
}

func ExampleEntry() {
	buf, err1 := json.Marshal(list)
	if err1 != nil {
		fmt.Printf("test: Entry{} -> [err:%v]\n", err1)
	} else {
		fmt.Printf("test: Entry{} -> %v\n", string(buf))
	}

	//Output:
	//fail

}
