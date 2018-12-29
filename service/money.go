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

type ManualMoney struct {
	OperatorID int
	UserID     int
	Amount     decimal.Decimal
	AuditScore decimal.Decimal
	ChangeType int
	Remark     string
}

func ManualAddMoney(mm ManualMoney) error {
	uInfo, err := models.GetUserByID(mm.UserID)
	if err != nil {
		if err == sql.ErrNoRows {
			return errors.New("账户不存在")
		}
		return err
	}

	if mm.Amount.LessThanOrEqual(MinMoney) == true {
		return errors.New(fmt.Sprintf("金额必须大于系统最金额:%s", MinMoney))
	}
	uWallet, err := models.GetUserWallet(mm.UserID)
	if err != nil {
		return errors.New("获取钱包出错")
	}

	_, err = models.GetCodeChangeMoneyType1(models.ChangeKindDeposit, mm.ChangeType)
	if err != nil {
		return errors.New("流水类型不存在")
	}

	opInfo, err := models.GetAdminUserByID(mm.OperatorID)
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
		ChangeType:    mm.ChangeType,
		BeforeBalance: uWallet.Balance,
		Amount:        mm.Amount,
		AfterBalance:  uWallet.Balance.Add(mm.Amount),
		RecordDate:    common.GetDateNowString(),
		RecordAt:      common.GetTimeNowString(),
		OperateType:   models.OperatorTypeAdmin,
		OperatorID:    opInfo.UserID,
		OperatorName:  opInfo.Name,
		Remark:        mm.Remark,
	}

	tx, err := common.BaseDb.Begin()
	if err != nil {
		return errors.New(fmt.Sprintf("事务开始失败: %s", err))
	}
	err = models.ChangeMoneyTx(tx, mm.UserID, mm.Amount)
	if err != nil {
		tx.Rollback()
		return errors.New(fmt.Sprintf("手动入款异常.%s", err))
	}

	_, err = models.CreateRecordMoneyChangeTx(tx, &rc)
	if err != nil {
		tx.Rollback()
		return errors.New(fmt.Sprintf("手动入款记录异常.%s", err))
	}

	if mm.AuditScore.GreaterThan(MinMoney) == true {
		err = models.ChangeAuditScoreTx(tx, mm.UserID, mm.AuditScore)
		if err != nil {
			tx.Rollback()
			return errors.New(fmt.Sprintf("稽核分记录异常.%s", err))
		}
	}

	err = tx.Commit()
	if err != nil {
		return errors.New(fmt.Sprintf("添加手动入款事务失败! 详情:%s", err))
	}
	return err

}

func ManualSubMoney(mm ManualMoney) error {
	uInfo, err := models.GetUserByID(mm.UserID)
	if err != nil {
		if err == sql.ErrNoRows {
			return errors.New("账户不存在")
		}
		return err
	}

	if mm.Amount.LessThanOrEqual(MinMoney) == true {
		return errors.New(fmt.Sprintf("金额必须大于系统最金额:%s", MinMoney))
	}
	uWallet, err := models.GetUserWallet(mm.UserID)
	if err != nil {
		return errors.New("获取钱包出错")
	}

	_, err = models.GetCodeChangeMoneyType1(models.ChangeKindWithdraw, mm.ChangeType)
	if err != nil {
		return errors.New("流水类型不存在")
	}

	opInfo, err := models.GetAdminUserByID(mm.OperatorID)
	if err != nil {
		if err == sql.ErrNoRows {
			return errors.New("操作用户不存在")
		}
		return err
	}
	mm.Amount = mm.Amount.Neg()
	rc := models.RecordMoneyChange{
		UserID:        uInfo.UserID,
		Name:          uInfo.Name,
		ChangeKind:    models.ChangeKindWithdraw,
		ChangeType:    mm.ChangeType,
		BeforeBalance: uWallet.Balance,
		Amount:        mm.Amount,
		AfterBalance:  uWallet.Balance.Add(mm.Amount),
		RecordDate:    common.GetDateNowString(),
		RecordAt:      common.GetTimeNowString(),
		OperateType:   models.OperatorTypeAdmin,
		OperatorID:    opInfo.UserID,
		OperatorName:  opInfo.Name,
		Remark:        mm.Remark,
	}

	tx, err := common.BaseDb.Begin()
	if err != nil {
		return errors.New(fmt.Sprintf("事务开始失败: %s", err))
	}
	err = models.ChangeMoneyTx(tx, mm.UserID, mm.Amount)
	if err != nil {
		tx.Rollback()
		return errors.New(fmt.Sprintf("手动出款异常.%s", err))
	}

	_, err = models.CreateRecordMoneyChangeTx(tx, &rc)
	if err != nil {
		tx.Rollback()
		return errors.New(fmt.Sprintf("手动出款记录异常.%s", err))
	}

	if mm.AuditScore.Abs().GreaterThan(MinMoney) == true {
		err = models.ChangeAuditScoreTx(tx, mm.UserID, mm.AuditScore)
		if err != nil {
			tx.Rollback()
			return errors.New(fmt.Sprintf("稽核分记录异常.%s", err))
		}
	}

	err = tx.Commit()
	if err != nil {
		return errors.New(fmt.Sprintf("添加手动入款事务失败! 详情:%s", err))
	}
	return err

}
