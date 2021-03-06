package models

import (
	"database/sql"
	"github.com/xiuos/mozi/common"
	"github.com/xiuos/mozi/xorm"
)

type UserRelation struct {
	UserID   int    `json:"user_id"`
	ParentID int    `json:"parent_id"`
	Parents  string `json:"parents"`
	UserType int    `json:"user_type"`
}

func (ur *UserRelation) TableName() string {
	return "user_relation"
}

func (ur *UserRelation) Field() []string {
	return []string{"user_id", "parent_id", "parents", "user_type"}
}

func (ur *UserRelation) FieldItem() []interface{} {
	return []interface{}{&ur.UserID, &ur.ParentID, &ur.Parents, &ur.UserType}
}

func CreateUserRelationTx(tx *sql.Tx, ur *UserRelation) error {
	o := xorm.Orm{}
	createSql, err := o.Table(ur.TableName()).Create().Do(ur)
	if err != nil {
		return err
	}
	_, err = tx.Exec(createSql, ur.FieldItem()...)
	return err
}

func GetUserRelation(uid int) (*UserRelation, error) {
	o := xorm.Orm{}
	ur := UserRelation{}
	querySql, err := o.Table(ur.TableName()).Query().Where("user_id = ?", uid).Do(&ur)
	if err != nil {
		return &ur, err
	}
	err = common.BaseDb.QueryRow(querySql, o.Args()...).Scan(ur.FieldItem()...)
	return &ur, err
}
