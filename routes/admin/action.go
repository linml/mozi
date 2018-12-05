package admin

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/xiuos/mozi/common"
	"github.com/xiuos/mozi/models"
	"github.com/xiuos/mozi/routes"
)

func FindCodeActionAdminModuleList(c *gin.Context) {
	_, err := routes.GetAdminLoginID(c)
	if err != nil {
		c.JSON(200, routes.ApiResult(common.CodeFail, fmt.Sprintf("%s", err), map[string]string{}))
		return
	}
	dataList, err := models.FindCodeActionAdminModule(map[string]string{"status": "1"})

	if err != nil {
		c.JSON(200, routes.ApiResult(common.CodeFail, fmt.Sprintf("%s", err), map[string]string{}))
	} else {
		c.JSON(200, routes.ApiResult(common.CodeOK, "", dataList))
	}
}

func FindCodeActionAdminList(c *gin.Context) {
	_, err := routes.GetAdminLoginID(c)
	if err != nil {
		c.JSON(200, routes.ApiResult(common.CodeFail, fmt.Sprintf("%s", err), map[string]string{}))
		return
	}
	dataList, err := models.FindCodeActionAdmin(map[string]string{"status": "1"})

	if err != nil {
		c.JSON(200, routes.ApiResult(common.CodeFail, fmt.Sprintf("%s", err), map[string]string{}))
	} else {
		c.JSON(200, routes.ApiResult(common.CodeOK, "", dataList))
	}
}

func PageFindLogRecordAdminActionList(c *gin.Context) {
	params := routes.ParamHelper{}
	params.GetQuery(c, "draw")
	params.GetQuery(c, "user_id")
	params.GetQuery(c, "name")
	params.GetQuery(c, "action_module")
	params.GetQuery(c, "action_id")
	params.GetQuery(c, "ip")
	params.GetQuery(c, "success")
	params.GetQuery(c, "operator_id")
	params.GetQuery(c, "operator_name")
	params.GetQuery(c, "operator_type")
	params.GetQuery(c, "start_at")
	params.GetQuery(c, "end_at")
	params.GetQuery(c, "curr_page")
	params.GetQuery(c, "page_row")
	currPage := common.GetInt(params.Get("curr_page"))
	pageRow := common.GetInt(params.Get("page_row"))

	if currPage < 1 {
		currPage = models.PageDefaultPage
	}

	if pageRow < 1 {
		pageRow = models.PageDefaultRow
	}
	pp := models.PageParams{CurrentPage: currPage, PageRow: pageRow, Params: params}
	pr, _, err := models.PageFindLogRecordAdminActionList(pp)
	if err != nil {
		c.JSON(200, routes.ApiResult(common.CodeFail, fmt.Sprintf("%s", err), map[string]string{}))
	} else {
		c.JSON(200, routes.ApiResult(common.CodeOK, "", pr))
	}
}
