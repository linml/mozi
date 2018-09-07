package models

import (
	"database/sql"
	"github.com/xiuos/mozi/common"
	"github.com/xiuos/xorm"
)

const (
	TableUser = "user"
)

type User struct {
	UserID   int
	Username string
	Password string
	Status   int
}

func (u *User) Field() []string {
	return []string{"user_id", "username", "password", "status"}
}

func (u *User) FieldItem() []interface{} {
	return []interface{}{&u.UserID, &u.Username, &u.Password, &u.Status}
}

func CreateUserTx(tx *sql.Tx, u *User) (sql.Result, error) {
	createSql := "INSERT INTO " + TableUser + " SET username = ?, password = ?, status = ?"
	rs, err := tx.Exec(createSql, u.Username, &u.Password, &u.Status)
	if err != nil {
		return nil, err
	}
	return rs, nil
}

func CreateUser(u *User) error {
	createSql := "INSERT INTO " + TableUser + " SET username = ?, password = ?, status = ?"
	_, err := common.BaseDb.Exec(createSql, u.Username, &u.Password, &u.Status)
	return err
}

func GetUserByID(id int) (*User, error) {
	o := xorm.Orm{}
	u := User{}
	querySql, err := o.Table(TableUser).Query().Where("user_id = ?", id).Do(&u)
	if err != nil {
		return &u, err
	}
	err = common.BaseDb.QueryRow(querySql, o.Args()...).Scan(u.FieldItem()...)
	return &u, err
}

func GetUserByName(name string) (*User, error) {
	o := xorm.Orm{}
	u := User{}
	querySql, err := o.Table(TableUser).Query().Where("username = ?", name).Do(&u)
	if err != nil {
		return &u, err
	}
	err = common.BaseDb.QueryRow(querySql, o.Args()...).Scan(u.FieldItem()...)
	return &u, err
}
