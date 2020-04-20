package cmd

import "fmt"

var UserCmd = Commands{
	20001: testUserMethod,
}

func testUserMethod() {
	fmt.Println("this is test method in user cmd")
}
