package admin

import (
	"github.com/gin-gonic/gin"
	"github.com/xiuos/mozi/common"
	"github.com/xiuos/mozi/models"
	"github.com/xiuos/mozi/routes"
	"net/http"
)

func TmplLogin(c *gin.Context) {
	if b := routes.CheckIsLogin(c); b == true {
		c.Redirect(http.StatusFound, "/index")
	}
	c.HTML(http.StatusOK, "login.html", gin.H{"title": "登录"})
}

func TmplLogout(c *gin.Context) {
	c.SetCookie(common.SID, "", -1, "/", "", false, true)
	c.Redirect(http.StatusFound, "/login")
}

func TmplRoot(c *gin.Context) {
	c.Redirect(http.StatusFound, "/login")
}

func TmplIndex(c *gin.Context) {
	uid, _ := routes.GetAdminLoginID(c)
	u, _ := models.GetAdminUserByID(uid)
	ip := c.ClientIP()
	c.HTML(http.StatusOK, "index.html", gin.H{"title": "管理后台", "admin_user": u.Name, "cur_ip": ip})
}

func TmplDashboard(c *gin.Context) {
	c.HTML(http.StatusOK, "dashboard.html", gin.H{"title": "管理后台"})
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

func TmplRecordAdminLogin(c *gin.Context) {
	c.HTML(http.StatusOK, "record_admin_login.html", gin.H{})
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

func TmplLottoOrder(c *gin.Context) {
	c.HTML(http.StatusOK, "record_lotto_order.html", gin.H{})
}

func TmplUserMoneyChange(c *gin.Context) {
	c.HTML(http.StatusOK, "record_money_change.html", gin.H{})
}

func TmplLottoReport(c *gin.Context) {
	c.HTML(http.StatusOK, "lotto_report.html", gin.H{})
}

func TmplRollNotice(c *gin.Context) {
	c.HTML(http.StatusOK, "roll_notice.html", gin.H{})
}

func TmplNormalNotice(c *gin.Context) {
	c.HTML(http.StatusOK, "normal_notice.html", gin.H{})
}
