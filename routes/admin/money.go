package admin

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/xiuos/mozi/common"
	"github.com/xiuos/mozi/models"
	"github.com/xiuos/mozi/routes"
)

func PageFindRecordMoneyList(c *gin.Context) {
	params := routes.ParamHelper{}
	params.GetQuery(c, "draw")
	params.GetQuery(c, "curr_page")
	params.GetQuery(c, "page_row")

	params.GetQuery(c, "id")
	params.GetQuery(c, "link_id")
	params.GetQuery(c, "user_id")
	params.GetQuery(c, "name")
	params.GetQuery(c, "game_kind")
	params.GetQuery(c, "game_type")
	params.GetQuery(c, "change_kind")
	params.GetQuery(c, "change_type")
	params.GetQuery(c, "before_balance")
	params.GetQuery(c, "amount")
	params.GetQuery(c, "after_balance")
	params.GetQuery(c, "record_at")
	params.GetQuery(c, "operate_type")
	params.GetQuery(c, "operator_id")
	params.GetQuery(c, "operator_name")
	params.GetQuery(c, "amount_min")
	params.GetQuery(c, "amount_max")
	params.GetQuery(c, "record_date_from")
	params.GetQuery(c, "record_date_to")
	params.GetQuery(c, "record_at_from")
	params.GetQuery(c, "record_at_to")

	currPage := common.GetInt(params.Get("curr_page"))
	pageRow := common.GetInt(params.Get("page_row"))

	if currPage < 1 {
		currPage = common.PageDefaultPage
	}

	if pageRow < 1 {
		pageRow = common.PageDefaultRow
	}

	p := common.PageParams{CurrentPage: currPage, PageRow: pageRow, Params: params}

	data, _, err := models.PageFindRecordMoneyList(p)
	if err != nil {
		c.JSON(200, routes.ApiResult(common.CodeFail, fmt.Sprintf("%s", err), map[string]string{}))
	} else {
		c.JSON(200, routes.ApiResult(common.CodeOK, "", data))
	}
}
