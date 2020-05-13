package cmd

import (
	"fmt"
	wsClient "github.com/e421083458/gin_scaffold/websocket/client"
	"github.com/pkg/errors"
	"reflect"
)

type Commands map[string]interface{}

var AllCommand Commands

func LoadCommands() {
	AllCommand = Commands{}
	mergeCommands(AllCommand, ContactCmd)
	mergeCommands(AllCommand, GroupCmd)
	mergeCommands(AllCommand, InfoCmd)
	mergeCommands(AllCommand, MsgCmd)
}

func mergeCommands(cmdMap1 Commands, cmdMap2 Commands) {
	for k, v := range cmdMap2 {
		if _, ok := cmdMap2[k]; ok {
			cmdMap1[k] = v
		}
	}
}

func CallCmd(msgData *wsClient.ClientMessage) error {
	fmt.Printf("ready to run cmd: %v\n", msgData)
	cmd, ok := AllCommand[msgData.Cname]
	if !ok {
		return errors.New("invalid command name:" + msgData.Cname)
	}
	fv := reflect.ValueOf(cmd)
	args := []reflect.Value{reflect.ValueOf(msgData)}
	fv.Call(args)
	return nil
}
