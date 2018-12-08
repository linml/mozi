package models

import (
	"fmt"
	"github.com/xiuos/mozi/common"
)

const (
	ActionAdminAddRole               = 1  // 添加角色
	ActionAdminDeleteRole            = 2  // 删除角色
	ActionAdminUpdateRole            = 3  // 设置角色
	ActionAdminSetRoleMenu           = 4  // 设置角色菜单
	ActionAdminSetAdminRole          = 5  // 修改管理员角色
	ActionAdminSetAdminStatus        = 6  // 修改管理员状态
	ActionAdminAddAdmin              = 7  // 添加管理员
	ActionAdminAddMember             = 8  // 添加用户
	ActionAdminSetCodeLottoStatus    = 9  // 设置彩票当前状态
	ActionAdminSetCodeLottoSortIndex = 10 // 设置彩票排序
	ActionAdminSetCodeLottoIsShow    = 11 // 设置彩票显示
	ActionAdminSetLottoOddsInfo      = 12 // 设置彩票赔率信息
)

type CodeActionAdmin struct {
	ModuleID  int    `json:"module_id"`
	ID        int    `json:"id"`
	Name      string `json:"name"`
	SortIndex int    `json:"sort_index"`
	Status    int    `json:"status"`
}

func (b *CodeActionAdmin) TableName() string {
	return "code_action_admin"
}

func (b *CodeActionAdmin) Field() []string {
	return []string{"module_id", "id", "name", "sort_index", "status"}
}

func (b *CodeActionAdmin) FieldItem() []interface{} {
	return []interface{}{&b.ModuleID, &b.ID, &b.Name, &b.SortIndex, &b.Status}
}

func FindCodeActionAdmin(param map[string]string) (*[]CodeActionAdmin, error) {
	arl := []CodeActionAdmin{}
	t := CodeActionAdmin{}
	querySql := fmt.Sprintf("SELECT module_id,id,name,sort_index,status FROM %s WHERE 1=1", t.TableName())
	sqlWhere, args := "", []interface{}{}
	if v, ok := param["module_id"]; ok {
		sqlWhere += " AND module_id = ?"
		args = append(args, v)
	}
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
		am := CodeActionAdmin{}
		err = rows.Scan(&am.ModuleID, &am.ID, &am.Name, &am.SortIndex, &am.Status)
		if err != nil {
			return &arl, err
		}
		arl = append(arl, am)
	}
	return &arl, nil
}
