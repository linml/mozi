package admin

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/shopspring/decimal"
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

func FindLottoMethodTemplateList(c *gin.Context) {
	params := routes.ParamHelper{}

	params.GetQuery(c, "id")
	params.GetQuery(c, "name")

	data, err := lotto.FindLottoMethodTemplateList(params)
	if err != nil {
		c.JSON(200, routes.ApiResult(common.CodeFail, fmt.Sprintf("%s", err), map[string]string{}))
	} else {
		c.JSON(200, routes.ApiResult(common.CodeOK, "", data))
	}
}

func FindLottOddsList(c *gin.Context) {
	params := routes.ParamHelper{}

	params.GetQuery(c, "lotto_id")
	params.GetQuery(c, "method_code")
	params.GetQuery(c, "play_code")

	data, err := lotto.FindLottoOddsList(params)
	if err != nil {
		c.JSON(200, routes.ApiResult(common.CodeFail, fmt.Sprintf("%s", err), map[string]string{}))
	} else {
		c.JSON(200, routes.ApiResult(common.CodeOK, "", data))
	}
}

func PageFindLottoResultList(c *gin.Context) {
	params := routes.ParamHelper{}
	params.GetQuery(c, "draw")
	params.GetQuery(c, "curr_page")
	params.GetQuery(c, "page_row")

	params.GetQuery(c, "lotto_id")
	params.GetQuery(c, "issue")
	params.GetQuery(c, "draw_number")
	params.GetQuery(c, "status")
	params.GetQuery(c, "game_type")
	params.GetQuery(c, "status")
	params.GetQuery(c, "issue_date")
	params.GetQuery(c, "start_time_from")
	params.GetQuery(c, "start_time_to")
	params.GetQuery(c, "stop_time_from")
	params.GetQuery(c, "stop_time_to")
	params.GetQuery(c, "result_time_from")
	params.GetQuery(c, "result_time_to")
	params.GetQuery(c, "issue_date_from")
	params.GetQuery(c, "issue_date_to")
	params.GetQuery(c, "update_time_from")
	params.GetQuery(c, "update_time_to")
	params.GetQuery(c, "sort_type")

	currPage := common.GetInt(params.Get("curr_page"))
	pageRow := common.GetInt(params.Get("page_row"))

	if currPage < 1 {
		currPage = common.PageDefaultPage
	}

	if pageRow < 1 {
		pageRow = common.PageDefaultRow
	}
	if params.Exist("lotto_id") == false {
		c.JSON(200, routes.ApiResult(common.CodeFail, "彩票编号不能为空", map[string]string{}))
		return
	}
	lottoID := common.GetInt(params.Get("lotto_id"))
	pp := common.PageParams{CurrentPage: currPage, PageRow: pageRow, Params: params}
	pr, _, err := lotto.PageFindLottoResultList(lottoID, pp)
	if err != nil {
		c.JSON(200, routes.ApiResult(common.CodeFail, fmt.Sprintf("%s", err), map[string]string{}))
	} else {
		c.JSON(200, routes.ApiResult(common.CodeOK, "", pr))
	}
}

func PageFindLottOddsList(c *gin.Context) {
	params := routes.ParamHelper{}
	params.GetQuery(c, "draw")
	params.GetQuery(c, "curr_page")
	params.GetQuery(c, "page_row")

	params.GetQuery(c, "lotto_id")
	params.GetQuery(c, "method_code")
	params.GetQuery(c, "play_code")
	params.GetQuery(c, "status")
	params.GetQuery(c, "is_show")
	params.GetQuery(c, "sort_type")

	currPage := common.GetInt(params.Get("curr_page"))
	pageRow := common.GetInt(params.Get("page_row"))

	if currPage < 1 {
		currPage = common.PageDefaultPage
	}

	if pageRow < 1 {
		pageRow = common.PageDefaultRow
	}

	pp := common.PageParams{CurrentPage: currPage, PageRow: pageRow, Params: params}
	pr, _, err := lotto.PageFindLottoOddsList(pp)
	if err != nil {
		c.JSON(200, routes.ApiResult(common.CodeFail, fmt.Sprintf("%s", err), map[string]string{}))
	} else {
		c.JSON(200, routes.ApiResult(common.CodeOK, "", pr))
	}
}

