package models

import (
	"database/sql"
	"github.com/xiuos/mozi/common"
	"github.com/xiuos/xorm"
	"github.com/shopspring/decimal"
)

const (
	TableUserWallet = "user_wallet"
)

const (
	WalletStatusUnActivate = 0 // 钱包状态未激活
	WalletStatusEnable     = 1 // 钱包状态启用
	WalletStatusDisable    = 2 // 钱包状态禁用
	WalletStatusMustRest   = 3 // 钱包状态需要重置

)

type UserWallet struct {
	UserID   int
	Password string
	Balance  decimal.Decimal
	Status   int
}

func (u *UserWallet) Field() []string {
	return []string{"user_id", "password", "balance", "status"}
}

func (u *UserWallet) FieldItem() []interface{} {
	return []interface{}{&u.UserID, &u.Password, &u.Balance, &u.Status}
}

func CreateUserWalletTx(tx *sql.Tx, uw *UserWallet) error {
	o := xorm.Orm{}
	createSql, err := o.Table(TableUserWallet).Create().Do(uw)
	if err != nil {
		return err
	}
	_, err = tx.Exec(createSql, uw.FieldItem()...)
	return err
}

func GetUserWallet(uid int) (*UserWallet, error) {
	o := xorm.Orm{}
	uw := UserWallet{}
	querySql, err := o.Table(TableUserWallet).Query().Where("user_id = ?", uid).Do(&uw)
	if err != nil {
		return &uw, err
	}
	err = common.BaseDb.QueryRow(querySql, o.Args()...).Scan(uw.FieldItem()...)
	return &uw, err
}
