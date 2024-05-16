package http

import "fmt"

func ExampleExchange_Invalid() {
	resp, status := Exchange(nil)

	fmt.Printf("test: Exchange(nil) -> [status:%v] [status-code:%v]\n", status, resp.StatusCode)

	//Output:
	//test: Exchange(nil) -> [status:Invalid Argument [error: request is nil]] [status-code:500]

}