func GetLottOdds(c *gin.Context) {
	params := routes.ParamHelper{}

	params.GetQueryNotEmpty(c, "lotto_id")
	params.GetQueryNotEmpty(c, "method_code")
	params.GetQueryNotEmpty(c, "play_code")

	if params.Exist("lotto_id") == false {
		c.JSON(200, routes.ApiResult(common.CodeFail, "彩票编号不能为空", map[string]string{}))
		return
	}
	if params.Exist("method_code") == false {
		c.JSON(200, routes.ApiResult(common.CodeFail, "玩法不能为空", map[string]string{}))
		return
	}
	if params.Exist("play_code") == false {
		c.JSON(200, routes.ApiResult(common.CodeFail, "玩法项不能为空", map[string]string{}))
		return
	}
	lid := common.GetInt(params.Get("lotto_id"))
	mCode := params.Get("method_code")
	pCode := params.Get("play_code")

	data, err := lotto.GetOdds(lid, mCode, pCode)
	if err != nil {
		c.JSON(200, routes.ApiResult(common.CodeFail, fmt.Sprintf("%s", err), map[string]string{}))
	} else {
		c.JSON(200, routes.ApiResult(common.CodeOK, "", data))
	}
}

func SetLottOddsInfo(c *gin.Context) {
	uid, err := routes.GetAdminLoginID(c)
	if err != nil {
		c.JSON(200, routes.ApiShowResult(common.CodeFail, fmt.Sprintf("%s", err)))
		return
	}
	params := routes.ParamHelper{}

	params.GetPostFormNotEmpty(c, "lotto_id")
	params.GetPostFormNotEmpty(c, "method_code")
	params.GetPostFormNotEmpty(c, "play_code")

	params.GetPostForm(c, "odds")
	params.GetPostForm(c, "odds_min")
	params.GetPostForm(c, "odds_max")
	params.GetPostForm(c, "bet_min")
	params.GetPostForm(c, "bet_max")
	params.GetPostForm(c, "status")
	params.GetPostForm(c, "is_show")

	if params.Exist("lotto_id") == false {
		c.JSON(200, routes.ApiShowResult(common.CodeFail, "彩票编号不能为空"))
		return
	}
	if params.Exist("method_code") == false {
		c.JSON(200, routes.ApiShowResult(common.CodeFail, "玩法不能为空"))
		return
	}
	if params.Exist("play_code") == false {
		c.JSON(200, routes.ApiShowResult(common.CodeFail, "玩法项不能为空"))
		return
	}

	lid := common.GetInt(params.Get("lotto_id"))
	mCode := params.Get("method_code")
	pCode := params.Get("play_code")
	beforeInfo, err := lotto.GetOdds(lid, mCode, pCode)
	if err != nil {
		c.JSON(200, routes.ApiShowResult(common.CodeFail, fmt.Sprintf("找不到该玩法%s", err)))
		return
	}

	lottoInfo, err := lotto.GetLotto(lid)
	if err != nil {
		c.JSON(200, routes.ApiShowResult(common.CodeFail, fmt.Sprintf("找不到该彩票%s", err)))
		return
	}
	content := fmt.Sprintf("修改赔率信息:彩票:%s;玩法:%s;玩法项:%s;", lottoInfo.Name, beforeInfo.MethodName, beforeInfo.PlayName)

	if params.Exist("odds") == true {
		odds, err := decimal.NewFromString(params.Get("odds"))
		if err != nil {
			c.JSON(200, routes.ApiShowResult(common.CodeFail, "请填入规范数字"))
			return
		}
		err = service.SetLottoOdds(lid, mCode, pCode, odds)
		if err != nil {
			content += fmt.Sprintf("设置赔率失败:%s", err)
		} else {
			content += fmt.Sprintf("设置赔率成功:原赔率%s,现赔率:%s", beforeInfo.Odds, odds)
		}
	}

	if params.Exist("odds_min") == true {
		oddsMin, err := decimal.NewFromString(params.Get("odds_min"))
		if err != nil {
			c.JSON(200, routes.ApiShowResult(common.CodeFail, "请填入规范数字"))
			return
		}
		err = service.SetLottoOddsMin(lid, mCode, pCode, oddsMin)
		if err != nil {
			content += fmt.Sprintf("设置最小赔率失败:%s", err)
		} else {
			content += fmt.Sprintf("设置最小赔率成功:原最小赔率%s,现最小赔率:%s", beforeInfo.OddsMin, oddsMin)
		}
	}

	if params.Exist("odds_max") == true {
		oddsMax, err := decimal.NewFromString(params.Get("odds_max"))
		if err != nil {
			c.JSON(200, routes.ApiShowResult(common.CodeFail, "请填入规范数字"))
			return
		}
		err = service.SetLottoOddsMax(lid, mCode, pCode, oddsMax)
		if err != nil {
			content += fmt.Sprintf("设置最大赔率失败:%s", err)
		} else {
			content += fmt.Sprintf("设置最大赔率成功:原最大赔率%s,现最大赔率:%s", beforeInfo.OddsMin, oddsMax)
		}
	}

	if params.Exist("bet_min") == true {
		betMin, err := decimal.NewFromString(params.Get("bet_min"))
		if err != nil {
			c.JSON(200, routes.ApiShowResult(common.CodeFail, "请填入规范数字"))
			return
		}
		err = service.SetLottoOddsBetMin(lid, mCode, pCode, betMin)
		if err != nil {
			content += fmt.Sprintf("设置最小额度失败:%s", err)
		} else {
			content += fmt.Sprintf("设置最小额度成功:原最小额度%s,现最小额度:%s", beforeInfo.BetMin, betMin)
		}
	}

	if params.Exist("bet_max") == true {
		betMax, err := decimal.NewFromString(params.Get("bet_max"))
		if err != nil {
			c.JSON(200, routes.ApiShowResult(common.CodeFail, "请填入规范数字"))
			return
		}
		err = service.SetLottoOddsBetMax(lid, mCode, pCode, betMax)
		if err != nil {
			content += fmt.Sprintf("设置最大额度失败:%s", err)
		} else {
			content += fmt.Sprintf("设置最大额度成功:原最大额度%s,现最大额度:%s", beforeInfo.BetMax, betMax)
		}
	}

	if params.Exist("status") == true {
		status := common.GetInt(params.Get("status"))
		err = service.SetLottoOddsStatus(lid, mCode, pCode, status)
		if err != nil {
			content += fmt.Sprintf("设置状态:%s", err)
		} else {
			content += fmt.Sprintf("设置状态成功:原状态%d,现状态:%d", beforeInfo.Status, status)
		}
	}
	if params.Exist("is_show") == true {
		isShow := common.GetInt(params.Get("status"))
		err = service.SetLottoOddsIsShow(lid, mCode, pCode, isShow)
		if err != nil {
			content += fmt.Sprintf("设置显示失败:%s", err)
		} else {
			content += fmt.Sprintf("设置显示成功:原显示%d,现显示:%d", beforeInfo.IsShow, isShow)
		}
	}
	if err != nil {
		c.JSON(200, routes.ApiShowResult(common.CodeFail, fmt.Sprintf("%s", err)))
	} else {
		c.JSON(200, routes.ApiShowResult(common.CodeOK, ""))
	}

	au, _ := models.GetAdminUserByID(uid)
	r := models.RecordAdminAction{
		ActionModule: models.ActionAdminModuleSystem,
		ActionID:     models.ActionAdminSetLottoOddsInfo,
		Content:      content,
		IP:           c.ClientIP(),
		RecordAt:     common.GetTimeNowString(),
		Success:      common.CodeOK,
		Message:      "操作成功",
		OperatorID:   uid,
		OperatorName: au.Name,
		OperatorType: models.OperatorTypeAdminSelf,
	}
	models.LogRecordAdminAction(&r)

}
