package admin

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/xiuos/mozi/common"
	"github.com/xiuos/mozi/models"
	"github.com/xiuos/mozi/routes"
	"github.com/xiuos/mozi/service"
)

func PageFindUserList(c *gin.Context) {
	params := routes.ParamHelper{}
	params.GetQuery(c, "draw")
	params.GetQuery(c, "user_id")
	params.GetQuery(c, "name")
	params.GetQuery(c, "login_ip")
	params.GetQuery(c, "status")
	params.GetQuery(c, "parent_id")
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
	pr, _, err := models.PageFindUserList(pp)
	if err != nil {
		c.JSON(200, routes.ApiResult(common.CodeFail, fmt.Sprintf("%s", err), map[string]string{}))
	} else {
		c.JSON(200, routes.ApiResult(common.CodeOK, "", pr))
	}
}

func PageFindAdminList(c *gin.Context) {
	params := routes.ParamHelper{}
	params.GetQuery(c, "draw")
	params.GetQuery(c, "user_id")
	params.GetQuery(c, "name")
	params.GetQuery(c, "status")
	params.GetQuery(c, "role")
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
	pr, _, err := models.PageFindAdmin(pp)
	if err != nil {
		c.JSON(200, routes.ApiResult(common.CodeFail, fmt.Sprintf("%s", err), map[string]string{}))
	} else {
		c.JSON(200, routes.ApiResult(common.CodeOK, "", pr))
	}
}

func PageFindUserRecordLogin(c *gin.Context) {
	params := routes.ParamHelper{}
	params.GetQuery(c, "draw")
	params.GetQuery(c, "name")
	params.GetQuery(c, "ip")
	params.GetQuery(c, "status")
	params.GetQuery(c, "start_at")
	params.GetQuery(c, "end_at")
	params.GetQuery(c, "status")
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
	pr, _, err := models.PageFindUserRecordLogin4Admin(pp)
	if err != nil {
		c.JSON(200, routes.ApiResult(common.CodeFail, fmt.Sprintf("%s", err), map[string]string{}))
	} else {
		c.JSON(200, routes.ApiResult(common.CodeOK, "", pr))
	}
}

// 管理员添加新用户
func AddMember(c *gin.Context) {
	uid, err := routes.GetAdminLoginID(c)
	if err != nil {
		c.JSON(200, routes.ApiShowResult(common.CodeFail, fmt.Sprintf("%s", err)))
		return
	}
	params := routes.ParamHelper{}
	params.GetPostForm(c, "is_new_line")
	params.GetPostForm(c, "parent_name")
	params.GetPostForm(c, "name")
	params.GetPostForm(c, "password")

	isNewLine := common.GetInt(params.Get("is_new_line"))

	param := map[string]string{}

	recordText := fmt.Sprintf("新增用户:用户名:%s;", params.Get("name"))
	if isNewLine == 1 {
		param["is_new_line"] = "1"
		recordText += "新开线;"
	} else {
		parentInfo, err := models.GetUserByName(params.Get("parent_name"))
		if err != nil {
			c.JSON(200, routes.ApiResult(common.CodeFail, fmt.Sprintf("%s", "找不到此上级用户"), map[string]string{}))
			return
		}
		param["parent_id"] = fmt.Sprintf("%d", parentInfo.UserID)
		recordText += fmt.Sprintf("指定上级:%s", parentInfo.Name)
	}
	err = service.RegisterUser(params.Get("name"), params.Get("password"), param, models.OperatorTypeAdmin)

	au, _ := models.GetAdminUserByID(uid)
	r := models.RecordAdminAction{
		ActionModule: models.ActionAdminModuleUser,
		ActionID:     models.ActionAdminAddMember,
		Content:      recordText,
		IP:           c.ClientIP(),
		RecordAt:     common.GetTimeNowString(),
		OperatorID:   uid,
		OperatorName: au.Name,
		OperatorType: models.OperatorTypeAdminSelf,
	}

	if err != nil {
		c.JSON(200, routes.ApiShowResult(common.CodeFail, fmt.Sprintf("%s", err)))
		r.Success = common.CodeFail
		r.Message = fmt.Sprintf("%s", err)
	} else {
		r.Success = common.CodeOK
		r.Message = "添加成功"
		c.JSON(200, routes.ApiShowResult(common.CodeOK, "添加成功"))
	}
	models.LogRecordAdminAction(&r)
}

func GetMemberInfos(c *gin.Context) {
	uid := common.GetInt(c.Query("user_id"))
	infos, err := service.GetInfos(uid)
	if err != nil {
		c.JSON(200, routes.ApiResult(common.CodeFail, fmt.Sprintf("%s", err), map[string]string{}))
	} else {
		c.JSON(200, routes.ApiResult(common.CodeOK, "", infos))
	}
}

func GetMemberProfile(c *gin.Context) {
	uid := common.GetInt(c.Query("user_id"))
	infos, err := models.GetUserProfile(uid)
	if err != nil {
		c.JSON(200, routes.ApiResult(common.CodeFail, fmt.Sprintf("%s", err), map[string]string{}))
	} else {
		c.JSON(200, routes.ApiResult(common.CodeOK, "", infos))
	}
}

