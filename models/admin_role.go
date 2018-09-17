package models

import (
	"fmt"
	"github.com/xiuos/mozi/common"
	"github.com/xiuos/xorm"
)

type AdminRole struct {
	ID     int    `json:"id"`
	Name   string `json:"name"`
	Menu   string `json:"menu"`
	Status int    `json:"status"`
}

func (am *AdminRole) TableName() string {
	return "admin_role"
}

func (am *AdminRole) Field() []string {
	return []string{"id", "name", "menu", "status"}
}

func (ar *AdminRole) FieldItem() []interface{} {
	return []interface{}{&ar.ID, &ar.Name, &ar.Menu, &ar.Status}
}

func CreateAdminRole(ar *AdminRole) error {
	createSql := fmt.Sprintf("INSERT INTO %s SET name = ?, menu = ?, status = ?", ar.TableName())
	_, err := common.BaseDb.Exec(createSql, ar.Name, ar.Menu, ar.Status)
	return err
}

func GetAdminRole(role int) (*AdminRole, error) {
	o := xorm.Orm{}
	ar := AdminRole{}
	querySql, err := o.Table(ar.TableName()).Query().Where("id = ?", role).Do(&ar)
	if err != nil {
		return &ar, err
	}
	err = common.BaseDb.QueryRow(querySql, o.Args()...).Scan(ar.FieldItem()...)
	return &ar, err
}
