package main

import (
	"fmt"
	"github.com/gin-gonic/contrib/sessions"
	"github.com/gin-gonic/gin"
	"github.com/xiuos/mozi/routes/game"

	"github.com/xiuos/mozi/common"
	"github.com/xiuos/xlog"
	"os"
	"runtime"
)

func init() {

	savePath, err := common.Conf.String("log", "path")
	if err != nil {
		os.Exit(1)
	}
	filePrefix, err := common.Conf.String("log", "lotto_file_prefix")
	if err != nil {
		os.Exit(1)
	}
	common.Logger = xlog.NewLogger(savePath, filePrefix, xlog.Ldate|xlog.Ltime|xlog.Lmicroseconds|xlog.Llongfile, xlog.LevelInfo)
}

func main() {
	runtime.GOMAXPROCS(runtime.NumCPU())
	app := gin.Default()
	store, _ := sessions.NewRedisStore(10, "tcp", "localhost:6379", "", []byte("secret"))
	app.Use(sessions.Sessions("session", store))

	apiV1 := app.Group("/lotto")
	{
		apiV1.POST("/draw", game.DrawLotto)
	}

	port, err := common.Conf.Int("port", "lotto_port")
	if err != nil {
		os.Exit(1)
	}
	common.Logger.Info("Start lotto!")
	app.Run(fmt.Sprintf(":%d", port))
}
