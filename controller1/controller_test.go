package controller1

import (
	"fmt"
	json2 "github.com/advanced-go/stdlib/json"
)

type Controller struct {
	Redirect ControllerRedirect `json:"redirect"`
	Routing  ControllerRouting  `json:"routing"`
}

var (
	testController = Controller{
		Redirect: ControllerRedirect{Location: "www.location.com", Status: "in-process"},
		Routing:  ControllerRouting{Scope: "region", Threshold: 10},
	}
)

func ExampleController() {
	buf, status := json2.Marshal(&testController)

	fmt.Printf("test: Controller() -> [status:%v] %v\n", status, string(buf))

	//Output:
	//fail
}
