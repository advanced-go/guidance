package resiliency

import (
	"fmt"
	"github.com/advanced-go/stdlib/uri"
	"net/url"
)

func ExampleExpansion() {
	host := ""
	path := ""
	vers := ""
	var values url.Values

	url := Expansion(host, path, vers, values)
	fmt.Printf("test: Expansion(\"%v\",\"%v\",\"%v\",\"%v\") -> [uri:%v]\n", host, path, vers, values, url)

	path = documentsPath
	url = Expansion(host, path, vers, values)
	fmt.Printf("test: Expansion(\"%v\",\"%v\",\"%v\",\"%v\") -> [uri:%v]\n", host, path, vers, values, url)

	host = "localhost:8081"
	url = Expansion(host, path, vers, values)
	fmt.Printf("test: Expansion(\"%v\",\"%v\",\"%v\",\"%v\") -> [uri:%v]\n", host, path, vers, values, url)

	vers = documentsV1
	url = Expansion(host, path, vers, values)
	fmt.Printf("test: Expansion(\"%v\",\"%v\",\"%v\",\"%v\") -> [uri:%v]\n", host, path, vers, values, url)

	host = "www.google.com"
	vers = "v2"
	values = uri.BuildValues("region=*&zone=west&sub-zone=texas")
	url = Expansion(host, path, vers, values)
	fmt.Printf("test: Expansion(\"%v\",\"%v\",\"%v\",\"%v\") -> [uri:%v]\n", host, path, vers, values, url)

	//Output:
	//test: Expansion("","","","map[]") -> [uri:%!(EXTRA string=)]
	//test: Expansion("","/github/advanced-go/documents:%sresiliency","","map[]") -> [uri:/github/advanced-go/documents:resiliency]
	//test: Expansion("localhost:8081","/github/advanced-go/documents:%sresiliency","","map[]") -> [uri:http://localhost:8081/github/advanced-go/documents:resiliency]
	//test: Expansion("localhost:8081","/github/advanced-go/documents:%sresiliency","v1","map[]") -> [uri:http://localhost:8081/github/advanced-go/documents:v1/resiliency]
	//test: Expansion("www.google.com","/github/advanced-go/documents:%sresiliency","v2","map[region:[*] sub-zone:[texas] zone:[west]]") -> [uri:https://www.google.com/github/advanced-go/documents:v2/resiliency?region=%2A&sub-zone=texas&zone=west]

}
