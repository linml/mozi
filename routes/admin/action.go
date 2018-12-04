package admin

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/xiuos/mozi/common"
	"github.com/xiuos/mozi/models"
	"github.com/xiuos/mozi/routes"
)

func FindCodeActionAdminModuleList(c *gin.Context) {
	_, err := routes.GetAdminLoginID(c)
	if err != nil {
		c.JSON(200, routes.ApiResult(common.CodeFail, fmt.Sprintf("%s", err), map[string]string{}))
		return
	}
	dataList, err := models.FindCodeActionAdminModule(map[string]string{"status": "1"})

	if err != nil {
		c.JSON(200, routes.ApiResult(common.CodeFail, fmt.Sprintf("%s", err), map[string]string{}))
	} else {
		c.JSON(200, routes.ApiResult(common.CodeOK, "", dataList))
	}
}

func FindCodeActionAdminList(c *gin.Context) {
	_, err := routes.GetAdminLoginID(c)
	if err != nil {
		c.JSON(200, routes.ApiResult(common.CodeFail, fmt.Sprintf("%s", err), map[string]string{}))
		return
	}
	dataList, err := models.FindCodeActionAdmin(map[string]string{"status": "1"})

	if err != nil {
		c.JSON(200, routes.ApiResult(common.CodeFail, fmt.Sprintf("%s", err), map[string]string{}))
	} else {
		c.JSON(200, routes.ApiResult(common.CodeOK, "", dataList))
	}
}
