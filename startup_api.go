package main

import (
	"fmt"
	"github.com/xiuos/mozi/common"
	"github.com/xiuos/mozi/routes/api"
	"github.com/xiuos/xlog"
	"os"
	"runtime"
	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/contrib/sessions"
)

func init() {

	savePath, err := common.Conf.String("log", "path")
	if err != nil {
		os.Exit(1)
	}
	filePrefix, err := common.Conf.String("log", "api_file_prefix")
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


	apiV1 := app.Group("/api/v1")
	{
		apiV1.POST("/register", api.Register)
		apiV1.POST("/login", api.Login)
	}

	apiAuthV1 := app.Group("/api/v1")
	{
		apiAuthV1.POST("/get_balance", api.Register)
	}


	port, err := common.Conf.Int("port", "api_port")
	if err != nil {
		os.Exit(1)
	}
	common.Logger.Info("Start API!")
	app.Run(fmt.Sprintf(":%d", port))
}
