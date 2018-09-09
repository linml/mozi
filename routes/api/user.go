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

// 用户注册接口。
// 参数 name: 用户名, password: 密码, ref: 推广码(非必须)
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

// 登录接口。
// 参数 name: 用户名, password: 密码
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
			session.Set(routes.SessionApiLoginID, u.UserID)
			session.Save()
		}

		c.JSON(200, routes.ApiResult(common.CodeOK, "", map[string]interface{}{
			"user_id": u.UserID,
			"name":    u.Name,
		}))
	}

}

// 获取用户余额
func GetBalance(c *gin.Context) {
	uid, err := routes.GetAPILoginID(c)
	w, err := models.GetUserWallet(uid)

	if err != nil {
		c.JSON(200, routes.ApiResult(common.CodeFail, fmt.Sprintf("%s", err), map[string]interface{}{
			"balance": w.Balance,
		}))
	} else {
		c.JSON(200, routes.ApiResult(common.CodeOK, "", map[string]interface{}{
			"balance": w.Balance,
		}))
	}
}

// 获取用户基本信息集合
func GetInfos(c *gin.Context) {
	uid, err := routes.GetAPILoginID(c)
	ui, err := service.GetInfos(uid)
	if err != nil {
		c.JSON(200, routes.ApiResult(common.CodeFail, fmt.Sprintf("%s", err), map[string]interface{}{}))
	} else {
		c.JSON(200, routes.ApiResult(common.CodeOK, "", ui))
	}
}

// 重置密码
func ResetPassword(c *gin.Context) {
	uid, err := routes.GetAPILoginID(c)
	oldPassword := c.PostForm("old_password")
	newPassword := c.PostForm("new_password")

	err = service.ResetUserPassword(uid, oldPassword, newPassword, models.OperatorTypeSelf)
	if err != nil {
		c.JSON(200, routes.ApiShowResult(common.CodeFail, fmt.Sprintf("%s", err)))
	} else {
		c.JSON(200, routes.ApiShowResult(common.CodeOK, ""))
	}
}
