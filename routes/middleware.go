package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/xiuos/mozi/common"
	"github.com/xiuos/mozi/models/errors"
	"github.com/xiuos/mozi/service"
	"net/http"
)

const (
	SessionApiLoginID   = "__session_api_login_id__"
	SessionAdminLoginID = "__session_admin_login_id__"
	SessionCaptcha      = "__captcha__"
)

func GetAPILoginID(c *gin.Context) (int, error) {
	uid, b := c.Get(SessionApiLoginID)
	if false == b {
		return -1, errors.Unauthorized{}
	}
	return uid.(int), nil
}

func GetAdminLoginID(c *gin.Context) (int, error) {
	uid, b := c.Get(SessionAdminLoginID)
	if false == b {
		return -1, errors.Unauthorized{}
	}
	return uid.(int), nil
}

func APIAuthMiddleWare() gin.HandlerFunc {
	return func(c *gin.Context) {
		sid, err := c.Cookie(common.SID)
		if err != nil {
			c.JSON(http.StatusUnauthorized, gin.H{
				"code": common.CodeUserNotLogin,
				"msg":  "请先登录",
				"data": gin.H{}})
			c.Abort()
			return
		}

		obj, err := service.GetUserSessionInfo(sid)
		if err != nil {
			c.JSON(http.StatusUnauthorized, gin.H{
				"code": common.CodeUserLoginExpired,
				"msg":  "请重新登录",
				"data": gin.H{}})
			c.Abort()
			return
		}
		c.Set(SessionApiLoginID, obj.UserID)
		c.Next()
		return

	}
}

func AdminAuthMiddleWare() gin.HandlerFunc {
	return func(c *gin.Context) {
		sid, err := c.Cookie(common.SID)
		if err != nil {
			c.JSON(http.StatusUnauthorized, gin.H{
				"code": common.CodeUserNotLogin,
				"msg":  "请先登录",
				"data": gin.H{}})
			c.Abort()
			return
		}

		obj, err := service.GetAdminSessionInfo(sid)
		if err != nil {
			c.JSON(http.StatusUnauthorized, gin.H{
				"code": common.CodeUserLoginExpired,
				"msg":  "您的账号已在别处登录,请重新登录",
				"data": gin.H{}})
			c.Abort()
			return
		}
		c.Set(SessionAdminLoginID, obj.UserID)
		c.Next()
		return

	}
}

func CheckIsLogin(c *gin.Context) bool {
	sid, err := c.Cookie(common.SID)
	if err != nil {
		return false
	}

	_, err = service.GetAdminSessionInfo(sid)
	if err != nil {
		return false
	}
	return true
}
