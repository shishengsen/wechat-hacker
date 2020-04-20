package cmd

import "fmt"

var UserLabelCmd = Commands{
	50001: testUserLabelMethod,
}

func testUserLabelMethod() {
	fmt.Println("this is test method in userLabel cmd")
}
