package resiliency

import (
	"fmt"
	"github.com/advanced-go/stdlib/core"
	"net/url"
)

var testV1 = []entryV1{
	{Origin: core.Origin{Region: "region1", Zone: "Zone1", Host: "www.host1.com"}, Status: "active", Timeout: "100ms", RateLimit: "125", RateBurst: "25"},
	{Origin: core.Origin{Region: "region1", Zone: "Zone2", Host: "www.host2.com"}, Status: "inactive", Timeout: "250ms", RateLimit: "100", RateBurst: "10"},
	{Origin: core.Origin{Region: "region2", Zone: "Zone1", Host: "www.google.com"}, Status: "removed", Timeout: "3s", RateLimit: "50", RateBurst: "5"},
}

var testV2 = []entryV2{
	{Origin: core.Origin{Region: "region1", Zone: "Zone1", Host: "www.host1.com"}, Version: "v2", Status: "active", Timeout: "100ms", RateLimit: "125", RateBurst: "25"},
	{Origin: core.Origin{Region: "region1", Zone: "Zone2", Host: "www.host1.com"}, Version: "v2", Status: "active", Timeout: "250ms", RateLimit: "100", RateBurst: "10"},
	{Origin: core.Origin{Region: "region2", Zone: "Zone1", Host: "www.google.com"}, Version: "v2", Status: "removed", Timeout: "3s", RateLimit: "50", RateBurst: "5"},
	{Origin: core.Origin{Region: "region2", Zone: "Zone1", Host: "www.yahoo.com"}, Version: "v2", Status: "not-active", Timeout: "3s", RateLimit: "50", RateBurst: "5"},
}

func ExampleAddEntries() {
	status := addEntries[entryV1](nil, testV1)
	fmt.Printf("test: AddEntriesV1() -> [status:%v] [add:%v] [total:%v]\n", status, len(testV1), len(listV1))
	//fmt.Printf("test: ListV1 -> [%v]\n", listV1)

	status = addEntries[entryV2](nil, testV2)
	fmt.Printf("test: AddEntriesV2() -> [status:%v] [add:%v] [total:%v]\n", status, len(testV2), len(listV2))
	//fmt.Printf("test: ListV2 -> [%v]\n", listV2)

	//Output:
	//test: AddEntriesV1() -> [status:OK] [add:3] [total:3]
	//test: AddEntriesV2() -> [status:OK] [add:4] [total:4]

}

func ExampleFilterEntries() {
	entries, status := filterEntries[entryV1](nil, nil)
	fmt.Printf("test: FilterEntriesV1() -> [status:%v] [entries:%v]\n", status, len(entries))

	values := make(url.Values)
	values.Add(core.Region, "region1")
	entries, status = filterEntries[entryV1](nil, values)
	fmt.Printf("test: FilterEntriesV1() -> [status:%v] [entries:%v]\n", status, len(entries))
	//fmt.Printf("test: EntriesV1 -> [%v]\n", entries)

	entries2, status2 := filterEntries[entryV2](nil, nil)
	fmt.Printf("test: FilterEntriesV2() -> [status:%v] [entries:%v]\n", status2, len(entries2))

	//Output:
	//test: FilterEntriesV1() -> [status:Not Found] [entries:0]
	//test: FilterEntriesV1() -> [status:OK] [entries:2]
	//test: FilterEntriesV2() -> [status:Not Found] [entries:0]

}
