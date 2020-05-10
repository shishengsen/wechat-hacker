package cmd

import (
	"fmt"
	"github.com/e421083458/gin_scaffold/proto"
)

var SetCmd = Commands{
	60001: testSetMethod,
}


func testSetMethod() {
	fmt.Println("this is test method in set cmd")
	proto_pb.Cmd.GetCid()
}
