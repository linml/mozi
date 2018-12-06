package api

import (
	"encoding/json"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/shopspring/decimal"
	"github.com/xiuos/mozi/common"
	"github.com/xiuos/mozi/models/lotto"
	"github.com/xiuos/mozi/routes"
	"github.com/xiuos/mozi/service"
)

type BetContent struct {
	MethodCode string          `json:"method_code"`
	BetContent string          `json:"bet_content"`
	Amount     decimal.Decimal `json:"amount"`
}

// 下单
// 参数：lotto_id，issue，bets
func Bet(c *gin.Context) {
	uid, err := routes.GetAPILoginID(c)
	if err != nil {
		c.JSON(200, routes.ApiShowResult(common.CodeFail, "登录校验失败"))
	}
	lottoID := common.GetInt(c.PostForm("lotto_id"))
	issue := c.PostForm("issue")
	//roomID := c.PostForm("room_id")
	bets := c.PostForm("bets")

	var betList []BetContent
	err = json.Unmarshal([]byte(bets), &betList)
	if err != nil {
		c.JSON(200, routes.ApiShowResult(common.CodeFail, "下注格式错误"))
		return
	}
	for i, _ := range betList {
		err = service.Bet(&lotto.Order{
			UserID:     uid,
			LottoID:    lottoID,
			Issue:      issue,
			MethodCode: betList[i].MethodCode,
			Content:    betList[i].BetContent,
			Amount:     betList[i].Amount,
		})
		if err != nil {
			c.JSON(200, routes.ApiShowResult(common.CodeFail, fmt.Sprintf("下注失败：%s", err)))
			return
		}
	}
	c.JSON(200, routes.ApiShowResult(common.CodeOK, "下单成功"))
	return

}
