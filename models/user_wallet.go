package models

import (
	"database/sql"
	"fmt"
	"github.com/shopspring/decimal"
	"github.com/xiuos/mozi/common"
	"github.com/xiuos/mozi/xorm"
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

func (u *UserWallet) TableName() string {
	return "user_wallet"
}

func (u *UserWallet) Field() []string {
	return []string{"user_id", "password", "balance", "status"}
}

func (u *UserWallet) FieldItem() []interface{} {
	return []interface{}{&u.UserID, &u.Password, &u.Balance, &u.Status}
}

func CreateUserWalletTx(tx *sql.Tx, uw *UserWallet) error {
	o := xorm.Orm{}
	createSql, err := o.Table(uw.TableName()).Create().Do(uw)
	if err != nil {
		return err
	}
	_, err = tx.Exec(createSql, uw.FieldItem()...)
	return err
}

func GetUserWallet(uid int) (*UserWallet, error) {
	o := xorm.Orm{}
	uw := UserWallet{}
	querySql, err := o.Table(uw.TableName()).Query().Where("user_id = ?", uid).Do(&uw)
	if err != nil {
		return &uw, err
	}
	err = common.BaseDb.QueryRow(querySql, o.Args()...).Scan(uw.FieldItem()...)
	return &uw, err
}

func SetWalletPassword(uid int, password string) error {
	hashPassword, err := common.HashPassword(password)
	if err != nil {
		return err
	}
	updateSql := fmt.Sprintf("UPDATE %s SET password = ? WHERE user_id = ?", TableUserWallet)
	_, err = common.BaseDb.Exec(updateSql, hashPassword, uid)
	return err
}

func SetWalletStatus(uid int, status int) error {

	updateSql := fmt.Sprintf("UPDATE %s SET status = ? WHERE user_id = ?", TableUserWallet)
	_, err := common.BaseDb.Exec(updateSql, status, uid)
	return err
}

// 余额变动
// money为正数表示加钱，负数表示减钱
func ChangeMoneyTx(tx *sql.Tx, uid int, money decimal.Decimal) error {
	updateSql := fmt.Sprintf("UPDATE %s SET balance=balance+? WHERE user_id = ?", TableUserWallet)
	_, err := tx.Exec(updateSql, money, uid)

	return err
}
