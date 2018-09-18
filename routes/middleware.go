package routes

import (
	"fmt"
	"github.com/gin-gonic/contrib/sessions"
	"github.com/gin-gonic/gin"
	"github.com/xiuos/mozi/common"
	"github.com/xiuos/mozi/models/errors"
	"net/http"
)

const (
	SessionApiLoginID   = "__session_api_login_id__"
	SessionAdminLoginID = "__session_admin_login_id__"
	SessionCaptcha = "__captcha__"
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
		session := sessions.Default(c)
		uid := session.Get(SessionApiLoginID)
		if uid == nil {
			c.JSON(http.StatusUnauthorized, gin.H{
				"code": common.CodeUserNotLogin,
				"msg":  "请先登录",
				"data": gin.H{}})
			c.Abort()
			return
		}
		c.Set(SessionApiLoginID, uid.(int))
		c.Next()
		return

	}
}

func AdminAuthMiddleWare() gin.HandlerFunc {
	return func(c *gin.Context) {
		session := sessions.Default(c)
		uid := session.Get(SessionAdminLoginID)
		if uid == nil {
			c.JSON(http.StatusUnauthorized, gin.H{
				"code": common.CodeUserNotLogin,
				"msg":  "请先登录",
				"data": gin.H{}})
			c.Abort()
			return
		}
		c.Set(SessionAdminLoginID, uid.(int))
		c.Next()
		return

	}
}

func CheckIsLogin(c *gin.Context) bool {
	session := sessions.Default(c)
	uid := session.Get(SessionAdminLoginID)
	fmt.Println(uid)
	if uid == nil {
		return false
	}
	return true
}
