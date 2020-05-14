package wsClient

import (
	"github.com/e421083458/gin_scaffold/proto"
	"github.com/e421083458/gin_scaffold/public"
)

type ClientMessage struct {
	Cmd *proto.Cmd `json:"cmd"`
	ClientConnId string `json:"clientConnId"`
	Wxid uint64 `json:"wxid"`
	Cname string `json:"cname"`
	Role public.RoleType `json:"role"`
}
