package routes

import (
	"github.com/afocus/captcha"
	"github.com/gin-gonic/contrib/sessions"
	"github.com/gin-gonic/gin"
	"github.com/xiuos/mozi/common"
	"image/png"
	"strings"
)

func GenCaptcha(c *gin.Context) {
	img, str := common.Captcha.Create(4, captcha.ALL)
	session := sessions.Default(c)
	session.Set(SessionCaptcha, strings.ToLower(str))
	session.Save()
	png.Encode(c.Writer, img)
}

func CheckCaptcha(c *gin.Context, cap string) bool {
	session := sessions.Default(c)
	str := session.Get(SessionCaptcha)
	if str == nil {
		return false
	}

	if str != strings.ToLower(cap) {
		return false
	}
	return true
}
