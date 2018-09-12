package service

import (
	"fmt"
	"github.com/shopspring/decimal"
	"github.com/xiuos/mozi/common"
	"github.com/xiuos/mozi/models"
	"github.com/xiuos/mozi/models/errors"
)

func Bet(o *models.Order) error {
	/*
		检查用户
		检查用户权限
		检查用户余额
		检查房间编号
		检查彩票编号
		检查彩票期号
		检查玩法编号
		检查下注信息
	*/

	u, err := models.GetUserByID(o.UserID)
	if err != nil {
		return err
	}

	if u.Status == models.UserStatusDisable {
		return errors.UserDisableErr{}
	}

	li, err := models.GetLotto(o.LottoID)
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

	uw, err := models.GetUserWallet(o.UserID)
	if err != nil {
		return err
	}
	if o.Amount.GreaterThan(uw.Balance) {
		return errors.New("账户余额不足")
	}

	err = models.MethodCheckBetLegal(o.MethodCode, o)
	if err != nil {
		return err
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
	err = models.CreateRecordLottoOrderTx(tx, o)
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
