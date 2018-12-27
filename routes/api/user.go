package api

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/xiuos/mozi/common"
	"github.com/xiuos/mozi/models"
	"github.com/xiuos/mozi/models/errors"
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

	err := service.RegisterUser(name, password, params, models.OperatorTypeSelf)
	if err != nil {
		c.JSON(200, routes.ApiShowResult(common.CodeFail, fmt.Sprintf("%s", err)))
	} else {
		uid, _ := service.GetUserIDByName(name)
		r := models.RecordUserAction{
			UserID:       uid,
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
		uid, err := service.GetUserIDByName(name)
		sessionID := common.RandString(32)
		if err == nil {
			err = service.SetUserOnline(sessionID, uid, 0)
			if err == nil {
				c.SetCookie(common.SID, sessionID, 0, "/", "", false, true)
			}
		}

		c.JSON(200, routes.ApiResult(common.CodeOK, "", map[string]interface{}{
			"user_id":  uid,
			"name":     name,
			common.SID: sessionID,
		}))
	}

}

// 获取用户余额
func GetBalance(c *gin.Context) {
	uid, err := routes.GetAPILoginID(c)
	balance, err := service.GetBalance(uid)

	if err != nil {
		c.JSON(200, routes.ApiResult(common.CodeFail, fmt.Sprintf("%s", err), map[string]interface{}{
			"balance": balance,
		}))
	} else {
		c.JSON(200, routes.ApiResult(common.CodeOK, "", map[string]interface{}{
			"balance": balance,
		}))
	}
}

// 获取用户基本信息集合
func GetInfos(c *gin.Context) {
	uid, err := routes.GetAPILoginID(c)
	if err != nil {
		c.JSON(200, routes.ApiShowResult(common.CodeFail, "登录校验失败"))
	}
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
	if err != nil {
		c.JSON(200, routes.ApiShowResult(common.CodeFail, "登录校验失败"))
	}
	oldPassword := c.PostForm("old_password")
	newPassword := c.PostForm("new_password")

	err = service.ResetUserPassword(uid, oldPassword, newPassword, models.OperatorTypeSelf)
	if err != nil {
		c.JSON(200, routes.ApiShowResult(common.CodeFail, fmt.Sprintf("%s", err)))
	} else {
		c.JSON(200, routes.ApiShowResult(common.CodeOK, ""))
	}
}

// 用户 profile
func UserProfile(c *gin.Context) {
	uid, err := routes.GetAPILoginID(c)
	if err != nil {
		c.JSON(200, routes.ApiShowResult(common.CodeFail, "登录校验失败"))
	}
	up, err := models.GetUserProfile(uid)

	if err != nil {
		c.JSON(200, routes.ApiResult(common.CodeFail, fmt.Sprintf("%s", err), map[string]string{}))
	} else {
		c.JSON(200, routes.ApiResult(common.CodeOK, "", up))
	}
}

// 设置真实姓名
// 参数：real_name
func SetProfileRealName(c *gin.Context) {
	uid, err := routes.GetAPILoginID(c)
	if err != nil {
		c.JSON(200, routes.ApiShowResult(common.CodeFail, "登录校验失败"))
	}
	realName := c.PostForm("real_name")
	up, err := models.GetUserProfile(uid)
	if up.RealName != "" {
		err = errors.RealNameExistErr{}
		c.JSON(200, routes.ApiShowResult(common.CodeFail, fmt.Sprintf("%s", err)))
	} else {
		err = service.SetUserProfileRealName(uid, realName)
		if err != nil {
			c.JSON(200, routes.ApiShowResult(common.CodeFail, fmt.Sprintf("%s", err)))
		} else {
			c.JSON(200, routes.ApiShowResult(common.CodeOK, ""))
		}
	}
}

// 设置昵称
// 参数：nickname
func SetProfileNicname(c *gin.Context) {
	uid, err := routes.GetAPILoginID(c)
	if err != nil {
		c.JSON(200, routes.ApiShowResult(common.CodeFail, "登录校验失败"))
	}
	nickname := c.PostForm("nickname")

	err = service.SetUserProfileNickname(uid, nickname)
	if err != nil {
		c.JSON(200, routes.ApiShowResult(common.CodeFail, fmt.Sprintf("%s", err)))
	} else {
		c.JSON(200, routes.ApiShowResult(common.CodeOK, ""))
	}

}

// 初始化资金密码

func InitUserWalletPassword(c *gin.Context) {
	// 参数： password，范围市6位数字
	uid, err := routes.GetAPILoginID(c)
	if err != nil {
		c.JSON(200, routes.ApiShowResult(common.CodeFail, "登录校验失败"))
	}

	password := c.PostForm("password")

	err = service.InitUserWalletPassword(uid, password)
	if err != nil {
		c.JSON(200, routes.ApiShowResult(common.CodeFail, fmt.Sprintf("%s", err)))
	} else {
		c.JSON(200, routes.ApiShowResult(common.CodeOK, ""))
	}
}

// 重置资金密码
func ResetUserWalletPassword(c *gin.Context) {
	uid, err := routes.GetAPILoginID(c)
	if err != nil {
		c.JSON(200, routes.ApiShowResult(common.CodeFail, "登录校验失败"))
	}

	oldPassword := c.PostForm("old_password")
	newPassword := c.PostForm("new_password")

	err = service.RestUserWalletPassword(uid, oldPassword, newPassword, models.OperatorTypeSelf)
	if err != nil {
		c.JSON(200, routes.ApiShowResult(common.CodeFail, fmt.Sprintf("%s", err)))
	} else {
		c.JSON(200, routes.ApiShowResult(common.CodeOK, ""))
	}
}

// 获取资金密码状态
func UserWalletPasswordStatus(c *gin.Context) {
	uid, err := routes.GetAPILoginID(c)
	if err != nil {
		c.JSON(200, routes.ApiShowResult(common.CodeFail, "登录校验失败"))
	}

	status, err := service.GetUserWalletPasswordStatus(uid)
	if err != nil {
		c.JSON(200, routes.ApiResult(common.CodeFail, fmt.Sprintf("%s", err), map[string]string{}))
	} else {
		c.JSON(200, routes.ApiResult(common.CodeOK, "", map[string]interface{}{
			"status": status,
		}))
	}
}

// 生成推广链接
func UserGenLink(c *gin.Context) {
	uid, err := routes.GetAPILoginID(c)
	if err != nil {
		c.JSON(200, routes.ApiShowResult(common.CodeFail, "登录校验失败"))
	}
	uType := common.GetInt(c.PostForm("user_type"))

	refInfo := models.UserLinks{UserID: uid, UserType: uType}

	err = service.GenUserRef(refInfo)
	if err != nil {
		c.JSON(200, routes.ApiShowResult(common.CodeFail, fmt.Sprintf("%s", err)))
	} else {
		r := models.RecordUserAction{
			UserID:       uid,
			ActionModule: models.ActionModuleUser,
			ActionID:     models.ActionCreateRef,
			Content:      "",
			IP:           c.ClientIP(),
			RecordAt:     common.GetTimeNowString(),
			Success:      common.CodeOK,
			Message:      "操作成功",
			OperatorType: models.OperatorTypeSelf,
		}
		models.LogRecordUserAction(&r)
		c.JSON(200, routes.ApiShowResult(common.CodeOK, ""))
	}
}
