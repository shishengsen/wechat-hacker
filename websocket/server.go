package websocket

import (
	"errors"
	"fmt"
	"github.com/e421083458/gin_scaffold/cmd"
	"github.com/e421083458/gin_scaffold/dao"
	"github.com/e421083458/gin_scaffold/proto"
	"github.com/e421083458/gin_scaffold/public"
	"github.com/e421083458/gin_scaffold/util"
	"github.com/e421083458/gin_scaffold/websocket/client"
	"github.com/e421083458/golang_common/lib"
	"github.com/garyburd/redigo/redis"
	"github.com/gorilla/websocket"
	uuId "github.com/satori/go.uuId"
	"log"
	"net/http"
	"strconv"
	"sync"
	"time"
)

// http升级websocket协议的配置
var wsUpgrader = websocket.Upgrader{
	// 允许所有CORS跨域请求
	CheckOrigin: func(r *http.Request) bool {
		return true
	},
}

// 客户端读写消息
type Message struct {
	MessageType int
	Data        []byte
}

// 客户端管理员
type ClientManager struct {
	Clients    map[*Client]bool
	broadCast  chan *Message
	register   chan *Client
	unregister chan *Client
}

// 客户端
type Client struct {
	Id        string          // 客户端ID
	wxId      uint64          // 客户端微信ID
	socket    *websocket.Conn // 客户端连接
	inChan    chan *Message   // 写channel(发送消息的管道)
	outChan   chan *Message   // 读channel(接收消息管道)
	mutex     sync.Mutex      // 避免重复关闭管道
	isClosed  bool
	closeChan chan byte // 关闭通知
}

var Manager = ClientManager{
	broadCast:  make(chan *Message),
	register:   make(chan *Client),
	unregister: make(chan *Client),
	Clients:    make(map[*Client]bool),
}

func (manager *ClientManager) start() {
	for {
		select {
		case conn := <-manager.register:
			manager.Clients[conn] = true
			connectMsg := &Message{MessageType: websocket.TextMessage, Data: []byte("/A new socket has connected.")}
			manager.send(connectMsg, conn)
			log.Printf(" [websocket] A socket connected: %s\n", conn.Id)
		case conn := <-manager.unregister:
			if _, ok := manager.Clients[conn]; ok {
				disconnectMsg := &Message{MessageType: websocket.TextMessage, Data: []byte("/A socket has disconnected.")}
				manager.send(disconnectMsg, conn)
				conn.wsClose()
				delete(manager.Clients, conn)
				log.Printf(" [websocket] A socket disconnected: %s\n", conn.Id)
			}
		case message := <-manager.broadCast:
			for conn := range manager.Clients {
				select {
				case conn.outChan <- message:
				default:
					conn.wsClose()
					delete(manager.Clients, conn)
				}
			}
		}
	}
}

func (manager *ClientManager) send(message *Message, ignore *Client) {
	for conn := range manager.Clients {
		if conn != ignore {
			conn.outChan <- message
		}
	}
}

func (c *Client) wsReadLoop() {
	for {
		//读一个message
		msgType, data, err := c.socket.ReadMessage()
		if err != nil {
			goto error
		}
		req := &Message{
			msgType,
			data,
		}
		runCmd(c, req)
		//放入请求队列-将原样返回输入的数据
		select {
		case c.inChan <- req:
		case <-c.closeChan:
			goto closed
		}
	}
error:
	c.wsClose()
closed:
}

// 存储注册客户端数据至redis
func runCmd(c *Client, msg *Message) {
	cmdObj := &proto.Cmd{}
	if err := util.DecodeCmd(msg.Data, cmdObj); err != nil{
		log.Printf(" [redis] decode cmd failed: %v", err)
	}
	execCmd := &wsClient.ClientMessage{
		Cmd:cmdObj,
		Wxid:uint64(cmdObj.Wid),
		ClientConnId:c.Id,
		Cname:cmdObj.Cname,
	}

	// 若为确认连接则存储至redis
	c.wxId = uint64(cmdObj.Wid)
	err := cmd.CallCmd(execCmd)
	if err != nil {
		log.Printf(" [redis] run command connect fail: %v", err)
	}

}

