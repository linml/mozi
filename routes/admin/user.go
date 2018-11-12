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
	params.GetQuery(c, "username")
	params.GetQuery(c, "login_ip")
	params.GetQuery(c, "status")
	params.GetQuery(c, "is_test")
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

func PageFindUserRecordLogin(c *gin.Context) {
	params := routes.ParamHelper{}
	params.GetQuery(c, "draw")
	params.GetQuery(c, "username")
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
	params := routes.ParamHelper{}
	params.GetPostForm(c, "is_new_line")
	params.GetPostForm(c, "parent_name")
	params.GetPostForm(c, "name")
	params.GetPostForm(c, "password")

	isNewLine := common.GetInt(params.Get("is_new_line"))

	param := map[string]string{}
	if isNewLine == 1 {
		param["is_new_line"] = "1"
	} else {
		parentInfo, err := models.GetUserByName(params.Get("parent_name"))
		if err != nil {
			c.JSON(200, routes.ApiResult(common.CodeFail, fmt.Sprintf("%s", "找不到此上级用户"), map[string]string{}))
			return
		}
		param["parent_id"] = fmt.Sprintf("%d", parentInfo.UserID)
	}
	fmt.Println(param)
	err := service.RegisterUser(params.Get("name"), params.Get("password"), param, models.OperatorTypeAdmin)

	if err != nil {
		c.JSON(200, routes.ApiShowResult(common.CodeFail, fmt.Sprintf("%s", err)))
	} else {
		c.JSON(200, routes.ApiShowResult(common.CodeOK, "添加成功"))
	}
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
