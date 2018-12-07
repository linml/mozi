package admin

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/xiuos/mozi/common"
	"github.com/xiuos/mozi/models"
	"github.com/xiuos/mozi/models/lotto"
	"github.com/xiuos/mozi/routes"
	"github.com/xiuos/mozi/service"
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

func GetCodeLotto(c *gin.Context) {
	params := routes.ParamHelper{}

	params.GetQuery(c, "lotto_id")

	data, err := lotto.GetLotto(common.GetInt(params.Get("lotto_id")))
	if err != nil {
		c.JSON(200, routes.ApiResult(common.CodeFail, fmt.Sprintf("%s", err), map[string]string{}))
	} else {
		c.JSON(200, routes.ApiResult(common.CodeOK, "", data))
	}
}

func SetCodeLottoStatus(c *gin.Context) {
	uid, err := routes.GetAdminLoginID(c)
	if err != nil {
		c.JSON(200, routes.ApiShowResult(common.CodeFail, fmt.Sprintf("%s", err)))
		return
	}
	params := routes.ParamHelper{}

	params.GetPostFormNotEmpty(c, "lotto_id")
	params.GetPostFormNotEmpty(c, "status")

	if !params.Exist("lotto_id") {
		c.JSON(200, routes.ApiShowResult(common.CodeFail, "彩票编号不能为空"))
		return
	}
	if !params.Exist("status") {
		c.JSON(200, routes.ApiShowResult(common.CodeFail, "状态不能为空"))
		return
	}
	lid := common.GetInt(params.Get("lotto_id"))
	status := common.GetInt(params.Get("status"))

	beforeInfo, _ := lotto.GetLotto(lid)
	err = service.SetCodeLottoStatus(lid, status)

	if err != nil {
		c.JSON(200, routes.ApiShowResult(common.CodeFail, fmt.Sprintf("%s", err)))
	} else {
		au, _ := models.GetAdminUserByID(uid)
		r := models.RecordAdminAction{
			ActionModule: models.ActionAdminModuleSystem,
			ActionID:     models.ActionAdminSetCodeLottoStatus,
			Content:      fmt.Sprintf("修改彩票状态:彩票:%s,原状态:%d,现状态:%d", beforeInfo.Name, beforeInfo.Status, status),
			IP:           c.ClientIP(),
			RecordAt:     common.GetTimeNowString(),
			Success:      common.CodeOK,
			Message:      "操作成功",
			OperatorID:   uid,
			OperatorName: au.Name,
			OperatorType: models.OperatorTypeAdminSelf,
		}
		models.LogRecordAdminAction(&r)
		c.JSON(200, routes.ApiShowResult(common.CodeOK, ""))
	}
}

func SetCodeLottoIsShow(c *gin.Context) {
	uid, err := routes.GetAdminLoginID(c)
	if err != nil {
		c.JSON(200, routes.ApiShowResult(common.CodeFail, fmt.Sprintf("%s", err)))
		return
	}
	params := routes.ParamHelper{}

	params.GetPostFormNotEmpty(c, "lotto_id")
	params.GetPostFormNotEmpty(c, "is_show")

	if !params.Exist("lotto_id") {
		c.JSON(200, routes.ApiShowResult(common.CodeFail, "彩票编号不能为空"))
		return
	}
	if !params.Exist("is_show") {
		c.JSON(200, routes.ApiShowResult(common.CodeFail, "显示状态不能为空"))
		return
	}
	lid := common.GetInt(params.Get("lotto_id"))
	isShow := common.GetInt(params.Get("is_show"))

	beforeInfo, _ := lotto.GetLotto(lid)
	err = service.SetCodeLottoIsShow(lid, isShow)

	if err != nil {
		c.JSON(200, routes.ApiShowResult(common.CodeFail, fmt.Sprintf("%s", err)))
	} else {
		au, _ := models.GetAdminUserByID(uid)
		r := models.RecordAdminAction{
			ActionModule: models.ActionAdminModuleSystem,
			ActionID:     models.ActionAdminSetCodeLottoIsShow,
			Content:      fmt.Sprintf("修改彩票显示:彩票:%s,原状态:%d,现状态:%d", beforeInfo.Name, beforeInfo.IsShow, isShow),
			IP:           c.ClientIP(),
			RecordAt:     common.GetTimeNowString(),
			Success:      common.CodeOK,
			Message:      "操作成功",
			OperatorID:   uid,
			OperatorName: au.Name,
			OperatorType: models.OperatorTypeAdminSelf,
		}
		models.LogRecordAdminAction(&r)
		c.JSON(200, routes.ApiShowResult(common.CodeOK, ""))
	}
}

func SetCodeLottoSortIndex(c *gin.Context) {
	uid, err := routes.GetAdminLoginID(c)
	if err != nil {
		c.JSON(200, routes.ApiShowResult(common.CodeFail, fmt.Sprintf("%s", err)))
		return
	}
	params := routes.ParamHelper{}

	params.GetPostFormNotEmpty(c, "lotto_id")
	params.GetPostFormNotEmpty(c, "sort_index")

	if !params.Exist("lotto_id") {
		c.JSON(200, routes.ApiShowResult(common.CodeFail, "彩票编号不能为空"))
		return
	}
	if !params.Exist("sort_index") {
		c.JSON(200, routes.ApiShowResult(common.CodeFail, "排序不能为空"))
		return
	}
	lid := common.GetInt(params.Get("lotto_id"))
	sortIndex := common.GetInt(params.Get("sort_index"))

	beforeInfo, _ := lotto.GetLotto(lid)
	err = service.SetCodeLottoSortIndex(lid, sortIndex)

	if err != nil {
		c.JSON(200, routes.ApiShowResult(common.CodeFail, fmt.Sprintf("%s", err)))
	} else {
		au, _ := models.GetAdminUserByID(uid)
		r := models.RecordAdminAction{
			ActionModule: models.ActionAdminModuleSystem,
			ActionID:     models.ActionAdminSetCodeLottoSortIndex,
			Content:      fmt.Sprintf("修改彩票排序:彩票:%s,原状态:%d,现状态:%d", beforeInfo.Name, beforeInfo.SortIndex, sortIndex),
			IP:           c.ClientIP(),
			RecordAt:     common.GetTimeNowString(),
			Success:      common.CodeOK,
			Message:      "操作成功",
			OperatorID:   uid,
			OperatorName: au.Name,
			OperatorType: models.OperatorTypeAdminSelf,
		}
		models.LogRecordAdminAction(&r)
		c.JSON(200, routes.ApiShowResult(common.CodeOK, ""))
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

func FindCodeLottoTypeList(c *gin.Context) {
	params := routes.ParamHelper{}

	params.GetQuery(c, "id")
	params.GetQuery(c, "name")

	data, err := lotto.FindCodeLottoTypeList(params)
	if err != nil {
		c.JSON(200, routes.ApiResult(common.CodeFail, fmt.Sprintf("%s", err), map[string]string{}))
	} else {
		c.JSON(200, routes.ApiResult(common.CodeOK, "", data))
	}
}
