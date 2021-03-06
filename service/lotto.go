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
		return errors.New("彩票编号错误")
	}

	if li.Status != models.LottoStatusEnable {
		return errors.New("该彩种暂未开放")
	}

	lr, err := lotto.GetLottoResult(o.LottoID, o.Issue)
	if err != nil {
		return errors.New("期号异常")
	}
	if lr.StopTime < common.GetTimeNowString() {
		return errors.New("当期已停止投注,请投注下一期")
	}
	if lr.Status != lotto.LottoResultOpen {
		return errors.New("当期已开奖,请投注下一期")
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

	pInfo, err := lotto.GetOdds(o.LottoID, o.MethodCode, o.PlayCode)
	if err != nil {
		fmt.Println(err)
		return errors.New("找不到该玩法")
	}

	mInfo, err := lotto.GetLottoMethodTemplate(li.LottoType, o.MethodCode)
	if err != nil {
		return errors.New("找不到该玩法[-1]")
	}
	if mInfo.OddsType == 1 {
		if o.PlayCode != o.BetContent {
			return errors.New("下注内容不匹配")
		}
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

	orderID := common.GetTimeNowString()[4:] + common.RandDigitString(8)
	o.LottoType = li.LottoType
	o.Odds = pInfo.Odds
	o.BetCount = betCount
	o.Name = u.Name
	o.OrderID = orderID
	o.GameKind = li.GameKind
	o.GameType = li.GameType
	o.BetTime = common.GetTimeNowString()
	o.UpdateTime = common.GetTimeNowString()
	o.BetDate = common.GetDateNowString()

	rc := models.RecordMoneyChange{
		UserID:        u.UserID,
		LinkID:        orderID,
		Name:          u.Name,
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

	err = models.ChangeMoneyTx(tx, o.UserID, o.Amount.Abs().Neg())
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

// 游客下单
func GuestBet(o *lotto.Order) error {
	/*
		检查用户
		检查彩票
		检查彩票期号
		检查用户权限
		检查玩法编号
		检查用户余额

		检查下注信息
	*/

	u, err := models.GetGuestUser(o.UserID)
	if err != nil {
		return err
	}

	li, err := lotto.GetLotto(o.LottoID)
	if err != nil {
		return errors.New("彩票编号错误")
	}

	if li.Status != models.LottoStatusEnable {
		return errors.New("该彩种暂未开放")
	}

	lr, err := lotto.GetLottoResult(o.LottoID, o.Issue)
	if err != nil {
		return errors.New("期号异常")
	}
	if lr.StopTime < common.GetTimeNowString() {
		return errors.New("当期已停止投注,请投注下一期")
	}
	if lr.Status != lotto.LottoResultOpen {
		return errors.New("当期已开奖,请投注下一期")
	}

	err = models.MethodCheckBetLegal(o.MethodCode, o)
	if err != nil {
		return err
	}

	pInfo, err := lotto.GetOdds(o.LottoID, o.MethodCode, o.PlayCode)
	if err != nil {
		fmt.Println(err)
		return errors.New("找不到该玩法")
	}

	mInfo, err := lotto.GetLottoMethodTemplate(li.LottoType, o.MethodCode)
	if err != nil {
		return errors.New("找不到该玩法[-1]")
	}
	if mInfo.OddsType == 1 {
		if o.PlayCode != o.BetContent {
			return errors.New("下注内容不匹配")
		}
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

	if o.Amount.GreaterThan(u.Balance) {
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

	orderID := common.GetTimeNowString()[4:] + common.RandDigitString(8)
	o.LottoType = li.LottoType
	o.Odds = pInfo.Odds
	o.BetCount = betCount
	o.Name = u.Name
	o.OrderID = orderID
	o.GameKind = li.GameKind
	o.GameType = li.GameType
	o.BetTime = common.GetTimeNowString()
	o.UpdateTime = common.GetTimeNowString()
	o.BetDate = common.GetDateNowString()

	tx, err := common.BaseDb.Begin()
	if err != nil {
		return errors.New(fmt.Sprintf("事务开始失败: %s", err))
	}
	err = lotto.CreateGuestRecordLottoOrderTx(tx, o)
	if err != nil {
		tx.Rollback()
		return errors.New(fmt.Sprintf("下单异常.%s", err))
	}

	err = models.GuestChangeMoneyTx(tx, o.UserID, o.Amount.Abs().Neg())
	if err != nil {
		tx.Rollback()
		return errors.New(fmt.Sprintf("扣款异常.%s", err))
	}

	err = tx.Commit()
	if err != nil {
		return errors.New(fmt.Sprintf("添加下单事务失败! 详情:%s", err))
	}
	return nil
}

// 期号生成
func CreateIssueCorn() {

	facList, err := lotto.FindIssueFactory(map[string]string{})
	if err != nil {
		fmt.Println(err)
		return
	}

	d, _ := time.ParseDuration("24h")
	for _, f := range *facList {

		if f.Status != models.LottoStatusEnable {
			continue
		}

		_, err := common.BaseDb.Exec("CALL CreateLottoIssueTable(?)", f.LottoID)
		if err != nil {
			fmt.Println(err)
		}

		day := time.Now()

		for i := 0; i < 3; i++ {
			fmt.Println("GEN ISSUE >>> lotto_id:", f.LottoID, "day:", day.Format("20060102"))
			issueInfoList, err := f.GenIssue(day)
			if err != nil {
				fmt.Println(err)
				continue
			}
			lotto.CreateIssueInfo(f.LottoID, issueInfoList)
			day = day.Add(d)
		}
	}
}

// 	赔率生成
func CheckAndCreateOdds() {
	lottoInfoList, err := lotto.FindCodeLottoList(map[string]string{})
	if err != nil {
		fmt.Println(err)
		return
	}
	lottoMethodTmpList, err := lotto.FindLottoMethodTemplateList(map[string]string{})
	if err != nil {
		fmt.Println(err)
		return
	}

	lottoOddTmpList, err := lotto.FindLottoOddsTemplateList(map[string]string{})
	if err != nil {
		fmt.Println(err)
		return
	}

	for i := range *lottoInfoList {
		for j := range *lottoMethodTmpList {
			if (*lottoInfoList)[i].LottoType != (*lottoMethodTmpList)[j].LottoType {
				continue
			}
			for k := range *lottoOddTmpList {
				if (*lottoMethodTmpList)[j].MethodCode == (*lottoOddTmpList)[k].MethodCode {
					lo := lotto.LottoOdds{
						LottoID:    (*lottoInfoList)[i].LottoID,
						MethodCode: (*lottoMethodTmpList)[j].MethodCode,
						PlayCode:   (*lottoOddTmpList)[k].PlayCode,
						MethodName: (*lottoOddTmpList)[k].MethodName,
						PlayName:   (*lottoOddTmpList)[k].PlayName,
						BaseOdds:   (*lottoOddTmpList)[k].BaseOdds,
						Odds:       (*lottoOddTmpList)[k].Odds,
						OddsMin:    (*lottoOddTmpList)[k].OddsMin,
						OddsMax:    (*lottoOddTmpList)[k].OddsMax,
						BetMin:     (*lottoOddTmpList)[k].BetMin,
						BetMax:     (*lottoOddTmpList)[k].BetMax,
						Status:     (*lottoOddTmpList)[k].Status,
						SortIndex:  (*lottoOddTmpList)[k].SortIndex,
						IsShow:     (*lottoOddTmpList)[k].IsShow,
					}
					err := lotto.CreateLottoOdds(lo)
					if err != nil {
						fmt.Println(err)
					}
				}
			}
		}
	}
}

// 	玩法组生成
func CheckAndCreateCMSLottoMethodGroup() {
	lottoInfoList, err := lotto.FindCodeLottoList(map[string]string{})
	if err != nil {
		fmt.Println(err)
		return
	}
	lmgList, err := lotto.FindCMSLottoMethodGroupTemplateList(map[string]string{})
	if err != nil {
		fmt.Println(err)
		return
	}

	lmgpList, err := lotto.FindCMSLottoMethodGroupPlayTemplateList(map[string]string{})
	if err != nil {
		fmt.Println(err)
		return
	}

	for i := range *lottoInfoList {
		for j := range *lmgList {
			if (*lottoInfoList)[i].LottoType != (*lmgList)[j].LottoType {
				continue
			}
			g := lotto.CMSLottoMethodGroup{
				LottoID:    (*lottoInfoList)[i].LottoID,
				GroupID:    (*lmgList)[j].GroupID,
				GroupName:  (*lmgList)[j].GroupName,
				GroupAlias: (*lmgList)[j].GroupAlias,
				SortIndex:  (*lmgList)[j].SortIndex,
				Status:     (*lmgList)[j].Status,
			}
			err := lotto.CreateLottoMethodGroup(g)
			if err != nil {
				fmt.Println(err)
			}
			for k := range *lmgpList {
				if (*lmgList)[j].GroupID == (*lmgpList)[k].GroupID {
					gp := lotto.CMSLottoMethodGroupPlay{
						LottoID:     (*lottoInfoList)[i].LottoID,
						GroupID:     (*lmgList)[j].GroupID,
						GroupName:   (*lmgList)[j].GroupName,
						GroupAlias:  (*lmgList)[j].GroupAlias,
						MethodCode:  (*lmgpList)[k].MethodCode,
						MethodName:  (*lmgpList)[k].MethodName,
						MethodAlias: (*lmgpList)[k].MethodAlias,
						SortIndex:   (*lmgpList)[k].SortIndex,
						Status:      (*lmgpList)[k].Status,
					}
					err := lotto.CreateLottoMethodGroupPlay(gp)
					if err != nil {
						fmt.Println(err)
					}
				}
			}
		}
	}
}

// 设置彩票状态
func SetCodeLottoStatus(lid int, status int) error {
	return lotto.SetCodeLottoInfo(lid, "status", status)
}

// 设置彩票排序
func SetCodeLottoSortIndex(lid int, sortIndex int) error {
	return lotto.SetCodeLottoInfo(lid, "sort_index", sortIndex)
}

// 设置彩票是否显示
func SetCodeLottoIsShow(lid int, isShow int) error {
	return lotto.SetCodeLottoInfo(lid, "is_show", isShow)
}

// 设置彩票赔率赔率
func SetLottoOdds(lid int, mCode string, pCode string, odds decimal.Decimal) error {
	return lotto.SetLottoOddsInfo(lid, mCode, pCode, "odds", odds)
}

// 设置彩票赔率最小赔率
func SetLottoOddsMin(lid int, mCode string, pCode string, oddsMin decimal.Decimal) error {
	return lotto.SetLottoOddsInfo(lid, mCode, pCode, "odds_min", oddsMin)
}

// 设置彩票赔率最大赔率
func SetLottoOddsMax(lid int, mCode string, pCode string, oddsMax decimal.Decimal) error {
	return lotto.SetLottoOddsInfo(lid, mCode, pCode, "odds_max", oddsMax)
}

// 设置彩票赔率最小额度
func SetLottoOddsBetMin(lid int, mCode string, pCode string, betMin decimal.Decimal) error {
	return lotto.SetLottoOddsInfo(lid, mCode, pCode, "bet_min", betMin)
}

// 设置彩票赔率最大额度
func SetLottoOddsBetMax(lid int, mCode string, pCode string, betMax decimal.Decimal) error {
	return lotto.SetLottoOddsInfo(lid, mCode, pCode, "bet_max", betMax)
}

// 设置彩票赔率玩法项状态
func SetLottoOddsStatus(lid int, mCode string, pCode string, status int) error {
	return lotto.SetLottoOddsInfo(lid, mCode, pCode, "status", status)
}

// 设置彩票赔率玩法项是否显示
func SetLottoOddsIsShow(lid int, mCode string, pCode string, isShow int) error {
	return lotto.SetLottoOddsInfo(lid, mCode, pCode, "is_show", isShow)
}

func CallDrawLotto(lottoID int, issue string, number string) error {
	err := lotto.CheckDrawNumberLegal(lottoID, number)
	if err != nil {
		return err
	}

	issueInfo, err := lotto.GetLottoResult(lottoID, issue)
	if err != nil {
		return err
	}

	if issueInfo.Status == lotto.LottoResultOpened {
		return errors.New("本期已开奖")
	} else {
		lotto.SettleLottoResult(lottoID, issue, number)
	}
	err = models.CalcIssue(lottoID, issue, number)
	return err
}
