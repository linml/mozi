package admin

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/shopspring/decimal"
	"github.com/xiuos/mozi/common"
	"github.com/xiuos/mozi/models"
	"github.com/xiuos/mozi/routes"
	"github.com/xiuos/mozi/service"
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

func FindCodeChangeMoneyKindList(c *gin.Context) {
	params := routes.ParamHelper{}

	params.GetQuery(c, "id")

	data, err := models.FindCodeChangeMoneyKindList(params)
	if err != nil {
		c.JSON(200, routes.ApiResult(common.CodeFail, fmt.Sprintf("%s", err), map[string]string{}))
	} else {
		c.JSON(200, routes.ApiResult(common.CodeOK, "", data))
	}
}

func FindCodeChangeMoneyTypeList(c *gin.Context) {
	params := routes.ParamHelper{}

	params.GetQuery(c, "change_kind")
	params.GetQuery(c, "id")

	data, err := models.FindCodeChangeMoneyTypeList(params)
	if err != nil {
		c.JSON(200, routes.ApiResult(common.CodeFail, fmt.Sprintf("%s", err), map[string]string{}))
	} else {
		c.JSON(200, routes.ApiResult(common.CodeOK, "", data))
	}
}

func ManualAddMoney(c *gin.Context) {
	uid, err := routes.GetAdminLoginID(c)
	if err != nil {
		c.JSON(200, routes.ApiShowResult(common.CodeFail, fmt.Sprintf("%s", err)))
		return
	}
	params := routes.ParamHelper{}

	params.GetPostFormNotEmpty(c, "user_id")
	params.GetPostFormNotEmpty(c, "amount")
	params.GetPostFormNotEmpty(c, "change_type")

	if params.Exist("user_id") == false {
		c.JSON(200, routes.ApiShowResult(common.CodeFail, "用户编号不能为空"))
		return
	}
	if params.Exist("amount") == false {
		c.JSON(200, routes.ApiShowResult(common.CodeFail, "金额不能为空"))
		return
	}
	if params.Exist("change_type") == false {
		c.JSON(200, routes.ApiShowResult(common.CodeFail, "变动类型不能为空"))
		return
	}
	userID := common.GetInt(params.Get("user_id"))
	changeType := common.GetInt(params.Get("change_type"))
	amount, _ := decimal.NewFromString(params.Get("amount"))

	err = service.ManualAddMoney(uid, userID, amount, changeType)

	if err != nil {
		c.JSON(200, routes.ApiShowResult(common.CodeFail, fmt.Sprintf("%s", err)))
	} else {
		c.JSON(200, routes.ApiShowResult(common.CodeOK, ""))
	}
}
