package models

import (
	"errors"
	"fmt"
	"github.com/shopspring/decimal"
	"github.com/xiuos/mozi/common"
	"github.com/xiuos/mozi/models/lotto"
	"strings"
)

func CalcIssue(lottoID int, issue string, number string) error {
	err := lotto.CheckDrawNumberLegal(lottoID, number)
	if err != nil {
		return err
	}

	drawNumberStrList := strings.Split(number, ",")

	lastID := 0
	for {
		orderList, err := lotto.FindOrder4Settle(lastID, lottoID, issue)
		if err != nil {
			fmt.Println(err)
			break
		}
		fmt.Println(len(*orderList))

		if len(*orderList) == 0 {
			break
		}

		for i := range *orderList {
			err = CalcOrder((*orderList)[i], drawNumberStrList)
			if (*orderList)[i].ID > lastID {
				lastID = (*orderList)[i].ID
			}
			if err != nil {
				common.Logger.Warn(fmt.Sprintf("%s", err))
			}
		}

	}
	return nil
}

func CalcOrder(order lotto.Order, drawNumber []string) error {
	// 检查投注内容
	// 检查投注注数
	// 检查中奖信息
	// 获取中奖金额和赔率
	// 设置订单信息
	// 中奖用户添加流水

	engine, err := lotto.GetMethodEngine(order.MethodCode)
	if err != nil {
		return err
	}
	err = engine.CheckBetLegal(&order)

	if err != nil {
		lotto.SetLottoOrderInfo(order.ID, "status", lotto.OrderStatusAbnormal)
		return err
	}

	realBetCount, err := engine.GetBetCount(&order)

	if err != nil {
		lotto.SetLottoOrderInfo(order.ID, "status", lotto.OrderStatusAbnormal)
		return err
	}

	if realBetCount <= 0 {
		lotto.SetLottoOrderInfo(order.ID, "status", lotto.OrderStatusAbnormal)
		return errors.New("下单数量错误")
	}

	winInfoList, err := engine.FindWin(&order, drawNumber)

	if err != nil {
		lotto.SetLottoOrderInfo(order.ID, "status", lotto.OrderStatusAbnormal)
		return err
	}

	TotalWin, _ := decimal.NewFromString("0")
	betCount, _ := decimal.NewFromString(fmt.Sprintf("%d", realBetCount))
	singleMoney := order.Amount.Div(betCount)
	if len(*winInfoList) > 0 {
		winContent := ""
		winCount := 0

		for _, winInfo := range *winInfoList {
			winContent += winInfo.Info
			winCount += winInfo.Count
			c, _ := decimal.NewFromString(fmt.Sprintf("%d", winInfo.Count))
			TotalWin = TotalWin.Add(singleMoney.Mul(winInfo.Odds).Mul(c))
		}

		order.WinContent = winContent
		order.WinCount = winCount
		order.DrawNumber = strings.Join(drawNumber, ",")
		order.Status = lotto.OrderStatusWinning
		order.Payout = TotalWin
		order.Profit = TotalWin.Sub(order.Amount)

	} else {
		order.WinContent = ""
		order.WinCount = 0
		order.DrawNumber = strings.Join(drawNumber, ",")
		order.Status = lotto.OrderStatusNotWinning
		order.Payout = TotalWin
		order.Profit = TotalWin.Sub(order.Amount)
	}

	order.Flag = 1
	order.CalcDate = common.GetDateNowString()
	order.UpdateTime = common.GetTimeNowString()

	tx, err := common.BaseDb.Begin()
	if err != nil {
		fmt.Println(err)
		return err
	}

	err = lotto.SettleUpdateLottoOrderTx(tx, &order)
	if err != nil {
		tx.Rollback()
		return err
	}

	if order.Status == lotto.OrderStatusWinning {

		userWallet, err := GetUserWallet(order.UserID)
		if err != nil {
			return errors.New("获取余额失败")
		}
		lottoInfo, err := lotto.GetLotto(order.LottoID)
		if err != nil {
			return errors.New("彩票编号错误")
		}

		rc := RecordMoneyChange{
			UserID:        order.UserID,
			LinkID:        fmt.Sprintf("%d", order.ID),
			Name:          order.Name,
			GameKind:      lottoInfo.GameKind,
			GameType:      lottoInfo.GameType,
			ChangeKind:    ChangeKindLotto,
			ChangeType:    ChangeTypeLottoPayout,
			BeforeBalance: userWallet.Balance,
			Amount:        order.Payout,
			AfterBalance:  userWallet.Balance.Add(order.Payout),
			RecordDate:    common.GetDateNowString(),
			RecordAt:      common.GetTimeNowString(),
			OperateType:   OperatorTypeSelf,
			Remark:        fmt.Sprintf("%s-%s期派奖", lottoInfo.Name, order.Issue),
		}

		_, err = CreateRecordMoneyChangeTx(tx, &rc)
		if err != nil {
			tx.Rollback()
			return errors.New(fmt.Sprintf("派彩记录异常.%s", err))
		}

		err = ChangeMoneyTx(tx, order.UserID, order.Payout)
		if err != nil {
			tx.Rollback()
			return errors.New(fmt.Sprintf("派奖记录异常.%s", err))
		}
	}
	err = tx.Commit()
	if err != nil {
		return errors.New(fmt.Sprintf("添加派奖事务失败! 详情:%s", err))
	}
	return nil
}
