package wsClient

type ClientMessage struct {
	Cmd  int      `json:"cmd"`
	Wxid uint64      `json:"wxid"`
	Data interface{} `json:"data"`
	ClientConnId string `json:"clientConnId"`
}
