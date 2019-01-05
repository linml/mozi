package api

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/xiuos/mozi/common"
	"github.com/xiuos/mozi/models"
	"github.com/xiuos/mozi/routes"
	"github.com/xiuos/mozi/service"
)

// 游客登录
func GuestLogin(c *gin.Context) {
	gu, err := models.GenGuestUser()
	if err != nil {
		c.JSON(200, routes.ApiResult(common.CodeFail, fmt.Sprintf("%s", err), map[string]string{}))
	} else {

		sid, err := service.SetAccountOnline(gu.UserID, service.PrefixGuestUser, 0)

		if err != nil {
			c.JSON(200, routes.ApiResult(common.CodeFail, fmt.Sprintf("%s", err), map[string]string{}))
		}

		c.SetCookie(common.SID, sid, 3600*2, "/", "", false, true)
		c.SetCookie("GUEST", "1", 3600*2, "/", "", false, true)
		c.JSON(200, routes.ApiResult(common.CodeOK, "", map[string]interface{}{
			"user_id":  gu.UserID,
			"name":     gu.Name,
			common.SID: sid,
		}))
	}
}

// 获取用户余额
func GetGuestBalance(c *gin.Context) {
	uid, err := routes.GetAPILoginID(c)
	balance, err := service.GetGuestBalance(uid)

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

// 获取游客用户基本信息集合
func GetGuestInfos(c *gin.Context) {
	uid, err := routes.GetAPILoginID(c)
	if err != nil {
		c.JSON(200, routes.ApiShowResult(common.CodeFail, "登录校验失败"))
	}
	ui, err := service.GetGuestInfos(uid)
	if err != nil {
		c.JSON(200, routes.ApiResult(common.CodeFail, fmt.Sprintf("%s", err), map[string]interface{}{}))
	} else {
		c.JSON(200, routes.ApiResult(common.CodeOK, "", ui))
	}
}
