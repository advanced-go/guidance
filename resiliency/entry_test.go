package resiliency

import (
	"encoding/json"
	"fmt"
	"github.com/advanced-go/stdlib/core"
	"github.com/advanced-go/stdlib/httpx"
	"github.com/advanced-go/stdlib/io"
	"net/http"
	"net/url"
)

const (
	entryV1Path    = "file://[cwd]/resiliencytest/entry-v1.json"
	entryV2Path    = "file://[cwd]/resiliencytest/entry-v2.json"
	emptyPath      = "file://[cwd]/resiliencytest/empty.json"
	emptyEntryPath = "file://[cwd]/resiliencytest/entry-empty.json"
)

var testV1 = []EntryV1{
	{Origin: core.Origin{Region: "region1", Zone: "Zone1", Host: "www.host1.com"}, Status: "active", Timeout: "100ms", RateLimit: "125", RateBurst: "25"},
	{Origin: core.Origin{Region: "region1", Zone: "Zone2", Host: "www.host2.com"}, Status: "inactive", Timeout: "250ms", RateLimit: "100", RateBurst: "10"},
	{Origin: core.Origin{Region: "region2", Zone: "Zone1", Host: "www.google.com"}, Status: "removed", Timeout: "3s", RateLimit: "50", RateBurst: "5"},
}

var testV2 = []EntryV2{
	{Origin: core.Origin{Region: "region1", Zone: "Zone1", Host: "www.host1.com"}, Version: "v2", Status: "active", Timeout: "100ms", RateLimit: "125", RateBurst: "25"},
	{Origin: core.Origin{Region: "region1", Zone: "Zone2", Host: "www.host1.com"}, Version: "v2", Status: "active", Timeout: "250ms", RateLimit: "100", RateBurst: "10"},
	{Origin: core.Origin{Region: "region2", Zone: "Zone1", Host: "www.google.com"}, Version: "v2", Status: "removed", Timeout: "3s", RateLimit: "50", RateBurst: "5"},
	{Origin: core.Origin{Region: "region2", Zone: "Zone1", Host: "www.yahoo.com"}, Version: "v2", Status: "not-active", Timeout: "3s", RateLimit: "50", RateBurst: "5"},
}

func ExampleCreateEntries() {
	buf, status := io.ReadFile(entryV1Path)
	fmt.Printf("test: ReadFile(\"%v\") -> [status:%v] [buff:%v]\n", entryV1Path, status, len(buf))

	entries1, status1 := createEntries[EntryV1](nil, nil)
	fmt.Printf("test: CreateEntries() -> [status:%v] [entries:%v]\n", status1, len(entries1))

	entries1, status1 = createEntries[EntryV1](nil, buf)
	fmt.Printf("test: CreateEntries() -> [status:%v] [entries:%v]\n", status1, len(entries1))

	buf, status = io.ReadFile(entryV2Path)
	fmt.Printf("test: ReadFile(\"%v\") -> [status:%v] [buff:%v]\n", entryV2Path, status, len(buf))

	entries2, status2 := createEntries[EntryV2](nil, buf)
	fmt.Printf("test: CreateEntries() -> [status:%v] [entries:%v]\n", status2, len(entries2))

	//Output:
	//test: ReadFile("file://[cwd]/resiliencytest/entry-v1.json") -> [status:OK] [buff:901]
	//test: CreateEntries() -> [status:Invalid Content] [entries:0]
	//test: CreateEntries() -> [status:OK] [entries:3]
	//test: ReadFile("file://[cwd]/resiliencytest/entry-v2.json") -> [status:OK] [buff:1291]
	//test: CreateEntries() -> [status:OK] [entries:4]

}

func ExampleAddEntries() {
	status := addEntries[EntryV1](nil, nil, testV1)
	fmt.Printf("test: AddEntriesV1() -> [status:%v] [add:%v] [total:%v]\n", status, len(testV1), len(listV1))
	//fmt.Printf("test: ListV1 -> [%v]\n", listV1)

	status = addEntries[EntryV2](nil, nil, testV2)
	fmt.Printf("test: AddEntriesV2() -> [status:%v] [add:%v] [total:%v]\n", status, len(testV2), len(listV2))
	//fmt.Printf("test: ListV2 -> [%v]\n", listV2)

	//Output:
	//test: AddEntriesV1() -> [status:OK] [add:3] [total:3]
	//test: AddEntriesV2() -> [status:OK] [add:4] [total:4]

}

