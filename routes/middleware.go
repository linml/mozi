package routes

import (
	"github.com/gin-gonic/contrib/sessions"
	"github.com/gin-gonic/gin"
	"github.com/xiuos/mozi/common"
	"github.com/xiuos/mozi/models/errors"
	"net/http"
)

const (
	SessionApiLoginID = "__session_api_login_id__"
)

func GetAPILoginID(c *gin.Context) (int, error) {
	uid, b := c.Get(SessionApiLoginID)
	if false == b {
		return -1, errors.Unauthorized{}
	}
	return uid.(int), nil
}

func APIAuthMiddleWare() gin.HandlerFunc {
	return func(c *gin.Context) {
		session := sessions.Default(c)
		session.Get("login_user_id")
		uid := session.Get("login_user_id")
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
