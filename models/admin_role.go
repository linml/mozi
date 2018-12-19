package models

import (
	"fmt"
	"github.com/xiuos/mozi/common"
	"github.com/xiuos/mozi/xorm"
)

type AdminRole struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
}

func (am *AdminRole) TableName() string {
	return "admin_role"
}

func (am *AdminRole) Field() []string {
	return []string{"id", "name"}
}

func (ar *AdminRole) FieldItem() []interface{} {
	return []interface{}{&ar.ID, &ar.Name}
}

func CreateAdminRole(ar *AdminRole) error {
	createSql := fmt.Sprintf("INSERT INTO %s SET name = ?", ar.TableName())
	_, err := common.BaseDb.Exec(createSql, ar.Name)
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

func FindAdminRole(param map[string]string) (*[]AdminRole, error) {
	arl := []AdminRole{}
	t := AdminRole{}
	querySql := fmt.Sprintf("SELECT id, name FROM %s WHERE 1=1", t.TableName())
	sqlWhere, args := "", []interface{}{}
	if v, ok := param["id"]; ok {
		sqlWhere += " AND id = ?"
		args = append(args, v)
	}
	if v, ok := param["name"]; ok {
		sqlWhere += " AND name = ?"
		args = append(args, v)
	}
	querySql += sqlWhere
	rows, err := common.BaseDb.Query(querySql, args...)
	if err != nil {
		fmt.Println(err)
		return &arl, err
	}
	for rows.Next() {
		am := AdminRole{}
		err = rows.Scan(&am.ID, &am.Name)
		if err != nil {
			return &arl, err
		}
		arl = append(arl, am)
	}
	return &arl, nil
}
