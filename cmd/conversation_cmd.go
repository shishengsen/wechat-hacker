package cmd

import "fmt"

var ConversationCmd = Commands{
	30001: testConversationMethod,
}

func testConversationMethod() {
	fmt.Println("this is test method from conversation cmd")
}
