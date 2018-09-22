package main

import (
	"github.com/gin-gonic/gin"
	"github.com/xiuos/mozi/common"
	"github.com/xiuos/mozi/routes/admin"
	"github.com/xiuos/xlog"
	"os"
	"github.com/gin-gonic/contrib/sessions"
	"github.com/xiuos/mozi/routes"
)

func init() {

	savePath, err := common.Conf.String("log", "path")
	if err != nil {
		os.Exit(1)
	}
	filePrefix, err := common.Conf.String("log", "admin_file_prefix")
	if err != nil {
		os.Exit(1)
	}
	common.Logger = xlog.NewLogger(savePath, filePrefix, xlog.Ldate|xlog.Ltime|xlog.Lmicroseconds|xlog.Llongfile, xlog.LevelInfo)
}

func main() {
	app := gin.Default()
	store, _ := sessions.NewRedisStore(10, "tcp", "localhost:6379", "", []byte("secret"))
	app.Use(sessions.Sessions("session", store))

	//gin.SetMode(gin.ReleaseMode)
	app.LoadHTMLGlob("templates/**/*")
	app.Static("/static", "./static")


	appV1 := app.Group("/")
	{
		appV1.GET("/captcha", routes.GenCaptcha)
		appV1.GET("/login", admin.TmplLogin)
		appV1.GET("/logout", admin.TmplLogout)
		appV1.POST("/login", admin.Login)
	}

	appAuthV1 := app.Group("/")
	appAuthV1.Use(routes.AdminAuthMiddleWare())
	{
		appAuthV1.GET("/index", admin.TmplIndex)
		appAuthV1.GET("/menu", admin.AdminMenu)
		appAuthV1.GET("/page/role_group", admin.TmplRoleGroup)
		appAuthV1.GET("/page/account/list", admin.TmplAccountList)
	}


	app.Run(":9980")

}
