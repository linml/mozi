package service

import (
	"database/sql"
	"fmt"
	"github.com/shopspring/decimal"
	"github.com/xiuos/mozi/common"
	"github.com/xiuos/mozi/models"
	"github.com/xiuos/mozi/models/errors"
)

var MinMoney, _ = decimal.NewFromString("0.01")

func ManualAddMoney(opId int, uid int, amount decimal.Decimal, changeType int) error {
	uInfo, err := models.GetUserByID(uid)
	if err != nil {
		if err == sql.ErrNoRows {
			return errors.New("账户不存在")
		}
		return err
	}

	if amount.LessThanOrEqual(decimal.NewFromFloat(0.0)) == true {
		return errors.New("金额必须大于0")
	}
	if amount.LessThanOrEqual(MinMoney) == true {
		return errors.New(fmt.Sprintf("金额必须大于系统最金额:%s", MinMoney))
	}
	uWallet, err := models.GetUserWallet(uid)
	if err != nil {
		return errors.New("获取钱包出错")
	}

	cInfo, err := models.GetCodeChangeMoneyType(changeType)
	if err != nil {
		return errors.New("流水类型不存在")
	}

	opInfo, err := models.GetAdminUserByID(opId)
	if err != nil {
		if err == sql.ErrNoRows {
			return errors.New("操作用户不存在")
		}
		return err
	}
	rc := models.RecordMoneyChange{
		UserID:        uInfo.UserID,
		Name:          uInfo.Name,
		ChangeKind:    models.ChangeKindDeposit,
		ChangeType:    changeType,
		BeforeBalance: uWallet.Balance,
		Amount:        amount,
		AfterBalance:  uWallet.Balance.Sub(amount),
		RecordDate:    common.GetDateNowString(),
		RecordAt:      common.GetTimeNowString(),
		OperateType:   models.OperatorTypeAdmin,
		OperatorID:    opInfo.UserID,
		OperatorName:  opInfo.Name,
		Remark:        fmt.Sprintf("手动存入账户:%s,类型:%s,金额:%s", uInfo.Name, cInfo.Name, amount),
	}

	tx, err := common.BaseDb.Begin()
	if err != nil {
		return errors.New(fmt.Sprintf("事务开始失败: %s", err))
	}
	err = models.ChangeMoneyTx(tx, uid, amount)
	if err != nil {
		tx.Rollback()
		return errors.New(fmt.Sprintf("手动入款异常.%s", err))
	}

	_, err = models.CreateRecordMoneyChangeTx(tx, &rc)
	if err != nil {
		tx.Rollback()
		return errors.New(fmt.Sprintf("手动入款记录异常.%s", err))
	}
	err = tx.Commit()
	if err != nil {
		return errors.New(fmt.Sprintf("添加手动入款事务失败! 详情:%s", err))
	}
	return err

}
