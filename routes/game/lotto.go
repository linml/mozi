package game

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/xiuos/mozi/common"
	"github.com/xiuos/mozi/routes"
	"github.com/xiuos/mozi/service"
)

func DrawLotto(c *gin.Context) {

	params := routes.ParamHelper{}

	params.GetPostFormNotEmpty(c, "lotto_id")
	params.GetPostFormNotEmpty(c, "issue")
	params.GetPostFormNotEmpty(c, "numbers")

	if params.Exist("lotto_id") == false {
		c.JSON(200, routes.ApiShowResult(common.CodeFail, "彩票不能为空"))
		return
	}
	if params.Exist("issue") == false {
		fmt.Println("33")
		c.JSON(200, routes.ApiShowResult(common.CodeFail, "期号不能为空"))
		return
	}

	if params.Exist("numbers") == false {
		c.JSON(200, routes.ApiShowResult(common.CodeFail, "开奖号码不能为空"))
		return
	}

	lottoID := common.GetInt(params.Get("lotto_id"))
	issue := params.Get("issue")
	numbers := params.Get("numbers")

	err := service.CallDrawLotto(lottoID, issue, numbers)

	if err != nil {
		c.JSON(200, routes.ApiShowResult(common.CodeFail, fmt.Sprintf("%s", err)))
	} else {
		c.JSON(200, routes.ApiShowResult(common.CodeOK, ""))
	}

}
