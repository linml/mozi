package models

import (
	"database/sql"
	"fmt"
	"github.com/shopspring/decimal"
	"github.com/xiuos/mozi/common"
	"strings"
)

type GuestUser struct {
	UserID   int             `json:"user_id"`
	Name     string          `json:"name"`
	Balance  decimal.Decimal `json:"balance"`
	CreateAt string          `json:"create_at"`
}

func (u *GuestUser) TableName() string {
	return "guest_user"
}

func (u *GuestUser) Field() []string {
	return []string{"user_id", "name", "balance", "create_at"}
}

func (u *GuestUser) FieldItem() []interface{} {
	return []interface{}{&u.UserID, &u.Name, &u.Balance, &u.CreateAt}
}

func GenGuestUser() (*GuestUser, error) {
	gu := GuestUser{}
	name := fmt.Sprintf("guest%s", common.RandDigitString(6))
	insertSql := fmt.Sprintf("INSERT INTO %s SET name=?,balance=?,create_at=?", gu.TableName())
	_, err := common.BaseDb.Exec(insertSql, name, 2000, common.GetTimeNowString())
	if err != nil {
		fmt.Println(err)
	}
	querySql := fmt.Sprintf("SELECT %s FROM %s WHERE name=?", strings.Join(gu.Field(), ","), gu.TableName())
	err = common.BaseDb.QueryRow(querySql, name).Scan(gu.FieldItem()...)
	return &gu, err
}

func GetGuestUser(uid int) (*GuestUser, error) {
	gu := GuestUser{}
	querySql := fmt.Sprintf("SELECT %s FROM %s WHERE user_id=?", strings.Join(gu.Field(), ","), gu.TableName())
	err := common.BaseDb.QueryRow(querySql, uid).Scan(gu.FieldItem()...)
	return &gu, err
}

// 余额变动
// money为正数表示加钱，负数表示减钱
func GuestChangeMoneyTx(tx *sql.Tx, uid int, money decimal.Decimal) error {
	gu := GuestUser{}
	updateSql := fmt.Sprintf("UPDATE %s SET balance=balance+? WHERE user_id = ?", gu.TableName())
	_, err := tx.Exec(updateSql, money, uid)

	return err
}
