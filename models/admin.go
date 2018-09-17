package models

import (
	"database/sql"
	"github.com/xiuos/mozi/common"
	"github.com/xiuos/xorm"
)

type AdminUser struct {
	UserID             int    `json:"user_id"`
	Name               string `json:"username"`
	Password           string `json:"password"`
	GoogleSecret       string `json:"google_secret"`
	GoogleSecretStatus int    `json:"google_secret_status"`
	Role               int    `json:"role"`
	Status             int    `json:"status"`
	CreatedAt          string `json:"created_at"`
}

func (u *AdminUser) TableName() string {
	return "admin"
}

func (u *AdminUser) Field() []string {
	return []string{"user_id", "name", "password", "google_secret", "google_secret_status", "role", "status", "created_at"}
}

func (u *AdminUser) FieldItem() []interface{} {
	return []interface{}{&u.UserID, &u.Name, &u.Password, &u.GoogleSecret, &u.GoogleSecretStatus, &u.Role, &u.Status, &u.CreatedAt}
}

func CreateAdminUserTx(tx *sql.Tx, u *AdminUser) (sql.Result, error) {
	createSql := "INSERT INTO " + u.TableName() + " SET name = ?, password = ?, google_secret = ?, google_secret_status = ?, role = ?, status = ?, created_at = ?"
	rs, err := tx.Exec(createSql, &u.Name, &u.Password, &u.GoogleSecret, &u.GoogleSecretStatus, &u.Role, &u.Status, &u.CreatedAt)
	if err != nil {
		return nil, err
	}
	return rs, nil
}

func GetAdminUserByID(uid int) (*AdminUser, error) {
	o := xorm.Orm{}
	u := AdminUser{}
	querySql, err := o.Table(u.TableName()).Query().Where("user_id = ?", uid).Do(&u)
	if err != nil {
		return &u, err
	}
	err = common.BaseDb.QueryRow(querySql, o.Args()...).Scan(u.FieldItem()...)
	return &u, err
}

func GetAdminUserByName(name string) (*AdminUser, error) {
	o := xorm.Orm{}
	u := AdminUser{}
	querySql, err := o.Table(u.TableName()).Query().Where("name = ?", name).Do(&u)
	if err != nil {
		return &u, err
	}
	err = common.BaseDb.QueryRow(querySql, o.Args()...).Scan(u.FieldItem()...)
	return &u, err
}
