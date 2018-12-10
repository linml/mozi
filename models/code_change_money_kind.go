package models

import (
	"fmt"
	"github.com/xiuos/mozi/common"
)

const (
	ChangeKindLotto    = 1 // 彩票
	ChangeKindDeposit  = 2 // 入款
	ChangeKindWithdraw = 3 // 出款
	ChangeKindTransfer = 4 // 转账
	ChangeKindActivity = 5 // 活动
	ChangeKindWage     = 6 // 工资
	ChangeKindDividend = 7 // 分红
)

type CodeChangeMoneyKind struct {
	ID        int    `json:"id"`
	Name      string `json:"name"`
	SortIndex int    `json:"sort_index"`
	Status    int    `json:"status"`
}

func (b *CodeChangeMoneyKind) TableName() string {
	return "code_money_change_kind"
}

func (b *CodeChangeMoneyKind) Field() []string {
	return []string{"id", "name", "sort_index", "status"}
}

func (b *CodeChangeMoneyKind) FieldItem() []interface{} {
	return []interface{}{&b.ID, &b.Name, &b.SortIndex, &b.Status}
}

func FindCodeChangeMoneyKindList(param map[string]string) (*[]CodeChangeMoneyKind, error) {
	arl := []CodeChangeMoneyKind{}
	t := CodeChangeMoneyKind{}
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
		am := CodeChangeMoneyKind{}
		err = rows.Scan(&am.ID, &am.Name, &am.SortIndex, &am.Status)
		if err != nil {
			return &arl, err
		}
		arl = append(arl, am)
	}
	return &arl, nil
}
