package models

import (
	"database/sql"
	"fmt"
	"github.com/xiuos/mozi/common"
	"github.com/xiuos/xorm"
)

type UserLinks struct {
	ID       int    `json:"id"`
	UserID   int    `json:"user_id"`
	Ref      string `json:"ref"`
	UserType int    `json:"user_type"`
}

func (ul *UserLinks) TableName() string {
	return "user_links"
}

func (ul *UserLinks) Field() []string {
	return []string{"id", "user_id", "ref", "user_type"}
}

func (ul *UserLinks) FieldItem() []interface{} {
	return []interface{}{&ul.ID, &ul.UserID, &ul.Ref, &ul.UserType}
}

func CreateUserLinksTx(tx *sql.Tx, ul *UserLinks) error {
	createSql := fmt.Sprintf("INSERT INTO %s SET user_id=?, ref=?, user_type=?", ul.TableName())
	_, err := tx.Exec(createSql, ul.UserID, ul.Ref, ul.UserType)
	return err
}

func GetUserLinkByRef(ref string) (*UserLinks, error) {
	o := xorm.Orm{}
	ul := UserLinks{}
	querySql, err := o.Table(ul.TableName()).Query().Where("ref = ?", ref).Do(&ul)
	if err != nil {
		return &ul, err
	}
	err = common.BaseDb.QueryRow(querySql, o.Args()...).Scan(ul.FieldItem()...)
	return &ul, err
}

func FindUserLinks(params map[string]string) (*[]UserLinks, error) {
	// todo
	ull := []UserLinks{}
	return &ull, nil
}
