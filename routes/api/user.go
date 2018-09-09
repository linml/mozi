package api

import (
	"fmt"
	"github.com/gin-gonic/contrib/sessions"
	"github.com/gin-gonic/gin"
	"github.com/xiuos/mozi/common"
	"github.com/xiuos/mozi/models"
	"github.com/xiuos/mozi/routes"
	"github.com/xiuos/mozi/service"
)

func Register(c *gin.Context) {
	params := routes.ParamHelper{}

	name := c.PostForm("name")
	password := c.PostForm("password")

	params.GetPostForm(c, "ref")
	params.Set("register_ip", c.ClientIP())

	err := service.RegisterUser(name, password, params)
	if err != nil {
		c.JSON(200, routes.ApiShowResult(common.CodeFail, fmt.Sprintf("%s", err)))
	} else {
		u, _ := models.GetUserByName(name)
		r := models.RecordUserAction{UserID: u.UserID,
			ActionModule: models.ActionModuleUser,
			ActionID:     models.ActionUserRegister,
			Content:      "",
			IP:           c.ClientIP(),
			RecordAt:     common.GetTimeNowString(),
			Success:      common.CodeOK,
			Message:      "注册成功",
			OperatorType: models.OperatorTypeSelf,
		}
		models.LogRecordUserAction(&r)

		c.JSON(200, routes.ApiShowResult(common.CodeOK, ""))
	}
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

	err := service.AuthLogin(name, password, params)
	if err != nil {
		c.JSON(200, routes.ApiResult(common.CodeFail, fmt.Sprintf("%s", err), map[string]string{}))
	} else {
		u, err := models.GetUserByName(name)
		if err == nil {
			session := sessions.Default(c)
			session.Set("login_user_id", u.UserID)
			session.Set("login_name", u.Name)
			session.Save()
		}

		c.JSON(200, routes.ApiResult(common.CodeOK, "", map[string]interface{}{
			"user_id": u.UserID,
			"name":    u.Name,
		}))
	}

}
