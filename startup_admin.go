package main

import (
	"github.com/gin-gonic/contrib/sessions"
	"github.com/gin-gonic/gin"
	"github.com/xiuos/mozi/common"
	"github.com/xiuos/mozi/routes"
	"github.com/xiuos/mozi/routes/admin"
	"github.com/xiuos/mozi/xlog"
	"os"
)

func init() {

	savePath, err := common.Conf.String("log", "path")
	if err != nil {
		os.Exit(1)
	}
	filePrefix, err := common.Conf.String("log", "admin_file_prefix")
	if err != nil {
		os.Exit(1)
	}
	common.Logger = xlog.NewLogger(savePath, filePrefix, xlog.Ldate|xlog.Ltime|xlog.Lmicroseconds|xlog.Llongfile, xlog.LevelInfo)
}

func main() {
	app := gin.Default()
	store, _ := sessions.NewRedisStore(10, "tcp", "localhost:6379", "", []byte("secret"))
	app.Use(sessions.Sessions("session", store))

	//gin.SetMode(gin.ReleaseMode)
	app.LoadHTMLGlob("templates/**/*")
	app.Static("/static", "./static")

	appV1 := app.Group("/")
	{
		appV1.GET("/", admin.TmplRoot)
		appV1.GET("/captcha", routes.GenCaptcha)
		appV1.GET("/login", admin.TmplLogin)
		appV1.GET("/logout", admin.TmplLogout)
		appV1.POST("/login", admin.Login)
	}

	appAuthV1 := app.Group("/")
	appAuthV1.Use(routes.AdminAuthMiddleWare())
	{
		appAuthV1.GET("/index", admin.TmplIndex)
		appAuthV1.GET("/menu", admin.AdminMenu)
		appAuthV1.GET("/html/admin_role", admin.TmplAdminRole)
		appAuthV1.GET("/html/admin_list", admin.TmplAdminList)
		appAuthV1.GET("/html/admin_role_menu", admin.TmplAdminRoleMenu)
		appAuthV1.GET("/html/member/list", admin.TmplMemberList)
		appAuthV1.GET("/html/member/record_login", admin.TmplRecordLogin)
		appAuthV1.GET("/html/member/user_bank", admin.TmplUserBank)
		appAuthV1.GET("/html/bank/code_bank", admin.TmplcodeBank)
		appAuthV1.GET("/html/record/admin_action", admin.TmplRecordAdminAction)
		appAuthV1.GET("/html/lotto/code_lotto", admin.TmplLottoCodeLotto)
		appAuthV1.GET("/html/lotto/result_lotto", admin.TmplLottoResultotto)
		appAuthV1.GET("/html/lotto/odds", admin.TmplLottoOdds)
		appAuthV1.GET("/html/lotto/order", admin.TmplLottoOrder)
		appAuthV1.GET("/html/lotto/report", admin.TmplLottoReport)
		appAuthV1.GET("/html/user/money_change", admin.TmplUserMoneyChange)

		appAuthV1.GET("/pages/user/list", admin.PageFindUserList)
		appAuthV1.GET("/pages/user/money_change/list", admin.PageFindRecordMoneyList)
		appAuthV1.GET("/pages/admin/list", admin.PageFindAdminList)
		appAuthV1.GET("/pages/user/record_login", admin.PageFindUserRecordLogin)
		appAuthV1.GET("/pages/user/bank/list", admin.PageFindUserBank)
		appAuthV1.GET("/pages/code_bank/list", admin.PageFindCodeBank)
		appAuthV1.GET("/pages/code_lotto/list", admin.PageFindCodeLottoList)
		appAuthV1.GET("/pages/lotto/result/list", admin.PageFindLottoResultList)
		appAuthV1.GET("/pages/lotto/order/list", admin.PageFindLottOrderList)
		appAuthV1.GET("/lotto/code_lotto", admin.GetCodeLotto)
		appAuthV1.POST("/lotto/code_lotto/status", admin.SetCodeLottoStatus)
		appAuthV1.POST("/lotto/code_lotto/is_show", admin.SetCodeLottoIsShow)
		appAuthV1.POST("/lotto/code_lotto/sort_index", admin.SetCodeLottoSortIndex)
		appAuthV1.GET("/lotto/code_lotto/list", admin.FindCodeLottoList)
		appAuthV1.GET("/lotto/code_lotto_type/list", admin.FindCodeLottoTypeList)
		appAuthV1.GET("/lotto/method_template/list", admin.FindLottoMethodTemplateList)
		appAuthV1.GET("/money/code_change_kind/list", admin.FindCodeChangeMoneyKindList)
		appAuthV1.GET("/money/code_change_type/list", admin.FindCodeChangeMoneyTypeList)
		appAuthV1.GET("/lotto/odds/list", admin.FindLottOddsList)
		appAuthV1.GET("/pages/lotto/odds/list", admin.PageFindLottOddsList)
		appAuthV1.GET("/pages/lotto/count_report/list", admin.PageFindAgentLottoReportList4Admin)
		appAuthV1.GET("/pages/lotto/report_day_count/list", admin.PageFindReportLottoDayCountList)
		appAuthV1.GET("/lotto/odds", admin.GetLottOdds)

		appAuthV1.POST("/lotto/odds/set", admin.SetLottOddsInfo)

		appAuthV1.GET("/member/infos", admin.GetMemberInfos)
		appAuthV1.POST("/member/add", admin.AddMember)
		appAuthV1.GET("/member/profile", admin.GetMemberProfile)
		appAuthV1.GET("/member/relation", admin.GetMemberRelation)
		appAuthV1.GET("/member/power", admin.GetMemberPower)
		appAuthV1.GET("/menu/node/list", admin.FindAdminMenuNode)
		appAuthV1.GET("/role/list", admin.FindAdminRole)
		appAuthV1.GET("/role_menu/list", admin.FindAdminRoleMenu)
		appAuthV1.POST("/role/add", admin.AddAdminRole)
		appAuthV1.POST("/role_menu/set", admin.UpdateRoleMenu)
		appAuthV1.POST("/admin/role/set", admin.SetAdminRole)
		appAuthV1.POST("/admin/status/set", admin.SetAdminStatus)
		appAuthV1.POST("/admin/add", admin.CreateAdmin)

		appAuthV1.GET("/action/action_module/list", admin.FindCodeActionAdminModuleList)
		appAuthV1.GET("/action/action_id/list", admin.FindCodeActionAdminList)
		appAuthV1.GET("/action/action_id/map", admin.CreateAdmin)
		appAuthV1.GET("/pages/action/record/list", admin.PageFindLogRecordAdminActionList)

	}
	app.Run(":9980")
}
