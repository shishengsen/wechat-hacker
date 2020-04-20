package websocket_test

import (
	"fmt"
	"github.com/e421083458/gin_scaffold/cmd"
	"github.com/e421083458/gin_scaffold/websocket"
	"golang.org/x/crypto/bcrypt"
	"reflect"
	"testing"
)

func TestWsConnection_WsWrite(t *testing.T) {
	cmd.LoadCommands()
	fmt.Printf("all commands: %v", cmd.AllCommand)
	//fmt.Printf("state cmd:%v", )
	return
	websocket.BatchMsg("send batch msg test")
	return
	fmt.Println(len("eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJleHAiOjE1ODU5MzU2ODcsImlhdCI6MTU4NTkzMjA4N30.VcDd34gD_fEOSZjE0kZkun_Ta79OHWa6KlQpJSlssJA"))
	passwordOk := "111111"
	hash, err := bcrypt.GenerateFromPassword([]byte(passwordOk), bcrypt.DefaultCost)
	if err != nil {
		fmt.Printf("err: %v", err)
	}
	encodePW := string(hash) // 保存在数据库的密码，虽然每次生成都不同，只需保存一份即可
	fmt.Printf("pwd: %s, len: %d\n", encodePW, len(encodePW))
	// 正确密码验证
	err = bcrypt.CompareHashAndPassword([]byte(encodePW), []byte(passwordOk))
	if err != nil {
		fmt.Println("pw wrong")
	} else {
		fmt.Println("pw ok")
	}

}


func CallMethod(method interface{}){
	// here method is a interface which is a type of func
	fv := reflect.ValueOf(method)
	//args := []reflect.Value{reflect.ValueOf(a)}
	fv.Call(nil)
}