package cmd

import "fmt"

var MsgCmd = Commands{
	40001: testMsgMethod,
}

func testMsgMethod() {
	fmt.Println("this is test method in msg cmd")
}
