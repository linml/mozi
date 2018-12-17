package admin

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/xiuos/mozi/common"
	"github.com/xiuos/mozi/models"
	"github.com/xiuos/mozi/routes"
)

func PageFindNoticeList(c *gin.Context) {
	params := routes.ParamHelper{}
	params.GetQuery(c, "draw")
	params.GetQuery(c, "curr_page")
	params.GetQuery(c, "page_row")

	params.GetQuery(c, "id")
	params.GetQuery(c, "class_id")
	params.GetQuery(c, "status")
	params.GetQuery(c, "start_at")
	params.GetQuery(c, "end_at")

	currPage := common.GetInt(params.Get("curr_page"))
	pageRow := common.GetInt(params.Get("page_row"))

	if currPage < 1 {
		currPage = common.PageDefaultPage
	}

	if pageRow < 1 {
		pageRow = common.PageDefaultRow
	}

	p := common.PageParams{CurrentPage: currPage, PageRow: pageRow, Params: params}
	pr, _, err := models.PageFindNoticeList(p)
	if err != nil {
		c.JSON(200, routes.ApiResult(common.CodeFail, fmt.Sprintf("%s", err), map[string]string{}))
	} else {
		c.JSON(200, routes.ApiResult(common.CodeOK, "", pr))
	}
}

func GetNotice(c *gin.Context) {
	params := routes.ParamHelper{}

	params.GetQueryNotEmpty(c, "id")
	if params.Exist("id") == false {
		c.JSON(200, routes.ApiResult(common.CodeFail, "id不能为空", map[string]string{}))
	}

	id := common.GetInt(params.Get("id"))

	data, err := models.GetNotice(id)
	if err != nil {
		c.JSON(200, routes.ApiResult(common.CodeFail, fmt.Sprintf("%s", err), map[string]string{}))
	} else {
		c.JSON(200, routes.ApiResult(common.CodeOK, "", data))
	}
}

func SetNotice(c *gin.Context) {
	uid, err := routes.GetAdminLoginID(c)
	if err != nil {
		c.JSON(200, routes.ApiShowResult(common.CodeFail, fmt.Sprintf("%s", err)))
		return
	}

	params := routes.ParamHelper{}

	params.GetPostFormNotEmpty(c, "id")
	params.GetPostFormNotEmpty(c, "class_id")
	params.GetPostFormNotEmpty(c, "title")
	params.GetPostFormNotEmpty(c, "content")
	params.GetPostFormNotEmpty(c, "sort_index")
	params.GetPostFormNotEmpty(c, "status")
	params.GetPostFormNotEmpty(c, "start_at")
	params.GetPostFormNotEmpty(c, "end_at")
	if params.Exist("id") == false {
		c.JSON(200, routes.ApiShowResult(common.CodeFail, "id不能为空"))
		return
	}
	if params.Exist("class_id") == false {
		c.JSON(200, routes.ApiShowResult(common.CodeFail, "class_id不能为空"))
		return
	}
	if params.Exist("content") == false {
		c.JSON(200, routes.ApiShowResult(common.CodeFail, "content不能为空"))
		return
	}
	if params.Exist("sort_index") == false {
		c.JSON(200, routes.ApiShowResult(common.CodeFail, "sort_index不能为空"))
		return
	}
	if params.Exist("status") == false {
		c.JSON(200, routes.ApiShowResult(common.CodeFail, "status不能为空"))
		return
	}
	if params.Exist("start_at") == false {
		c.JSON(200, routes.ApiShowResult(common.CodeFail, "start_at不能为空"))
		return
	}
	if params.Exist("end_at") == false {
		c.JSON(200, routes.ApiShowResult(common.CodeFail, "end_at不能为空"))
		return
	}

	notice := models.Notice{
		ID:        common.GetInt(params.Get("id")),
		ClassID:   common.GetInt(params.Get("class_id")),
		Title:     params.Get("title"),
		Content:   params.Get("content"),
		SortIndex: common.GetInt(params.Get("sort_index")),
		Status:    common.GetInt(params.Get("status")),
		StartAt:   params.Get("start_at"),
		EndAt:     params.Get("end_at"),
		UpdateAt:  common.GetTimeNowString(),
	}

	err = models.SetNotice(&notice)
	content := fmt.Sprintf("编号:%d;类别:%d;标题:%s;内容:%s;排序:%d;状态:%d;起始:%s;截止:%s", common.GetInt(params.Get("id")), common.GetInt(params.Get("class_id")),
		params.Get("title"), params.Get("content"), common.GetInt(params.Get("sort_index")), common.GetInt(params.Get("status")), params.Get("start_at"), params.Get("end_at"))

	au, _ := models.GetAdminUserByID(uid)
	r := models.RecordAdminAction{
		ActionModule: models.ActionAdminModuleSystem,
		ActionID:     models.ActionAdminUpdateNotice,
		Content:      content,
		IP:           c.ClientIP(),
		RecordAt:     common.GetTimeNowString(),
		Success:      common.CodeOK,
		Message:      "操作成功",
		OperatorID:   uid,
		OperatorName: au.Name,
		OperatorType: models.OperatorTypeAdminSelf,
	}

	if err != nil {
		r.Success = common.CodeFail
		r.Message = fmt.Sprintf("%s", err)
		c.JSON(200, routes.ApiShowResult(common.CodeFail, fmt.Sprintf("%s", err)))
	} else {
		c.JSON(200, routes.ApiShowResult(common.CodeOK, ""))
	}
	models.LogRecordAdminAction(&r)
}

