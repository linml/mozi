package admin

import (
	"fmt"
	"github.com/gin-gonic/contrib/sessions"
	"github.com/gin-gonic/gin"
	"github.com/xiuos/mozi/common"
	"github.com/xiuos/mozi/routes"
	"github.com/xiuos/mozi/service"
)

func Login(c *gin.Context) {
	params := routes.ParamHelper{}

	name := c.PostForm("name")
	password := c.PostForm("password")
	code := c.PostForm("code")
	if b := routes.CheckCaptcha(c, code); b == false {
		c.JSON(200, routes.ApiResult(common.CodeFail, "验证码错误", map[string]string{}))
		return
	}

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
