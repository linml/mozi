package admin

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/xiuos/mozi/common"
	"github.com/xiuos/mozi/models"
	"github.com/xiuos/mozi/routes"
	"github.com/xiuos/mozi/service"
	"strings"
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

func AddAdminRole(c *gin.Context) {
	uid, err := routes.GetAdminLoginID(c)
	if err != nil {
		c.JSON(200, routes.ApiShowResult(common.CodeFail, fmt.Sprintf("%s", err)))
		return
	}
	params := routes.ParamHelper{}
	params.GetPostForm(c, "name")
	name := params.Get("name")
	if name == "" {
		c.JSON(200, routes.ApiShowResult(common.CodeFail, "角色名不能为空"))
		return
	}
	ar := models.AdminRole{Name: name}
	err = models.CreateAdminRole(&ar)

	if err != nil {
		c.JSON(200, routes.ApiShowResult(common.CodeFail, fmt.Sprintf("%s", err)))
	} else {
		au, _ := models.GetAdminUserByID(uid)
		r := models.RecordAdminAction{
			ActionModule: models.ActionAdminModuleSystem,
			ActionID:     models.ActionAdminAddRole,
			Content:      fmt.Sprintf("添加新角色:%s", name),
			IP:           c.ClientIP(),
			RecordAt:     common.GetTimeNowString(),
			Success:      common.CodeOK,
			Message:      "添加成功",
			OperatorID:   uid,
			OperatorName: au.Name,
			OperatorType: models.OperatorTypeAdminSelf,
		}
		models.LogRecordAdminAction(&r)
		c.JSON(200, routes.ApiShowResult(common.CodeOK, ""))
	}
}

func UpdateRoleMenu(c *gin.Context) {
	uid, err := routes.GetAdminLoginID(c)
	if err != nil {
		c.JSON(200, routes.ApiShowResult(common.CodeFail, fmt.Sprintf("%s", err)))
		return
	}
	params := routes.ParamHelper{}
	params.GetPostForm(c, "role_id")
	params.GetPostForm(c, "menu_id_list")

	roleID := common.GetInt(params.Get("role_id"))
	menuIDList := strings.Split(params.Get("menu_id_list"), ",")
	mIDList, _ := common.S2IList(menuIDList)
	err = service.UpdateRoleMenu(roleID, mIDList)

	if err != nil {
		c.JSON(200, routes.ApiShowResult(common.CodeFail, fmt.Sprintf("%s", err)))
	} else {
		au, _ := models.GetAdminUserByID(uid)
		r := models.RecordAdminAction{
			ActionModule: models.ActionAdminModuleSystem,
			ActionID:     models.ActionAdminSetRoleMenu,
			Content:      fmt.Sprintf("设置角色菜单:%s", params.Get("menu_id_list")),
			IP:           c.ClientIP(),
			RecordAt:     common.GetTimeNowString(),
			Success:      common.CodeOK,
			Message:      "添加成功",
			OperatorID:   uid,
			OperatorName: au.Name,
			OperatorType: models.OperatorTypeAdminSelf,
		}
		models.LogRecordAdminAction(&r)
		c.JSON(200, routes.ApiShowResult(common.CodeOK, ""))
	}
}