// 从redis中移除注册客户端
func redisClose(c *Client) {
	msgObj := &wsClient.ClientMessage{
		ClientConnId: c.Id,
		Wxid:         c.wxId,
		Cmd:          &proto.Cmd{
			Cname:"CmdWw322",
		},
		Cname:"CmdWw322",
	}
	if err := cmd.CallCmd(msgObj); err != nil {
		log.Printf(" [redis] run command disconnect fail: %v", err)
	}
}

func (c *Client) wsWriteLoop() {
	for {
		select {
		// 取一个应答
		case msg := <-c.outChan:
			// 写给websocket
			if err := c.socket.WriteMessage(msg.MessageType, msg.Data); err != nil {
				goto error
			}
		case <-c.closeChan:
			goto closed
		}
	}
error:
	c.wsClose()
closed:
}

// 发送心跳检测
func (c *Client) procLoop() {
	// 启动一个goroutine发送心跳
	go func() {
		for {
			time.Sleep(2 * time.Second)
			if err := c.wsWrite(websocket.TextMessage, []byte("pong")); err != nil {
				log.Printf(" [websocket] heartbeat fail")
				c.wsClose()
				break
			}
		}
	}()

	// 启动一个goroutine读取redis中单聊消息并写入channel
	redC, err := public.NewRedisConn()
	if err == nil {
		go func() {
			for {
				time.Sleep(1 * time.Second)
				// 从redis读取消息发到客户端
				wxId, err := redis.String(redC.Do("get", c.Id))
				if err != nil || wxId == "" {
					continue
				}
				singleMsgQueue := wxId + "-msg-queue"
				data, err := redis.String(redC.Do("rpop", singleMsgQueue))
				if err != nil || data == "" {
					continue
				}
				if err := c.wsWrite(websocket.TextMessage, []byte(data)); err != nil {
					log.Printf(" [websocket] read single msg fail")
					c.wsClose()
					break
				}
			}
		}()
	}

	go func() {
		for {
			msg, err := c.wsRead()
			if err != nil {
				log.Printf(" [websocket] client:%s read fail", c.Id)
				break
			}
			err = c.wsWrite(msg.MessageType, msg.Data)
			if err != nil {
				log.Printf(" [websocket] client:%s write fail", c.Id)
				break
			}
		}
	}()
}

//MessageType int `json:"messageType"` // 消息类型：1-文本，2-图片，3-图文链接，4-音频，5-视频
//Content interface{} `json:"content"`

// 客户端消息写入队列
func (c *Client) wsWrite(messageType int, data []byte) error {
	select {
	case c.outChan <- &Message{messageType, data,}:
	case <-c.closeChan:
		return errors.New("websocket closed")
	}
	return nil
}

// 客户端读取消息队列
func (c *Client) wsRead() (*Message, error) {
	select {
	case msg := <-c.inChan:
		return msg, nil
	case <-c.closeChan:
	}
	return nil, errors.New("websocket closed")
}

// 关闭连接
func (c *Client) wsClose() {
	redisClose(c)
	c.socket.Close()
	c.mutex.Lock()
	defer c.mutex.Unlock()
	if !c.isClosed {
		c.isClosed = true
		close(c.closeChan)
	}
}

func Run() {
	go Manager.start()
	http.HandleFunc("/wsClient", wsClientHandler)
	http.HandleFunc("/wsCustomer", wsCustomerHandler)
	log.Printf(" [websocket] WebsocketServerRun:%s\n", lib.GetStringConf("base.websocket.addr"))
	err := http.ListenAndServe(lib.GetStringConf("base.websocket.addr"), nil)
	if err != nil {
		log.Printf(" [websocket error] WebsocketServerRun: error %s\n", err.Error())
	}
}

