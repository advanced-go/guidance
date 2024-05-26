package module

import (
	"fmt"
	"github.com/advanced-go/stdlib/core"
	"net/url"
)

func ExampleBuildDocumentsPath() {
	values := make(url.Values)
	values.Add(core.RegionKey, "us")
	values.Add(core.HostKey, "www.google.com")
	values.Add(core.ZoneKey, "west")

	// No version, no query
	path := BuildDocumentsPath("", nil)
	fmt.Printf("test: BuildDocumentsPath(\"%v\",\"%v\") -> [%v]\n", "", "", path)

	// Version, no query
	path = BuildDocumentsPath(Ver1, nil)
	fmt.Printf("test: BuildDocumentsPath(\"%v\",\"%v\") -> [%v]\n", Ver1, "", path)

	// No version, include query
	path = BuildDocumentsPath("", values)
	fmt.Printf("test: BuildDocumentsPath(\"%v\",\"%v\") -> [%v]\n", "", values.Encode(), path)

	// Version plus query
	path = BuildDocumentsPath(Ver2, values)
	fmt.Printf("test: BuildDocumentsPath(\"%v\",\"%v\") -> [%v]\n", Ver2, values.Encode(), path)

	//Output:
	//test: BuildDocumentsPath("","") -> [github/advanced-go/documents:resiliency1]
	//test: BuildDocumentsPath("v1","") -> [github/advanced-go/documents:v1/resiliency1]
	//test: BuildDocumentsPath("","host=www.google.com&region=us&zone=west") -> [github/advanced-go/documents:resiliency1?host=www.google.com&region=us&zone=west]
	//test: BuildDocumentsPath("v2","host=www.google.com&region=us&zone=west") -> [github/advanced-go/documents:v2/resiliency1?host=www.google.com&region=us&zone=west]

}

func ExampleBuildDocumentsResource() {
	values := make(url.Values)
	values.Add(core.RegionKey, "us")
	values.Add(core.HostKey, "www.google.com")
	values.Add(core.ZoneKey, "west")

	// Resource path only
	path := BuildDocumentsResource("", nil)
	fmt.Printf("test: BuildDocumentsResource(\"%v\",\"%v\") -> [%v]\n", "", "", path)

	// Resource path only
	path = BuildDocumentsResource("", values)
	fmt.Printf("test: BuildDocumentsResource(\"%v\",\"%v\") -> [%v]\n", "", values.Encode(), path)

	// Resource path only
	path = BuildDocumentsResource(Ver1, nil)
	fmt.Printf("test: BuildDocumentsResource(\"%v\",\"%v\") -> [%v]\n", Ver1, "", path)

	path = BuildDocumentsResource(Ver2, values)
	fmt.Printf("test: BuildDocumentsResource(\"%v\",\"%v\") -> [%v]\n", Ver2, values.Encode(), path)

	//Output:
	//test: BuildDocumentsResource("","") -> [resiliency1]
	//test: BuildDocumentsResource("","host=www.google.com&region=us&zone=west") -> [resiliency1?host=www.google.com&region=us&zone=west]
	//test: BuildDocumentsResource("v1","") -> [v1/resiliency1]
	//test: BuildDocumentsResource("v2","host=www.google.com&region=us&zone=west") -> [v2/resiliency1?host=www.google.com&region=us&zone=west]

}