func ExampleFilterEntries() {
	entries, status := filterEntries[EntryV1](nil, listV1, nil)
	fmt.Printf("test: FilterEntriesV1() -> [status:%v] [entries:%v]\n", status, len(entries))

	values := make(url.Values)
	values.Add(core.RegionKey, "regIon1")
	entries, status = filterEntries[EntryV1](nil, listV1, values)
	fmt.Printf("test: FilterEntriesV1() -> [status:%v] [entries:%v]\n", status, len(entries))
	//fmt.Printf("test: EntriesV1 -> [%v]\n", entries)

	entries2, status2 := filterEntries[EntryV2](nil, listV2, nil)
	fmt.Printf("test: FilterEntriesV2() -> [status:%v] [entries:%v]\n", status2, len(entries2))

	values.Set(core.RegionKey, "regIon2")
	values.Add(core.ZoneKey, "zone1")
	values.Add(core.HostKey, "www.google.com")
	entries2, status2 = filterEntries[EntryV2](nil, listV2, values)
	fmt.Printf("test: FilterEntriesV2() -> [status:%v] [entries:%v]\n", status2, entries2)

	//Output:
	//test: FilterEntriesV1() -> [status:OK] [entries:3]
	//test: FilterEntriesV1() -> [status:OK] [entries:2]
	//test: FilterEntriesV2() -> [status:OK] [entries:4]
	//test: FilterEntriesV2() -> [status:OK] [entries:[{{region2 Zone1  www.google.com } v2 removed   3s 50 5}]]

}

func ExampleGetEntries_Empty() {
	h := make(http.Header)
	h.Add(httpx.ContentLocation, emptyPath)
	entries, status := getEntries[EntryV1](nil, h, nil)
	fmt.Printf("test: GetEntries() -> [status:%v] [entries:%v]\n", status, len(entries))

	h.Set(httpx.ContentLocation, emptyEntryPath)
	entries, status = getEntries[EntryV1](nil, h, nil)
	fmt.Printf("test: GetEntries() -> [status:%v] [entries:%v]\n", status, len(entries))

	//Output:
	//test: GetEntries() -> [status:Not Found] [entries:0]
	//test: GetEntries() -> [status:Not Found] [entries:0]

}

func ExampleGetEntries_V1() {
	values := make(url.Values)
	values.Add(core.RegionKey, "region1")
	entries1, status1 := getEntries[EntryV1](nil, nil, values)
	fmt.Printf("test: GetEntries() -> [status:%v] [entries:%v]\n", status1, len(entries1))

	buf, status := io.ReadFile(entryV1Path)
	fmt.Printf("test: ReadFile(\"%v\") -> [status:%v] [buff:%v]\n", entryV1Path, status, len(buf))

	h := make(http.Header)
	h.Add(httpx.ContentLocation, entryV1Path)
	entries1, status1 = getEntries[EntryV1](nil, h, nil)
	fmt.Printf("test: GetEntries() -> [status:%v] [entries:%v]\n", status1, len(entries1))

	//Output:
	//test: GetEntries() -> [status:OK] [entries:2]
	//test: ReadFile("file://[cwd]/resiliencytest/entry-v1.json") -> [status:OK] [buff:901]
	//test: GetEntries() -> [status:OK] [entries:3]

}

func ExampleGetEntries_V2() {
	values := make(url.Values)
	values.Add(core.ZoneKey, "zonE1")
	entries1, status1 := getEntries[EntryV2](nil, nil, values)
	fmt.Printf("test: GetEntries() -> [status:%v] [entries:%v]\n", status1, len(entries1))

	buf, status := io.ReadFile(entryV2Path)
	fmt.Printf("test: ReadFile(\"%v\") -> [status:%v] [buff:%v]\n", entryV2Path, status, len(buf))

	values = make(url.Values)
	values.Add(httpx.ContentLocation, entryV2Path)
	h := make(http.Header)
	h.Add(httpx.ContentLocation, entryV2Path)
	entries1, status1 = getEntries[EntryV2](nil, h, nil)
	fmt.Printf("test: GetEntries() -> [status:%v] [entries:%v]\n", status1, len(entries1))

	//Output:
	//test: GetEntries() -> [status:OK] [entries:3]
	//test: ReadFile("file://[cwd]/resiliencytest/entry-v2.json") -> [status:OK] [buff:1291]
	//test: GetEntries() -> [status:OK] [entries:4]

}

func _ExampleOutput() {
	//buff, err := json.Marshal(testV1)
	//fmt.Printf("%v\n", err)
	//fmt.Printf("%v\n", string(buff))

	buff, err := json.Marshal(testV2)
	fmt.Printf("%v\n", err)
	fmt.Printf("%v\n", string(buff))

	//Output:
	//fail
}