func SingleMsg(wxid uint64, msg string) {
	redC, err := public.NewRedisConn()
	if err != nil {
		log.Printf("[redis error] new redis connect err: %v", err)
		return
	}
	defer redC.Close()
	_, err = redC.Do("lpush", strconv.Itoa(int(wxid))+"-msg-queue", msg)
	if err != nil {
		fmt.Printf("push err: %v", err)
	}
}

func BatchMsg(msg string) {
	redC, err := public.NewRedisConn()
	if err != nil {
		log.Printf("[redis error] new redis connect err: %v", err)
		return
	}
	clientWxids, err := redis.Int64s(redC.Do("smembers", public.RegisterKey))
	if err != nil {
		log.Printf("[redis error] new redis connect err: %v", err)
		return
	}
	for _, wxid := range clientWxids {
		SingleMsg(uint64(wxid), msg)
		fmt.Printf("client wxid: %v\n", wxid)
	}
}

func wsClientHandler(resp http.ResponseWriter, req *http.Request) {
	respHeader := http.Header{}
	//token 校验
	token := req.Header.Get("Sec-WebSocket-Protocol")
	if len(token) == 0 {
		log.Printf(" [websocket] handle error with no token")
		return
	}
	req.Header.Add("Authorization", "Bearer "+token)
	log.Printf(" [websocket] token:%s", token)
	// 应答客户端告知升级连接为websocket
	respHeader.Add("Sec-WebSocket-Protocol", token)
	req.Header.Set("Authorization", "Bearer "+token)
	verifyRs := util.ValidateWsToken(req, public.RoleClient)
	if !verifyRs {
		log.Printf(" [websocket] shake hands error: invalid token %s", token)
		return
	}
	socket, err := wsUpgrader.Upgrade(resp, req, respHeader)
	if err != nil {
		return
	}
	clientId := uuId.NewV4().String()
	c := &Client{
		Id:        clientId,
		wxId:      0,
		socket:    socket,
		inChan:    make(chan *Message, 1000),
		outChan:   make(chan *Message, 1000),
		closeChan: make(chan byte),
		isClosed:  false,
	}
	Manager.register <- c

	// 处理器
	go c.procLoop()
	// 读协程
	go c.wsReadLoop()
	// 写协程
	go c.wsWriteLoop()
}

func wsCustomerHandler(resp http.ResponseWriter, req *http.Request) {
	respHeader := http.Header{}
	//token 校验
	token := req.Header.Get("Sec-WebSocket-Protocol")
	if len(token) == 0 {
		log.Printf(" [websocket] handle error with no token")
		return
	}
	req.Header.Add("Authorization", "Bearer "+token)
	log.Printf(" [websocket] token:%s", token)
	// 应答客户端告知升级连接为websocket
	respHeader.Add("Sec-WebSocket-Protocol", token)
	req.Header.Set("Authorization", "Bearer "+token)
	verifyRs := util.ValidateWsToken(req, public.RoleCustomer)
	if !verifyRs {
		log.Printf(" [websocket] shake hands error: invalid token %s", token)
		return
	}
	// 根据token从DB查出客户ID并将wxid设置为客户ID
	var userDao = &dao.User{}
	user, err := userDao.FindByToken(nil, token); if err != nil {
		log.Printf(" [websocket] customer not found:  %s", token)
		return
	}
	socket, err := wsUpgrader.Upgrade(resp, req, respHeader)
	if err != nil {
		return
	}
	clientId := uuId.NewV4().String()
	c := &Client{
		Id:        clientId,
		wxId:      uint64(user.CustomerId),
		socket:    socket,
		inChan:    make(chan *Message, 1000),
		outChan:   make(chan *Message, 1000),
		closeChan: make(chan byte),
		isClosed:  false,
	}
	Manager.register <- c

	// 处理器
	go c.procLoop()
	// 读协程
	go c.wsReadLoop()
	// 写协程
	go c.wsWriteLoop()
}
