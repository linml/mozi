package main

import (
	"fmt"
	"github.com/gin-gonic/contrib/sessions"
	"github.com/gin-gonic/gin"
	"github.com/xiuos/mozi/common"
	"github.com/xiuos/mozi/routes"
	"github.com/xiuos/mozi/routes/api"
	"github.com/xiuos/mozi/xlog"
	"os"
	"runtime"
)

func init() {

	savePath, err := common.Conf.String("log", "path")
	if err != nil {
		fmt.Println("日志文件错误 ", err)
		os.Exit(1)
	}
	filePrefix, err := common.Conf.String("log", "guest_api_file_prefix")
	if err != nil {
		fmt.Println("日志文件头错误 ", err)
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
		apiV1.POST("/login/guest", api.GuestLogin)
	}

	apiAuthV1 := app.Group("/api/v1")
	apiAuthV1.Use(routes.APIGuestAuthMiddleWare())
	{
		apiAuthV1.GET("/get_balance", api.GetGuestBalance)
		apiAuthV1.GET("/infos", api.GetGuestInfos)
		apiAuthV1.POST("/lotto/bet", api.GuestBet)
	}

	port, err := common.Conf.Int("port", "guest_api_port")
	if err != nil {
		fmt.Println("启动端口错误", err)
		os.Exit(1)
	}
	common.Logger.Info("Start API!")
	app.Run(fmt.Sprintf(":%d", port))
}
