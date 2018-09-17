package admin

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"github.com/xiuos/mozi/routes"
	"github.com/xiuos/mozi/common"
	"fmt"
	"github.com/xiuos/mozi/service"
	"github.com/gin-gonic/contrib/sessions"
)

func TmplLogin(c *gin.Context) {
	c.HTML(http.StatusOK, "login.html", gin.H{"title": "登录"})
}

func Login(c *gin.Context) {
	params := routes.ParamHelper{}

	name := c.PostForm("name")
	password := c.PostForm("password")

	params.Set("ip", c.ClientIP())
	params.Set("user_agent", c.Request.Header.Get("User-Agent"))

	scheme := "http://"
	if c.Request.TLS != nil {
		scheme = "https://"
	}
	params.Set("url", scheme+c.Request.Host+c.Request.RequestURI)
	err := service.AuthAdminLogin(name, password, params)
	if err != nil {
		c.JSON(200, routes.ApiResult(common.CodeFail, fmt.Sprintf("%s", err), map[string]string{}))
	} else {
		uid, err := service.GetAdminUserIDByName(name)
		if err == nil {
			session := sessions.Default(c)
			session.Set(routes.SessionAdminLoginID, uid)
			session.Save()
		}

		c.JSON(200, routes.ApiResult(common.CodeOK, "", map[string]interface{}{
			"user_id": uid,
			"name":    name,
		}))
	}
}


func TmplIndex(c *gin.Context) {
	c.HTML(http.StatusOK, "index.html", gin.H{"title": "管理后台"})
}

func Render(c *gin.Context)  {
	c.HTML(http.StatusOK, "login.html", gin.H{"title": "管理后台"})
}