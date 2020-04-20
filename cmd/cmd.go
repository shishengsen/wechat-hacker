package cmd

import (
	"fmt"
	wsClient "github.com/e421083458/gin_scaffold/websocket/client"
	"github.com/pkg/errors"
	"reflect"
)

type Commands map[int]interface{}

var AllCommand Commands

func LoadCommands() {
	AllCommand = Commands{}
	mergeCommands(AllCommand, StateCmd)
	mergeCommands(AllCommand, ConversationCmd)
	mergeCommands(AllCommand, MsgCmd)
	mergeCommands(AllCommand, SetCmd)
	mergeCommands(AllCommand, UserCmd)
	mergeCommands(AllCommand, UserLabelCmd)
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
	cmd, ok := AllCommand[msgData.Cmd]
	if !ok {
		return errors.New("invalid command code")
	}
	fv := reflect.ValueOf(cmd)
	args := []reflect.Value{reflect.ValueOf(msgData)}
	fv.Call(args)
	return nil
}
