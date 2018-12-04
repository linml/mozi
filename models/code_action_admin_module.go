package models

import (
	"fmt"
	"github.com/xiuos/mozi/common"
)

const (
	ActionAdminModuleUser     = 1 // 用户模块
	ActionAdminModuleMoney    = 2 // 资金模块
	ActionAdminModuleActivity = 3 // 活动模块
	ActionAdminModuleSystem   = 4 // 系统模块
)

type CodeActionAdminModule struct {
	ID        int    `json:"id"`
	Name      string `json:"name"`
	SortIndex int    `json:"sort_index"`
	Status    int    `json:"status"`
}

func (b *CodeActionAdminModule) TableName() string {
	return "code_action_admin_module"
}

func (b *CodeActionAdminModule) Field() []string {
	return []string{"id", "name", "sort_index", "status"}
}

func (b *CodeActionAdminModule) FieldItem() []interface{} {
	return []interface{}{&b.ID, &b.Name, &b.SortIndex, &b.Status}
}

func FindCodeActionAdminModule(param map[string]string) (*[]CodeActionAdminModule, error) {
	arl := []CodeActionAdminModule{}
	t := CodeActionAdminModule{}
	querySql := fmt.Sprintf("SELECT id,name,sort_index,status FROM %s WHERE 1=1", t.TableName())
	sqlWhere, args := "", []interface{}{}
	if v, ok := param["id"]; ok {
		sqlWhere += " AND id = ?"
		args = append(args, v)
	}
	if v, ok := param["name"]; ok {
		sqlWhere += " AND name = ?"
		args = append(args, v)
	}
	if v, ok := param["status"]; ok {
		sqlWhere += " AND status = ?"
		args = append(args, v)
	}
	querySql += sqlWhere
	rows, err := common.BaseDb.Query(querySql, args...)
	if err != nil {
		fmt.Println(err)
		return &arl, err
	}
	for rows.Next() {
		am := CodeActionAdminModule{}
		err = rows.Scan(&am.ID, &am.Name, &am.SortIndex, &am.Status)
		if err != nil {
			return &arl, err
		}
		arl = append(arl, am)
	}
	return &arl, nil
}
