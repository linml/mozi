package models

import (
	"database/sql"
	"fmt"
	"github.com/xiuos/mozi/common"
	"github.com/xiuos/xorm"
)

const (
	TableUser = "users"
)

const (
	UserStatusDisable      = 0 // 用户状态禁用
	UserStatusEnable       = 1 // 用户状态启用
	USER_STATUS_UNACTIVATE = 2 // 用户状态未激活
)

type User struct {
	UserID   int
	Name     string
	Password string
	Status   int
}

func (u *User) Field() []string {
	return []string{"user_id", "name", "password", "status"}
}

func (u *User) FieldItem() []interface{} {
	return []interface{}{&u.UserID, &u.Name, &u.Password, &u.Status}
}

func CreateUserTx(tx *sql.Tx, u *User) (sql.Result, error) {
	createSql := "INSERT INTO " + TableUser + " SET name = ?, password = ?, status = ?"
	rs, err := tx.Exec(createSql, u.Name, &u.Password, &u.Status)
	if err != nil {
		return nil, err
	}
	return rs, nil
}

func CreateUser(u *User) error {
	createSql := "INSERT INTO " + TableUser + " SET name = ?, password = ?, status = ?"
	_, err := common.BaseDb.Exec(createSql, u.Name, &u.Password, &u.Status)
	return err
}

func GetUserByID(uid int) (*User, error) {
	o := xorm.Orm{}
	u := User{}
	querySql, err := o.Table(TableUser).Query().Where("user_id = ?", uid).Do(&u)
	if err != nil {
		return &u, err
	}
	err = common.BaseDb.QueryRow(querySql, o.Args()...).Scan(u.FieldItem()...)
	return &u, err
}

func GetUserByName(name string) (*User, error) {
	o := xorm.Orm{}
	u := User{}
	querySql, err := o.Table(TableUser).Query().Where("name = ?", name).Do(&u)
	if err != nil {
		return &u, err
	}
	err = common.BaseDb.QueryRow(querySql, o.Args()...).Scan(u.FieldItem()...)
	return &u, err
}

func IsUserExist(name string) (bool, error) {
	querySql := fmt.Sprintf("SELECT COUNT(*) as count FROM %s Where name = ?", TableUser)
	var c int
	err := common.BaseDb.QueryRow(querySql, name).Scan(&c)
	if err != nil {
		return false, err
	}
	if c > 0 {
		return true, nil
	} else {
		return false, nil
	}
}

func SetPassword(uid int, password string) error {
	hashedPassword, err := common.HashPassword(password)

	if err != nil {
		return err
	}
	updateSql := fmt.Sprintf("UPDATE %s SET password = ? WHERE user_id = ?", TableUser)
	_, err = common.BaseDb.Exec(updateSql, hashedPassword, uid)
	return err
}
