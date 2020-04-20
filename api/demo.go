package api

type Demo struct {
}

//func DemoRegister(router *gin.RouterGroup) {
//	demo := Demo{}
//	router.GET("/index", demo.Index)
//	router.GET("/bind", demo.Bind)
//	router.GET("/dao", demo.Dao)
//	router.GET("/redis", demo.Redis)
//}

//func (demo *Demo) Index(c *gin.Context) {
//	clientId,_:= c.GetQuery("clientId")
//	message,_:= c.GetQuery("message")
//
//	websocket.SinglePush(clientId, &websocket.Message{Data: []byte(message)})
//	manager := &websocket.ClientManager{}
//	//var id = ""
//	//for c := range manager.Clients {
//	//	id = c.Id
//	//}
//	middleware.ResponseSuccess(c, len(manager.Clients))
//	return
//}

//func (demo *Demo) Dao(c *gin.Context) {
//	if area,err:=(&dao.Area{}).Find(c,c.DefaultQuery("id","1"));err!=nil{
//		middleware.ResponseError(c,501,err)
//	}else{
//		js,_:=json.Marshal(area)
//		middleware.ResponseSuccess(c, string(js))
//	}
//	return
//}

//func (demo *Demo) Redis(c *gin.Context) {
//	redisKey:="redis_key"
//	lib.RedisConfDo(public.GetTraceContext(c),"default",
//		"SET",redisKey,"redis_value")
//	redisValue,err:=redis.String(lib.RedisConfDo(public.GetTraceContext(c),"default",
//		"GET",redisKey))
//	if err!=nil{
//		middleware.ResponseError(c,501,err)
//		return
//	}
//	middleware.ResponseSuccess(c, redisValue)
//	return
//}
//
//func (demo *Demo) Bind(c *gin.Context) {
//	st:=&dto.InStruct{}
//	if err:=st.BindingValidParams(c);err!=nil{
//		middleware.ResponseError(c,500,err)
//		return
//	}
//	middleware.ResponseSuccess(c, "")
//	return
//}