package module

import (
	"fmt"
	"github.com/advanced-go/stdlib/uri"
)

func ExampleBuildDocumentsPath() {
	values := uri.BuildValues("region=us&host=www.google.com&zone=*")
	//values.Add(core.RegionKey, "us")
	//values.Add(core.HostKey, "www.google.com")
	//values.Add(core.ZoneKey, "west")

	// No version, no query
	path := BuildDocumentsPath("", nil)
	fmt.Printf("test: BuildDocumentsPath(\"%v\",\"%v\") -> [%v]\n", "", "", path)

	// Version, no query
	path = BuildDocumentsPath(Ver1, nil)
	fmt.Printf("test: BuildDocumentsPath(\"%v\",\"%v\") -> [%v]\n", Ver1, "", path)

	// No version, include query
	path = BuildDocumentsPath("", values)
	fmt.Printf("test: BuildDocumentsPath(\"%v\",\"%v\") -> [%v]\n", "", values, path)

	// Version plus query
	path = BuildDocumentsPath(Ver2, values)
	fmt.Printf("test: BuildDocumentsPath(\"%v\",\"%v\") -> [%v]\n", Ver2, values, path)

	//Output:
	//test: BuildDocumentsPath("","") -> [/github/advanced-go/documents:resiliency]
	//test: BuildDocumentsPath("v1","") -> [/github/advanced-go/documents:v1/resiliency]
	//test: BuildDocumentsPath("","map[host:[www.google.com] region:[us] zone:[*]]") -> [/github/advanced-go/documents:resiliency?host=www.google.com&region=us&zone=%2A]
	//test: BuildDocumentsPath("v2","map[host:[www.google.com] region:[us] zone:[*]]") -> [/github/advanced-go/documents:v2/resiliency?host=www.google.com&region=us&zone=%2A]

}

/*
func ExampleBuildDocumentsResource() {
	values := make(url.Values)
	values.Add(core.RegionKey, "us")
	values.Add(core.HostKey, "www.google.com")
	values.Add(core.ZoneKey, "*")

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
	//test: BuildDocumentsResource("","") -> [resiliency]
	//test: BuildDocumentsResource("","host=www.google.com&region=us&zone=%2A") -> [resiliency?host=www.google.com&region=us&zone=%2A]
	//test: BuildDocumentsResource("v1","") -> [v1/resiliency]
	//test: BuildDocumentsResource("v2","host=www.google.com&region=us&zone=%2A") -> [v2/resiliency?host=www.google.com&region=us&zone=%2A]

}


*/
