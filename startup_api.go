package main

import (
	"fmt"
	"github.com/gin-gonic/contrib/sessions"
	"github.com/gin-gonic/gin"
	"github.com/xiuos/mozi/common"
	"github.com/xiuos/mozi/routes"
	"github.com/xiuos/mozi/routes/api"
	"github.com/xiuos/xlog"
	"os"
	"runtime"
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
	apiAuthV1.Use(routes.APIAuthMiddleWare())
	{
		apiAuthV1.GET("/get_balance", api.GetBalance)
		apiAuthV1.GET("/infos", api.GetInfos)
		apiAuthV1.POST("/password/reset", api.ResetPassword)
		apiAuthV1.GET("/profile", api.UserProfile)
		apiAuthV1.POST("/profile/real_name", api.SetProfileRealName)
		apiAuthV1.POST("/profile/nickname", api.SetProfileNicname)
		apiAuthV1.POST("/wallet/password/init", api.InitUserWalletPassword)
		apiAuthV1.POST("/wallet/password/reset", api.ResetUserWalletPassword)
		apiAuthV1.GET("/wallet/password/status", api.UserWalletPasswordStatus)
	}

	port, err := common.Conf.Int("port", "api_port")
	if err != nil {
		os.Exit(1)
	}
	common.Logger.Info("Start API!")
	app.Run(fmt.Sprintf(":%d", port))
}