func AddNotice(c *gin.Context) {
	uid, err := routes.GetAdminLoginID(c)
	if err != nil {
		c.JSON(200, routes.ApiShowResult(common.CodeFail, fmt.Sprintf("%s", err)))
		return
	}

	params := routes.ParamHelper{}

	params.GetPostFormNotEmpty(c, "class_id")
	params.GetPostFormNotEmpty(c, "title")
	params.GetPostFormNotEmpty(c, "content")
	params.GetPostFormNotEmpty(c, "sort_index")
	params.GetPostFormNotEmpty(c, "status")
	params.GetPostFormNotEmpty(c, "start_at")
	params.GetPostFormNotEmpty(c, "end_at")

	if params.Exist("class_id") == false {
		c.JSON(200, routes.ApiShowResult(common.CodeFail, "class_id不能为空"))
		return
	}
	if params.Exist("content") == false {
		c.JSON(200, routes.ApiShowResult(common.CodeFail, "content不能为空"))
		return
	}
	if params.Exist("sort_index") == false {
		c.JSON(200, routes.ApiShowResult(common.CodeFail, "sort_index不能为空"))
		return
	}
	if params.Exist("status") == false {
		c.JSON(200, routes.ApiShowResult(common.CodeFail, "status不能为空"))
		return
	}
	if params.Exist("start_at") == false {
		c.JSON(200, routes.ApiShowResult(common.CodeFail, "start_at不能为空"))
		return
	}
	if params.Exist("end_at") == false {
		c.JSON(200, routes.ApiShowResult(common.CodeFail, "end_at不能为空"))
		return
	}

	notice := models.Notice{
		ID:        common.GetInt(params.Get("id")),
		ClassID:   common.GetInt(params.Get("class_id")),
		Title:     params.Get("title"),
		Content:   params.Get("content"),
		SortIndex: common.GetInt(params.Get("sort_index")),
		Status:    common.GetInt(params.Get("status")),
		StartAt:   params.Get("start_at"),
		EndAt:     params.Get("end_at"),
		UpdateAt:  common.GetTimeNowString(),
	}

	err = models.AddNotice(&notice)
	content := fmt.Sprintf("类别:%d;标题:%s;内容:%s;排序:%d;状态:%d;起始:%s;截止:%s", common.GetInt(params.Get("class_id")),
		params.Get("title"), params.Get("content"), common.GetInt(params.Get("sort_index")), common.GetInt(params.Get("status")), params.Get("start_at"), params.Get("end_at"))

	au, _ := models.GetAdminUserByID(uid)
	r := models.RecordAdminAction{
		ActionModule: models.ActionAdminModuleSystem,
		ActionID:     models.ActionAdminAddNotice,
		Content:      content,
		IP:           c.ClientIP(),
		RecordAt:     common.GetTimeNowString(),
		Success:      common.CodeOK,
		Message:      "操作成功",
		OperatorID:   uid,
		OperatorName: au.Name,
		OperatorType: models.OperatorTypeAdminSelf,
	}

	if err != nil {
		r.Success = common.CodeFail
		r.Message = fmt.Sprintf("%s", err)
		c.JSON(200, routes.ApiShowResult(common.CodeFail, fmt.Sprintf("%s", err)))
	} else {
		c.JSON(200, routes.ApiShowResult(common.CodeOK, ""))
	}
	models.LogRecordAdminAction(&r)
}
