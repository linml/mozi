package cms

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/xiuos/mozi/common"
	"github.com/xiuos/mozi/routes"
	"github.com/xiuos/mozi/service"
)

func CMSFindHallLotto(c *gin.Context) {

	dataList, err := service.CMSFindHallLottoList()
	if err != nil {
		c.JSON(200, routes.ApiResult(common.CodeFail, fmt.Sprintf("%s", err), map[string]string{}))
	} else {
		c.JSON(200, routes.ApiResult(common.CodeOK, "", dataList))
	}
}

func CMSHomeInit(c *gin.Context) {
	dataList, err := service.CMSHomeInit()
	if err != nil {
		c.JSON(200, routes.ApiResult(common.CodeFail, fmt.Sprintf("%s", err), map[string]string{}))
	} else {
		c.JSON(200, routes.ApiResult(common.CodeOK, "", dataList))
	}
}

func CMSBetPlay(c *gin.Context) {
	params := routes.ParamHelper{}
	params.GetQueryNotEmpty(c, "lotto_id")
	if params.Exist("lotto_id") == false {
		c.JSON(200, routes.ApiResult(common.CodeFail, "彩票不能为空", map[string]string{}))
		return
	}
	lid := common.GetInt(params.Get("lotto_id"))
	data, _ := service.CMSBetPlayInfo(1, lid)
	c.JSON(200, routes.ApiResult(common.CodeOK, "", data))
}
