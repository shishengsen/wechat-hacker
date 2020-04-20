package cmd

import "fmt"

var SetCmd = Commands{
	60001: testSetMethod,
}


func testSetMethod() {
	fmt.Println("this is test method in set cmd")
}