func GetMemberRelation(c *gin.Context) {
	uid := common.GetInt(c.Query("user_id"))
	infos, err := models.GetUserRelation(uid)
	if err != nil {
		c.JSON(200, routes.ApiResult(common.CodeFail, fmt.Sprintf("%s", err), map[string]string{}))
	} else {
		c.JSON(200, routes.ApiResult(common.CodeOK, "", infos))
	}
}

func GetMemberPower(c *gin.Context) {
	uid := common.GetInt(c.Query("user_id"))
	infos, err := models.GetUserPower(uid)
	if err != nil {
		c.JSON(200, routes.ApiResult(common.CodeFail, fmt.Sprintf("%s", err), map[string]string{}))
	} else {
		c.JSON(200, routes.ApiResult(common.CodeOK, "", infos))
	}
}

func SetAdminRole(c *gin.Context) {
	uid, err := routes.GetAdminLoginID(c)
	if err != nil {
		c.JSON(200, routes.ApiShowResult(common.CodeFail, fmt.Sprintf("%s", err)))
		return
	}
	params := routes.ParamHelper{}
	params.GetPostForm(c, "user_id")
	params.GetPostForm(c, "role_id")

	userID := common.GetInt(params.Get("user_id"))
	roleID := common.GetInt(params.Get("role_id"))

	beforeUser, err := models.GetAdminUserByID(userID)
	if err != nil {
		c.JSON(200, routes.ApiShowResult(common.CodeFail, fmt.Sprintf("找不到该用户%s", err)))
		return
	}
	var beforeRoleName, nowRoleName string
	bInfo, err := models.GetAdminRole(beforeUser.Role)
	if err != nil {

	} else {
		beforeRoleName = bInfo.Name
	}

	err = service.SetAdminRole(userID, roleID)
	afterUser, err := models.GetAdminUserByID(userID)
	nInfo, err := models.GetAdminRole(afterUser.Role)
	if err != nil {

	} else {
		nowRoleName = nInfo.Name
	}

	if err != nil {
		c.JSON(200, routes.ApiShowResult(common.CodeFail, fmt.Sprintf("%s", err)))
	} else {
		au, _ := models.GetAdminUserByID(uid)
		r := models.RecordAdminAction{
			ActionModule: models.ActionAdminModuleSystem,
			ActionID:     models.ActionAdminSetAdminRole,
			Content:      fmt.Sprintf("修改管理员角色:用户名:%s,原角色:%s,现角色:%s", beforeUser.Name, beforeRoleName, nowRoleName),
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

func SetAdminStatus(c *gin.Context) {
	uid, err := routes.GetAdminLoginID(c)
	if err != nil {
		c.JSON(200, routes.ApiShowResult(common.CodeFail, fmt.Sprintf("%s", err)))
		return
	}
	params := routes.ParamHelper{}
	params.GetPostForm(c, "user_id")
	params.GetPostForm(c, "status")

	userID := common.GetInt(params.Get("user_id"))
	status := common.GetInt(params.Get("status"))

	beforeUser, err := models.GetAdminUserByID(userID)
	if err != nil {
		c.JSON(200, routes.ApiShowResult(common.CodeFail, fmt.Sprintf("找不到该用户%s", err)))
		return
	}

	err = service.SetAdminStatus(userID, status)

	if err != nil {
		c.JSON(200, routes.ApiShowResult(common.CodeFail, fmt.Sprintf("%s", err)))
	} else {
		au, _ := models.GetAdminUserByID(uid)
		r := models.RecordAdminAction{
			ActionModule: models.ActionAdminModuleSystem,
			ActionID:     models.ActionAdminSetAdminStatus,
			Content:      fmt.Sprintf("修改管理员状态:用户名:%s,原状态:%d,现状态:%d", beforeUser.Name, beforeUser.Status, status),
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

func CreateAdmin(c *gin.Context) {
	uid, err := routes.GetAdminLoginID(c)
	if err != nil {
		c.JSON(200, routes.ApiShowResult(common.CodeFail, fmt.Sprintf("%s", err)))
		return
	}
	params := routes.ParamHelper{}
	params.GetPostForm(c, "name")
	params.GetPostForm(c, "password")
	params.GetPostForm(c, "role")

	name := params.Get("name")
	password := params.Get("password")
	role := common.GetInt(params.Get("role"))

	err = service.AddAdmin(name, password, role)

	if err != nil {
		c.JSON(200, routes.ApiShowResult(common.CodeFail, fmt.Sprintf("%s", err)))
	} else {
		au, _ := models.GetAdminUserByID(uid)
		r := models.RecordAdminAction{
			ActionModule: models.ActionAdminModuleSystem,
			ActionID:     models.ActionAdminAddAdmin,
			Content:      fmt.Sprintf("添加管理员:用户名:%s", name),
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
