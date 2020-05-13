package websocket_test

import (
	"fmt"
	"github.com/e421083458/gin_scaffold/proto"
	"github.com/e421083458/gin_scaffold/util"
	"reflect"
	"testing"
)

func TestWsConnection_WsWrite(t *testing.T) {
	//cmd.LoadCommands()
	//fmt.Printf("all commands: %v", cmd.AllCommand)
	////fmt.Printf("state cmd:%v", )
	//return
	//websocket.BatchMsg("send batch msg test")
	//return
	//fmt.Println(len("eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJleHAiOjE1ODU5MzU2ODcsImlhdCI6MTU4NTkzMjA4N30.VcDd34gD_fEOSZjE0kZkun_Ta79OHWa6KlQpJSlssJA"))
	//passwordOk := "111111"
	//hash, err := bcrypt.GenerateFromPassword([]byte(passwordOk), bcrypt.DefaultCost)
	//if err != nil {
	//	fmt.Printf("err: %v", err)
	//}
	//encodePW := string(hash) // 保存在数据库的密码，虽然每次生成都不同，只需保存一份即可
	//fmt.Printf("pwd: %s, len: %d\n", encodePW, len(encodePW))
	// 正确密码验证
	//err = bcrypt.CompareHashAndPassword([]byte(encodePW), []byte(passwordOk))
	//if err != nil {
	//	fmt.Println("pw wrong")
	//} else {
	//	fmt.Println("pw ok")
	//}
	user := &proto.User{
		Id:123123,
		Name:"测试名称",
	}
	//l := unsafe.Sizeof(user)
	//pb := (*[1024]byte)(unsafe.Pointer(&user))
	//fmt.Println("Struct:", user)
	//fmt.Println("Bytes:", (*pb)[:l])
	//
	//buf := bytes.NewBuffer(nil)
	//enc := gob.NewEncoder(buf)
	//err := enc.Encode(user)
	//if err != nil {
	//	fmt.Printf("err:%v\n", err)
	//}
	//fmt.Printf("user:%v\n", buf.Bytes())

	byte, err := util.EncodeCmd(user); if err != nil{
		fmt.Println("err", err)
	}
	fmt.Println(byte)

	//buf := new(bytes.Buffer)
	//
	//if err := binary.Write(buf, binary.LittleEndian, *user); err != nil {
	//	fmt.Printf("err:%v\n", err)
	//}
	cmd := &proto.Cmd{
		Cid: "connectDeviceId",
		Cname:"CmdWw331",
		Wid:18018080808080,
		Data:byte,
	}
	byteCmd, err := util.EncodeCmd(cmd); if err != nil{
		fmt.Println("err", err)
	}
	fmt.Println(byteCmd)
	realCmd := &proto.Cmd{}
	if err = util.DecodeCmd(byteCmd, realCmd); err != nil{
		fmt.Println("err", err)
	}
	fmt.Println(realCmd)
	realUser := &proto.User{}
	if err = util.DecodeCmd(realCmd.Data, realUser); err != nil{
		fmt.Println("err", err)
	}
	fmt.Println(realUser)

	//bufCmd := new(bytes.Buffer)
	//bufEnc := gob.NewEncoder(bufCmd)
	//err = bufEnc.Encode(cmd)
	//if err != nil {
	//	fmt.Printf("err:%v\n", err)
	//}
	//fmt.Printf("cmd:%v\n", bufCmd.Bytes())

	//if err := binary.Write(bufCmd, binary.LittleEndian, *cmd); err != nil {
	//	fmt.Printf("err:%v\n", err)
	//}
	//fmt.Printf("bufCmd:%v\n", bufCmd)
}





func CallMethod(method interface{}){
	// here method is a interface which is a type of func
	fv := reflect.ValueOf(method)
	//args := []reflect.Value{reflect.ValueOf(a)}
	fv.Call(nil)
}