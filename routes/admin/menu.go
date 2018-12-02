package admin

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/xiuos/mozi/common"
	"github.com/xiuos/mozi/models"
	"github.com/xiuos/mozi/routes"
)

func AdminMenu(c *gin.Context) {
	uid, err := routes.GetAdminLoginID(c)
	if err != nil {
		c.JSON(200, routes.ApiResult(common.CodeFail, fmt.Sprintf("%s", err), map[string]string{}))
		return
	}
	t, err := models.GetAdminMenuTree(uid)
	if err != nil {
		c.JSON(200, routes.ApiResult(common.CodeFail, fmt.Sprintf("%s", err), map[string]string{}))
	} else {
		c.JSON(200, routes.ApiResult(common.CodeOK, "", t.Child))
	}
}

func FindAdminMenuNode(c *gin.Context) {
	_, err := routes.GetAdminLoginID(c)
	if err != nil {
		c.JSON(200, routes.ApiResult(common.CodeFail, fmt.Sprintf("%s", err), map[string]string{}))
		return
	}
	nodeList, err := models.FindAdminMenu(map[string]string{})
	if err != nil {
		c.JSON(200, routes.ApiResult(common.CodeFail, fmt.Sprintf("%s", err), map[string]string{}))
	} else {
		c.JSON(200, routes.ApiResult(common.CodeOK, "", nodeList))
	}
}

func FindAdminRole(c *gin.Context) {
	_, err := routes.GetAdminLoginID(c)
	if err != nil {
		c.JSON(200, routes.ApiResult(common.CodeFail, fmt.Sprintf("%s", err), map[string]string{}))
		return
	}

	params := routes.ParamHelper{}

	dataList, err := models.FindAdminRole(params)
	if err != nil {
		c.JSON(200, routes.ApiResult(common.CodeFail, fmt.Sprintf("%s", err), map[string]string{}))
	} else {
		c.JSON(200, routes.ApiResult(common.CodeOK, "", dataList))
	}
}

func FindAdminRoleMenu(c *gin.Context) {
	_, err := routes.GetAdminLoginID(c)
	if err != nil {
		c.JSON(200, routes.ApiResult(common.CodeFail, fmt.Sprintf("%s", err), map[string]string{}))
		return
	}
	params := routes.ParamHelper{}
	params.GetQuery(c, "role_id")
	dataList, err := models.FindRoleMenu(params)
	if err != nil {
		c.JSON(200, routes.ApiResult(common.CodeFail, fmt.Sprintf("%s", err), map[string]string{}))
	} else {
		c.JSON(200, routes.ApiResult(common.CodeOK, "", dataList))
	}
}
