package middleware

import (
	"encoding/json"
	"github.com/e421083458/gin_scaffold/public"
	"github.com/gin-gonic/gin"
)

type ResponseCode int

//1000以下为通用码，1000以上为用户自定义码
const (
	SuccessCode ResponseCode = iota
	UndefErrorCode
	ValidErrorCode
	InternalErrorCode

	InvalidRequestErrorCode ResponseCode = 401
	CustomizeCode           ResponseCode = 1000

	GROUPALL_SAVE_FLOWERROR ResponseCode = 2001
)

type Response struct {
	ErrorCode ResponseCode `json:"errCode"`
	ErrorMsg  string       `json:"errMsg"`
	Data      interface{}  `json:"data"`
	TraceId   interface{}  `json:"traceId"`
}

func ResponseError(c *gin.Context, code ResponseCode, err error) {
	//trace, _ := c.Get("trace")
	//traceContext, _ := trace.(*lib.TraceContext)
	traceId := ""
	traceContext := public.ParserWeContext(c)
	if traceContext != nil {
		traceId = traceContext.GetTraceId()
	}
	c.Header("X-Trace-Id", traceId)
	resp := &Response{ErrorCode: code, ErrorMsg: err.Error(), Data: "", TraceId: traceId}
	c.JSON(200, resp)
	response, _ := json.Marshal(resp)
	c.Set("response", string(response))
	c.AbortWithError(200, err)
}

func ResponseSuccess(c *gin.Context, data interface{}) {
	traceId := ""
	traceContext := public.ParserWeContext(c)
	if traceContext != nil {
		traceId = traceContext.GetTraceId()
	}
	resp := &Response{ErrorCode: SuccessCode, ErrorMsg: "", Data: data, TraceId: traceId}
	c.JSON(200, resp)
	response, _ := json.Marshal(resp)
	c.Set("response", string(response))
}
