package main

import (
	"github.com/e421083458/gin_scaffold/cmd"
	"github.com/e421083458/gin_scaffold/public"
	"github.com/e421083458/gin_scaffold/router"
	"github.com/e421083458/gin_scaffold/websocket"
	"github.com/e421083458/golang_common/lib"
	"os"
	"os/signal"
	"syscall"
)

func main()  {
	lib.InitModule("./conf/dev/",[]string{"base","mysql","redis",})
	defer lib.Destroy()
	public.InitMysql()
	cmd.LoadCommands()
	public.InitValidate()
	router.HttpServerRun()
	websocket.Run()
	quit := make(chan os.Signal)
	signal.Notify(quit, syscall.SIGKILL, syscall.SIGQUIT, syscall.SIGINT, syscall.SIGTERM)
	<-quit

	router.HttpServerStop()
}