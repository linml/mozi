package models

import (
	"database/sql"
	"fmt"
	"github.com/shopspring/decimal"
	"github.com/xiuos/mozi/common"
	"github.com/xiuos/mozi/xorm"
)

const (
	TableUserScore = "user_score"
)

type UserScore struct {
	UserID     int
	GameScore  decimal.Decimal
	AuditScore decimal.Decimal
	UpScore    decimal.Decimal
}

func (u *UserScore) TableName() string {
	return "user_score"
}

func (u *UserScore) Field() []string {
	return []string{"user_id", "game_score", "audit_score", "up_score"}
}

func (u *UserScore) FieldItem() []interface{} {
	return []interface{}{&u.UserID, &u.GameScore, &u.AuditScore, &u.UpScore}
}

func CreateUserScoreTx(tx *sql.Tx, us *UserScore) error {
	o := xorm.Orm{}
	createSql, err := o.Table(us.TableName()).Create().Do(us)
	if err != nil {
		return err
	}
	_, err = tx.Exec(createSql, us.FieldItem()...)
	return err
}

func GetUserScore(uid int) (*UserScore, error) {
	o := xorm.Orm{}
	us := UserScore{}
	querySql, err := o.Table(us.TableName()).Query().Where("user_id = ?", uid).Do(&us)
	if err != nil {
		return &us, err
	}
	err = common.BaseDb.QueryRow(querySql, o.Args()...).Scan(us.FieldItem()...)
	return &us, err
}

// score为正数表示加钱，负数表示减钱
func ChangeAuditScoreTx(tx *sql.Tx, uid int, score decimal.Decimal) error {
	s := UserScore{}
	updateSql := fmt.Sprintf("UPDATE %s SET audit_score=audit_score+? WHERE user_id = ?", s.TableName())
	_, err := tx.Exec(updateSql, score, uid)

	return err
}
