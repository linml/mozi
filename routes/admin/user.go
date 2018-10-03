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
