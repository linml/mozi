package service

import (
	"fmt"
	"github.com/shopspring/decimal"
	"github.com/xiuos/mozi/common"
	"github.com/xiuos/mozi/models"
	"github.com/xiuos/mozi/models/errors"
	"github.com/xiuos/mozi/models/lotto"
	"time"
)

// 下单
func Bet(o *lotto.Order) error {
	/*
		检查用户
		检查彩票
		检查彩票期号
		检查用户权限
		检查玩法编号
		检查用户余额

		检查下注信息
	*/

	u, err := models.GetUserByID(o.UserID)
	if err != nil {
		return err
	}

	if u.Status == models.UserStatusDisable {
		return errors.UserDisableErr{}
	}

	li, err := lotto.GetLotto(o.LottoID)
	if err != nil {
		fmt.Println(err)
		return errors.New("彩票编号错误")
	}

	if li.Status != models.LottoStatusEnable {
		return errors.New("该彩种暂未开放")
	}

	lr, err := models.GetLottoResult(o.LottoID, o.Issue)
	if err != nil {
		return errors.New("期号异常")
	}
	if lr.CloseTime < common.GetTimeNowString() {
		return errors.New("当期已停止投注,请投注下一期")
	}

	uPower, err := models.GetUserPower(o.UserID)
	if err != nil {
		return err
	}
	if uPower.PowerBet != models.PowerEnable {
		return errors.New("账户无下注权限")
	}

	err = models.MethodCheckBetLegal(o.MethodCode, o)
	if err != nil {
		return err
	}

	pInfo, err := models.GetPlayInfo(o.LottoID, o.MethodCode, o.Content)
	if err != nil {
		return err
	}
	if pInfo.Status != 1 {
		return errors.New("该玩法已停售")
	}

	if pInfo.BetMin.GreaterThan(o.Amount) {
		return errors.New("投注金额过低")
	}

	if pInfo.BetMax.LessThan(o.Amount) {
		return errors.New("投注金额过高")
	}

	uw, err := models.GetUserWallet(o.UserID)
	if err != nil {
		return err
	}

	if o.Amount.GreaterThan(uw.Balance) {
		return errors.New("账户余额不足")
	}

	betCount, err := models.MethodBetCount(o.MethodCode, o)
	if err != nil {
		return err
	}
	if betCount < 1 {
		return errors.New("下注注数异常")
	}
	minS, _ := decimal.NewFromString("0.001")
	bcD, _ := decimal.NewFromString(fmt.Sprintf("%d", betCount))
	if o.Amount.Div(bcD).LessThan(minS) {
		return errors.New(fmt.Sprintf("单注最低金额不能低于%s", minS))
	}

	o.Odds = pInfo.Odds
	o.Count = betCount
	o.Username = u.Name
	o.OrderNo = common.GetTimeNowString() + common.RandString(6)
	o.GameKind = li.GameKind
	o.GameType = li.GameType
	o.BetTime = common.GetTimeNowString()
	o.BetDate = common.GetDateNowString()

	rc := models.RecordMoneyChange{
		UserID:        u.UserID,
		Username:      u.Name,
		GameKind:      li.GameKind,
		GameType:      li.GameType,
		ChangeKind:    models.ChangeKindLotto,
		ChangeType:    models.ChangeTypeLottoBet,
		BeforeBalance: uw.Balance,
		Amount:        o.Amount.Abs().Neg(),
		AfterBalance:  uw.Balance.Sub(o.Amount.Abs()),
		RecordDate:    common.GetDateNowString(),
		RecordAt:      common.GetTimeNowString(),
		OperateType:   models.OperatorTypeSelf,
		Remark:        fmt.Sprintf("%s-%s期下单", li.Name, o.Issue),
	}

	tx, err := common.BaseDb.Begin()
	if err != nil {
		return errors.New(fmt.Sprintf("事务开始失败: %s", err))
	}
	err = lotto.CreateRecordLottoOrderTx(tx, o)
	if err != nil {
		tx.Rollback()
		return errors.New(fmt.Sprintf("下单异常.%s", err))
	}

	err = models.ChangeMoneyTx(tx, o.UserID, o.Amount.Abs())
	if err != nil {
		tx.Rollback()
		return errors.New(fmt.Sprintf("扣款异常.%s", err))
	}

	_, err = models.CreateRecordMoneyChangeTx(tx, &rc)
	if err != nil {
		tx.Rollback()
		return errors.New(fmt.Sprintf("扣款记录异常.%s", err))
	}
	err = tx.Commit()
	if err != nil {
		return errors.New(fmt.Sprintf("添加下单事务失败! 详情:%s", err))
	}
	return nil
}

// 期号生成
func CreateIssueCorn() {

	facList, err := models.FindIssueFactory(map[string]string{})
	if err != nil {
		fmt.Println(err)
		return
	}

	d, _ := time.ParseDuration("24h")
	for _, f := range *facList {

		if f.Status != models.LottoStatusEnable {
			continue
		}

		day := time.Now()

		for i := 0; i < 3; i++ {
			fmt.Println("GEN ISSUE >>> lotto_id:", f.LottoID, "day:", day.Format("20060102"))
			issueInfoList, err := f.GenIssue(day)
			if err != nil {
				fmt.Println(err)
				continue
			}
			models.CreateIssueInfo(issueInfoList)
			day = day.Add(d)
		}
	}
}

func SetCodeLottoStatus(lid int, status int) error {
	return lotto.SetCodeLottoInfo(lid, "status", status)
}

func SetCodeLottoSortIndex(lid int, sortIndex int) error {
	return lotto.SetCodeLottoInfo(lid, "sort_index", sortIndex)
}

func SetCodeLottoIsShow(lid int, isShow int) error {
	return lotto.SetCodeLottoInfo(lid, "is_show", isShow)
}
