package admin

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/xiuos/mozi/common"
	"github.com/xiuos/mozi/models/lotto"
	"github.com/xiuos/mozi/routes"
)

func PageFindCodeLottoList(c *gin.Context) {
	params := routes.ParamHelper{}
	params.GetQuery(c, "draw")
	params.GetQuery(c, "curr_page")
	params.GetQuery(c, "page_row")

	params.GetQuery(c, "lotto_id")
	params.GetQuery(c, "name")
	params.GetQuery(c, "lotto_type")
	params.GetQuery(c, "game_kind")
	params.GetQuery(c, "game_type")
	params.GetQuery(c, "status")
	params.GetQuery(c, "sort_index")
	params.GetQuery(c, "introduction")

	currPage := common.GetInt(params.Get("curr_page"))
	pageRow := common.GetInt(params.Get("page_row"))

	if currPage < 1 {
		currPage = common.PageDefaultPage
	}

	if pageRow < 1 {
		pageRow = common.PageDefaultRow
	}
	pp := common.PageParams{CurrentPage: currPage, PageRow: pageRow, Params: params}
	pr, _, err := lotto.PageFindCodeLottoList(pp)
	if err != nil {
		c.JSON(200, routes.ApiResult(common.CodeFail, fmt.Sprintf("%s", err), map[string]string{}))
	} else {
		c.JSON(200, routes.ApiResult(common.CodeOK, "", pr))
	}
}

func FindCodeLottoList(c *gin.Context) {
	params := routes.ParamHelper{}

	params.GetQuery(c, "lotto_id")
	params.GetQuery(c, "name")
	params.GetQuery(c, "lotto_type")
	params.GetQuery(c, "game_kind")
	params.GetQuery(c, "game_type")
	params.GetQuery(c, "status")
	params.GetQuery(c, "sort_index")
	params.GetQuery(c, "introduction")

	data, err := lotto.FindCodeLottoList(params)
	if err != nil {
		c.JSON(200, routes.ApiResult(common.CodeFail, fmt.Sprintf("%s", err), map[string]string{}))
	} else {
		c.JSON(200, routes.ApiResult(common.CodeOK, "", data))
	}
}
