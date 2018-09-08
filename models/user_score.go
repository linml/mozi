package models

import (
	"database/sql"
	"github.com/xiuos/mozi/common"
	"github.com/xiuos/xorm"
	"github.com/shopspring/decimal"
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

func (u *UserScore) Field() []string {
	return []string{"user_id", "game_score", "audit_score", "up_score"}
}

func (u *UserScore) FieldItem() []interface{} {
	return []interface{}{&u.UserID, &u.GameScore, &u.AuditScore, &u.UpScore}
}

func CreateUserScoreTx(tx *sql.Tx, uw *UserScore) error {
	o := xorm.Orm{}
	createSql, err := o.Table(TableUserScore).Create().Do(uw)
	if err != nil {
		return err
	}
	_, err = tx.Exec(createSql, uw.FieldItem()...)
	return err
}

func GetUserScore(uid int) (*UserScore, error) {
	o := xorm.Orm{}
	us := UserScore{}
	querySql, err := o.Table(TableUserScore).Query().Where("user_id = ?", uid).Do(&us)
	if err != nil {
		return &us, err
	}
	err = common.BaseDb.QueryRow(querySql, o.Args()...).Scan(us.FieldItem()...)
	return &us, err
}
