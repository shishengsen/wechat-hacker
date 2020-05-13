package public

import (
	"crypto/md5"
	"encoding/hex"
	"github.com/gin-gonic/gin"
	"math/rand"
	"strconv"
	"strings"
	"time"
)

const WeContextKey = "_we_context"

type WeContext struct {
	// 请求唯一标识
	traceId string
	// 用户信息
	user User
	// 客户端 IP 信息
	clientIp string
	// 客户端系统
	clientSystem string

	// 存储变量
	storage map[string]interface{}
}

type RoleType int

const (
	RoleManager RoleType = iota
	RoleCustomer
	RoleClient
)

type User struct {
	// 用户ID
	Id uint32 `json:"id"`
	// 姓名
	UserName string `json:"userName"`
	// 手机号
	Phone string `json:"phone"`
	// 角色
	Role RoleType `json:"role"`
}

// 创建自定义上下文
func NewWeContext() WeContext {
	return WeContext{
		traceId: TraceId(),
	}
}

func TraceId() string {
	h := md5.New()
	rand.Seed(time.Now().UnixNano())
	h.Write([]byte(strconv.FormatInt(rand.Int63(), 10)))
	h.Write([]byte("-"))
	h.Write([]byte(strconv.FormatInt(time.Now().UnixNano(), 10)))
	h.Write([]byte("-"))
	h.Write([]byte(strconv.FormatInt(int64(rand.Int31()), 10)))
	return hex.EncodeToString(h.Sum([]byte("wehacker")))
}

func (wc *WeContext) GetTraceId() string {
	return wc.traceId
}

func (wc *WeContext) SetUser(user User) bool {
	if user.Id <= 0 {
		return false
	}

	// 同一个请求上下文中，只能被设置一次
	if wc.user.Id > 0 {
		return false
	}

	wc.user = user
	return true
}

// 获得用户ID
func (wc *WeContext) GetUserId() uint32 {
	return wc.user.Id
}

// 获得用户信息
func (wc *WeContext) GetUser() User {
	return wc.user
}

// 获得客户端IP
func (wc *WeContext) GetClientIp() string {
	return wc.clientIp
}

// 获取客户端系统
func (wc *WeContext) GetClientSystem() string {
	return wc.clientSystem
}

func (wc *WeContext) Deadline() (deadline time.Time, ok bool) {
	return
}

func (wc *WeContext) Done() <-chan struct{} {
	return nil
}

func (wc *WeContext) Err() error {
	return nil
}

func (wc *WeContext) Value(key interface{}) interface{} {
	if keyAsString, ok := key.(string); ok {
		val, _ := wc.storage[keyAsString]
		return val
	}

	return nil
}

func (wc *WeContext) Set(key string, value interface{}) {
	if len(key) == 0 {
		return
	}

	// 如果是用户信息，则存储到用户中
	if user, ok := value.(User); ok {
		wc.SetUser(user)
		return
	}

	// 延迟初始化
	if wc.storage == nil {
		wc.storage = make(map[string]interface{})
	}

	wc.storage[key] = value
}

// 将已变更的数据，存储到 gin 上下文中，继续传输
func (wc *WeContext) Storage(ctx *gin.Context) bool {
	ctx.Set(WeContextKey, *wc)
	return true
}

// 从 gin 上下文中解析自定义上下文
func ParserWeContext(ctx *gin.Context) *WeContext {
	if v, ok := ctx.Get(WeContextKey); ok {
		rctx := v.(WeContext)
		rctx.clientIp = ctx.ClientIP()
		rctx.clientSystem = GetClientSystem(ctx)
		return &rctx
	}
	return nil
}

func GetClientSystem(ctx *gin.Context) string {
	val := ctx.GetHeader("User-Agent")
	if len(val) > 0 {
		if strings.Contains(val, "android") {
			return "android"
		} else if strings.Contains(val, "iphone") || strings.Contains(val, "ipad") {
			return "ios"
		}
	}
	return ""
}



// 获取请求参数session
func GetRequestSession(ctx *gin.Context) string {
	sessionId, _ := ctx.GetQuery("sess")
	if len(sessionId) > 0 {
		return sessionId
	}

	sessionId, _ = ctx.Cookie("sess")
	if len(sessionId) > 0 {
		return sessionId
	}
	return ""
}

// 获取合法session
func GetSession(ctx *gin.Context) string {
	return ctx.GetString("_session")
}

// 设置合法session (合法的sessionId才会存入context中)
func SetSession(ctx *gin.Context, sessionId string) {
	ctx.Set("_session", sessionId)
}

func GetErrMsg(ctx *gin.Context) string {
	return ctx.GetString("_err")
}

func SetErrMsg(ctx *gin.Context, msg string) {
	ctx.Set("_err", msg)
}

