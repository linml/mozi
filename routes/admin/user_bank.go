package admin

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/xiuos/mozi/common"
	"github.com/xiuos/mozi/models"
	"github.com/xiuos/mozi/routes"
)

func PageFindUserBank(c *gin.Context) {
	params := routes.ParamHelper{}
	params.GetQuery(c, "draw")
	params.GetQuery(c, "name")
	params.GetQuery(c, "real_name")
	params.GetQuery(c, "card_no")
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
	pr, _, err := models.PageFindUserBank4AdminList(pp)
	if err != nil {
		c.JSON(200, routes.ApiResult(common.CodeFail, fmt.Sprintf("%s", err), map[string]string{}))
	} else {
		c.JSON(200, routes.ApiResult(common.CodeOK, "", pr))
	}
}
