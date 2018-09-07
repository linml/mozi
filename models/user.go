package models

import (
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
