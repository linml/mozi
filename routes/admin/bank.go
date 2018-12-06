package admin

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/xiuos/mozi/common"
	"github.com/xiuos/mozi/models"
	"github.com/xiuos/mozi/routes"
)

func PageFindCodeBank(c *gin.Context) {
	params := routes.ParamHelper{}
	params.GetQuery(c, "draw")
	params.GetQuery(c, "id")
	params.GetQuery(c, "bank_code")
	params.GetQuery(c, "bank_name")
	params.GetQuery(c, "group")
	params.GetQuery(c, "status")
	params.GetQuery(c, "curr_page")
	params.GetQuery(c, "page_row")
	currPage := common.GetInt(params.Get("curr_page"))
	pageRow := common.GetInt(params.Get("page_row"))

	if currPage < 1 {
		currPage = common.PageDefaultPage
	}

	if pageRow < 1 {
		pageRow = common.PageDefaultRow
	}
	pp := common.PageParams{CurrentPage: currPage, PageRow: pageRow, Params: params}
	pr, _, err := models.PageFindCodeBankList(pp)
	if err != nil {
		c.JSON(200, routes.ApiResult(common.CodeFail, fmt.Sprintf("%s", err), map[string]string{}))
	} else {
		c.JSON(200, routes.ApiResult(common.CodeOK, "", pr))
	}
}
