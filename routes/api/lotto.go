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
	PlayCode   string          `json:"play_code"`
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

	params := routes.ParamHelper{}

	params.GetPostFormNotEmpty(c, "lotto_id")
	params.GetPostFormNotEmpty(c, "issue")
	params.GetPostFormNotEmpty(c, "bets")

	if params.Exist("lotto_id") == false {
		c.JSON(200, routes.ApiShowResult(common.CodeFail, "彩票不能为空"))
		return
	}
	if params.Exist("issue") == false {
		c.JSON(200, routes.ApiShowResult(common.CodeFail, "期号不能为空"))
		return
	}

	if params.Exist("bets") == false {
		c.JSON(200, routes.ApiShowResult(common.CodeFail, "下注内容不能为空"))
		return
	}

	lottoID := common.GetInt(params.Get("lotto_id"))
	issue := params.Get("issue")
	bets := params.Get("bets")
	//ip := common.InetAton(c.ClientIP())

	var betList []BetContent
	err = json.Unmarshal([]byte(bets), &betList)
	if err != nil {
		c.JSON(200, routes.ApiShowResult(common.CodeFail, "下注格式错误"))
		return
	}
	for i := range betList {
		err = service.Bet(&lotto.Order{
			UserID:     uid,
			LottoID:    lottoID,
			Issue:      issue,
			MethodCode: betList[i].MethodCode,
			PlayCode:   betList[i].PlayCode,
			BetContent: betList[i].BetContent,
			Amount:     betList[i].Amount,
			//IP:         ip,
		})
		if err != nil {
			c.JSON(200, routes.ApiShowResult(common.CodeFail, fmt.Sprintf("下注失败：%s", err)))
			return
		}
	}
	c.JSON(200, routes.ApiShowResult(common.CodeOK, "下单成功"))
	return
}

// 游客下单
// 参数：lotto_id，issue，bets
func GuestBet(c *gin.Context) {
	uid, err := routes.GetAPILoginID(c)
	if err != nil {
		c.JSON(200, routes.ApiShowResult(common.CodeFail, "登录校验失败"))
	}

	params := routes.ParamHelper{}

	params.GetPostFormNotEmpty(c, "lotto_id")
	params.GetPostFormNotEmpty(c, "issue")
	params.GetPostFormNotEmpty(c, "bets")

	if params.Exist("lotto_id") == false {
		c.JSON(200, routes.ApiShowResult(common.CodeFail, "彩票不能为空"))
		return
	}
	if params.Exist("issue") == false {
		c.JSON(200, routes.ApiShowResult(common.CodeFail, "期号不能为空"))
		return
	}

	if params.Exist("bets") == false {
		c.JSON(200, routes.ApiShowResult(common.CodeFail, "下注内容不能为空"))
		return
	}

	lottoID := common.GetInt(params.Get("lotto_id"))
	issue := params.Get("issue")
	bets := params.Get("bets")
	//ip := common.InetAton(c.ClientIP())

	var betList []BetContent
	err = json.Unmarshal([]byte(bets), &betList)
	if err != nil {
		c.JSON(200, routes.ApiShowResult(common.CodeFail, "下注格式错误"))
		return
	}
	for i := range betList {
		err = service.GuestBet(&lotto.Order{
			UserID:     uid,
			LottoID:    lottoID,
			Issue:      issue,
			MethodCode: betList[i].MethodCode,
			PlayCode:   betList[i].PlayCode,
			BetContent: betList[i].BetContent,
			Amount:     betList[i].Amount,
			//IP:         ip,
		})
		if err != nil {
			c.JSON(200, routes.ApiShowResult(common.CodeFail, fmt.Sprintf("下注失败：%s", err)))
			return
		}
	}
	c.JSON(200, routes.ApiShowResult(common.CodeOK, "下单成功"))
	return
}

func GetCurrIssueInfo(c *gin.Context) {
	params := routes.ParamHelper{}
	params.GetQueryNotEmpty(c, "lotto_id")
	if params.Exist("lotto_id") == false {
		c.JSON(200, routes.ApiResult(common.CodeFail, "彩票不能为空", map[string]string{}))
		return
	}
	lid, err := common.GetInt2(params.Get("lotto_id"))
	if err != nil {
		c.JSON(200, routes.ApiResult(common.CodeFail, "彩票编号错误", map[string]string{}))
		return
	}
	_, err = lotto.GetLotto(lid)
	if err != nil {
		c.JSON(200, routes.ApiResult(common.CodeFail, "彩票编号错误", map[string]string{}))
		return
	}
	currIssue, err := lotto.GetCurIssue(lid)
	if err != nil {
		c.JSON(200, routes.ApiResult(common.CodeFail, "暂无数据", map[string]string{}))
		return
	} else {
		c.JSON(200, routes.ApiResult(common.CodeOK, "", currIssue))
		return
	}

}
