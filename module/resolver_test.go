package module

import (
	"fmt"
	"github.com/advanced-go/stdlib/core"
	"net/url"
)

func ExampleBuildDocumentsPath() {
	// No version, no query
	path := BuildDocumentsPath("", nil)
	fmt.Printf("test: BuildDocumentsPath(\"%v\",\"%v\") -> [%v]\n", "", "", path)

	// Version, no query
	path = BuildDocumentsPath(Ver1, nil)
	fmt.Printf("test: BuildDocumentsPath(\"%v\",\"%v\") -> [%v]\n", Ver1, "", path)

	values := make(url.Values)
	values.Add(core.RegionKey, "us")
	values.Add(core.HostKey, "www.google.com")
	values.Add(core.ZoneKey, "west")

	// No version, include query
	path = BuildDocumentsPath("", values)
	fmt.Printf("test: BuildDocumentsPath(\"%v\",\"%v\") -> [%v]\n", "", values.Encode(), path)

	// Version plus query
	path = BuildDocumentsPath(Ver2, values)
	fmt.Printf("test: BuildDocumentsPath(\"%v\",\"%v\") -> [%v]\n", Ver2, values.Encode(), path)

	//Output:
	//test: BuildDocumentsPath("","") -> [github/advanced-go/documents:resiliency]
	//test: BuildDocumentsPath("v1","") -> [github/advanced-go/documents:v1/resiliency]
	//test: BuildDocumentsPath("","host=www.google.com&region=us&zone=west") -> [github/advanced-go/documents:resiliency?host=www.google.com&region=us&zone=west]
	//test: BuildDocumentsPath("v2","host=www.google.com&region=us&zone=west") -> [github/advanced-go/documents:v2/resiliency?host=www.google.com&region=us&zone=west]
	
}
