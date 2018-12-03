package models

import (
	"fmt"
	"github.com/xiuos/mozi/common"
)

type AdminRoleMenu struct {
	RoleID int `json:"role_id"`
	MenuID int `json:"menu_id"`
}

func (am *AdminRoleMenu) TableName() string {
	return "admin_role_menu"
}

func (am *AdminRoleMenu) Field() []string {
	return []string{"role_id", "menu_id"}
}

func (ar *AdminRoleMenu) FieldItem() []interface{} {
	return []interface{}{&ar.RoleID, &ar.MenuID}
}

func AddRoleMenu(ar *AdminRoleMenu) error {
	createSql := fmt.Sprintf("INSERT INTO %s SET role_id = ?, menu_id = ?", ar.TableName())
	_, err := common.BaseDb.Exec(createSql, ar.RoleID, ar.MenuID)
	return err
}

func DelRoleMenu(ar *AdminRoleMenu) error {
	deleteSql := fmt.Sprintf("DELETE FROM %s WHERE role_id = ? AND menu_id = ?", ar.TableName())
	_, err := common.BaseDb.Exec(deleteSql, ar.RoleID, ar.MenuID)
	return err
}

func FindRoleMenu(param map[string]string) (*[]AdminRoleMenu, error) {
	aml := []AdminRoleMenu{}
	querySql := "SELECT role_id, menu_id FROM admin_role_menu  WHERE 1=1 "
	sqlWhere, args := "", []interface{}{}
	if v, ok := param["role_id"]; ok {
		sqlWhere += " AND role_id = ?"
		args = append(args, v)
	}
	if v, ok := param["menu_id"]; ok {
		sqlWhere += " AND menu_id = ?"
		args = append(args, v)
	}
	querySql += sqlWhere
	rows, err := common.BaseDb.Query(querySql, args...)
	if err != nil {
		return &aml, err
	}
	for rows.Next() {
		am := AdminRoleMenu{}
		err = rows.Scan(&am.RoleID, &am.MenuID)
		if err != nil {
			return &aml, err
		}
		aml = append(aml, am)
	}
	return &aml, nil
}
