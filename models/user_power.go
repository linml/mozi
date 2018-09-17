package models

import (
	"database/sql"
	"github.com/xiuos/mozi/common"
	"github.com/xiuos/xorm"
)

type UserPower struct {
	UserID        int `json:"user_id"`
	PowerBet      int `json:"power_bet"`
	PowerLogin    int `json:"power_login"`
	PowerDeposit  int `json:"power_deposit"`
	PowerWithdraw int `json:"power_withdraw"`
}

func (up *UserPower) TableName() string {
	return "user_power"
}

func (up *UserPower) Field() []string {
	return []string{"user_id", "power_bet", "power_login", "power_deposit", "power_withdraw"}
}

func (up *UserPower) FieldItem() []interface{} {
	return []interface{}{&up.UserID, &up.PowerBet, &up.PowerLogin, &up.PowerDeposit, &up.PowerWithdraw}
}

func CreateUserPowerTx(tx *sql.Tx, up *UserPower) error {
	createSql := "INSERT INTO " + up.TableName() + " SET  user_id = ?, power_bet = ?, power_login = ?, power_deposit = ?, power_withdraw = ?"
	_, err := tx.Exec(createSql, &up.UserID, &up.PowerBet, &up.PowerLogin, &up.PowerDeposit, &up.PowerWithdraw)
	return err
}

func GetUserPower(uid int) (*UserPower, error) {
	o := xorm.Orm{}
	up := UserPower{}
	querySql, err := o.Table(up.TableName()).Query().Where("user_id = ?", uid).Do(&up)
	if err != nil {
		return &up, err
	}
	err = common.BaseDb.QueryRow(querySql, o.Args()...).Scan(up.FieldItem()...)
	return &up, err
}
