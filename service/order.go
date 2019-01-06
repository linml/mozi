package service

import (
	"github.com/shopspring/decimal"
	"github.com/xiuos/mozi/common"
	"github.com/xiuos/mozi/models/errors"
	"github.com/xiuos/mozi/models/lotto"
)

type LottoOrder4User struct {
	ID         int             `json:"id"`
	OrderID    string          `json:"order_id"`
	UserID     int             `json:"user_id"`
	Name       string          `json:"name"`
	Issue      string          `json:"issue"`
	LottoID    int             `json:"lotto_id"`
	LottoName  string          `json:"lotto_name"`
	MethodCode string          `json:"method_code"`
	MethodName string          `json:"method_name"`
	PlayCode   string          `json:"play_code"`
	PlayName   string          `json:"play_name"`
	BetContent string          `json:"bet_content"`
	Odds       decimal.Decimal `json:"odds"`
	Amount     decimal.Decimal `json:"amount"`
	Payout     decimal.Decimal `json:"payout"`
	BetTime    string          `json:"bet_time"`
	Status     int             `json:"status"`
}

func PageFindLottoOrder4UserList(pageParam common.PageParams) (*common.PageResult, error) {
	retpr := common.PageResult{}
	if _, ok := pageParam.Params["user_id"]; !ok {
		return &retpr, errors.New("用户编码错误")
	}
	pr, ol, err := lotto.PageFindLottoOrderList(pageParam)
	if err != nil {
		return &retpr, errors.New("查询异常")
	}
	o4ul := []LottoOrder4User{}
	for i := range *ol {
		o4u := LottoOrder4User{
			ID:         (*ol)[i].ID,
			OrderID:    (*ol)[i].OrderID,
			UserID:     (*ol)[i].UserID,
			Name:       (*ol)[i].Name,
			Issue:      (*ol)[i].Issue,
			LottoID:    (*ol)[i].LottoID,
			MethodCode: (*ol)[i].MethodCode,
			PlayCode:   (*ol)[i].PlayCode,
			BetContent: (*ol)[i].BetContent,
			Odds:       (*ol)[i].Odds,
			Amount:     (*ol)[i].Amount,
			Payout:     (*ol)[i].Payout,
			BetTime:    (*ol)[i].BetTime,
			Status:     (*ol)[i].Status,
		}
		lottoInfo, _ := lotto.GetLotto((*ol)[i].LottoID)
		o4u.LottoName = lottoInfo.Name
		oddsInfo, _ := lotto.GetOdds((*ol)[i].LottoID, (*ol)[i].MethodCode, (*ol)[i].PlayCode)
		o4u.MethodName = oddsInfo.MethodName
		o4u.PlayName = oddsInfo.PlayName

		t, err := common.Str2Time(o4u.BetTime)
		if err == nil {
			o4u.BetTime = common.Time2Str1(t, "2006-01-02 15:04:05")
		}
		o4ul = append(o4ul, o4u)
	}

	pr.RecordData = make([]interface{}, len(o4ul))
	for i, d := range o4ul {
		pr.RecordData[i] = d
	}
	return pr, err
}
