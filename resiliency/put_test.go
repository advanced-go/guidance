package resiliency

import "fmt"

func ExampleCreateEntries() {
	entries, status := createEntries[entryV1](nil, nil)

	fmt.Printf("test: CreateEntries() -> [status:%v] [entries:%v]\n", status, len(entries))

	//Output:
	//test: CreateEntries() -> [status:Invalid Content] [entries:0]

}
