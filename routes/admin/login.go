package admin

import (
	"fmt"
	"github.com/gin-gonic/contrib/sessions"
	"github.com/gin-gonic/gin"
	"github.com/xiuos/mozi/common"
	"github.com/xiuos/mozi/models"
	"github.com/xiuos/mozi/routes"
	"github.com/xiuos/mozi/service"
	"net/http"
)

func TmplLogin(c *gin.Context) {
	if b := routes.CheckIsLogin(c); b == true {
		c.Redirect(http.StatusFound, "/index")
	}
	c.HTML(http.StatusOK, "login.html", gin.H{"title": "登录"})
}

func TmplLogout(c *gin.Context) {
	session := sessions.Default(c)
	session.Delete(routes.SessionAdminLoginID)
	session.Save()
	c.Redirect(http.StatusFound, "/login")
}

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

func TmplRoot(c *gin.Context) {
	c.Redirect(http.StatusFound, "/login")
}

func TmplIndex(c *gin.Context) {
	uid, _ := routes.GetAdminLoginID(c)
	u, _ := models.GetAdminUserByID(uid)
	c.HTML(http.StatusOK, "index.html", gin.H{"title": "管理后台", "admin_user": u.Name})
}

func TmplAdminRole(c *gin.Context) {
	c.HTML(http.StatusOK, "admin_role.html", gin.H{"title": "管理后台"})
}

func TmplAdminList(c *gin.Context) {
	c.HTML(http.StatusOK, "admin_list.html", gin.H{"title": "管理后台"})
}

func TmplAdminRoleMenu(c *gin.Context) {
	c.HTML(http.StatusOK, "admin_role_menu.html", gin.H{"title": "管理后台"})
}

func TmplMemberList(c *gin.Context) {
	c.HTML(http.StatusOK, "user_list.html", gin.H{})
}

func TmplRecordLogin(c *gin.Context) {
	c.HTML(http.StatusOK, "record_login.html", gin.H{})
}

func TmplUserBank(c *gin.Context) {
	c.HTML(http.StatusOK, "user_bank.html", gin.H{})
}

func TmplcodeBank(c *gin.Context) {
	c.HTML(http.StatusOK, "code_bank.html", gin.H{})
}

func TmplRecordAdminAction(c *gin.Context) {
	c.HTML(http.StatusOK, "record_admin_action.html", gin.H{})
}

func TmplLottoCodeLotto(c *gin.Context) {
	c.HTML(http.StatusOK, "code_lotto.html", gin.H{})
}

func TmplLottoResultotto(c *gin.Context) {
	c.HTML(http.StatusOK, "lotto_result.html", gin.H{})
}

func TmplLottoOdds(c *gin.Context) {
	c.HTML(http.StatusOK, "lotto_odds.html", gin.H{})
}
