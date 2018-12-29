package models

import (
	"fmt"
	"github.com/xiuos/mozi/common"
	"strings"
)

const (
	ChangeTypeLottoBet    = 101 // 彩票-下注
	ChangeTypeLottoPayout = 102 // 彩票-派彩
	ChangeTypeLottoCancel = 103 // 彩票-撤单

	ChangeKindDepositManual   = 200 //入款-人工入款
	ChangeKindDepositCencel   = 201 //入款-冲账-取消出款
	ChangeKindDepositRepeate  = 202 //入款-冲账-重复出款
	ChangeKindDepositPayoff   = 203 //入款-入款优惠(活动)
	ChangeKindDepositActivity = 204 //入款-活动优惠(活动)
	ChangeKindDepositZero     = 205 //负数额度归零
	ChangeKindDepositBonus    = 206 //红利
	ChangeKindDepositOther    = 220 //入款-其他(活动)

	ChangeKindWithdrawManual   = 300 //出款-人工出款
	ChangeKindWithdrawPayoff   = 303 //出款-人工扣除-入款优惠(活动)
	ChangeKindWithdrawActivity = 304 //出款-人工扣除-活动优惠(活动)
	ChangeKindWithdrawBonus    = 306 //人工扣除-红利
	ChangeKindWithdrawOther    = 320 //出款-其他
)

type CodeChangeMoneyType struct {
	ChangeKind int    `json:"change_kind"`
	ID         int    `json:"id"`
	Name       string `json:"name"`
	SortIndex  int    `json:"sort_index"`
	Status     int    `json:"status"`
}

func (b *CodeChangeMoneyType) TableName() string {
	return "code_money_change_type"
}

func (b *CodeChangeMoneyType) Field() []string {
	return []string{"change_kind", "id", "name", "sort_index", "status"}
}

func (b *CodeChangeMoneyType) FieldItem() []interface{} {
	return []interface{}{&b.ChangeKind, &b.ID, &b.Name, &b.SortIndex, &b.Status}
}

func GetCodeChangeMoneyType(id int) (*CodeChangeMoneyType, error) {
	t := CodeChangeMoneyType{}
	querySql := fmt.Sprintf("SELECT %s FROM %s WHERE id=?", strings.Join(t.Field(), ","), t.TableName())
	err := common.BaseDb.QueryRow(querySql, id).Scan(t.FieldItem()...)
	return &t, err
}
func GetCodeChangeMoneyType1(changeKind int, id int) (*CodeChangeMoneyType, error) {
	t := CodeChangeMoneyType{}
	querySql := fmt.Sprintf("SELECT %s FROM %s WHERE change_kind=? AND id=? ", strings.Join(t.Field(), ","), t.TableName())
	err := common.BaseDb.QueryRow(querySql, changeKind, id).Scan(t.FieldItem()...)
	return &t, err
}

func FindCodeChangeMoneyTypeList(param map[string]string) (*[]CodeChangeMoneyType, error) {
	arl := []CodeChangeMoneyType{}
	t := CodeChangeMoneyType{}
	querySql := fmt.Sprintf("SELECT change_kind,id,name,sort_index,status FROM %s WHERE 1=1", t.TableName())
	sqlWhere, args := "", []interface{}{}
	if v, ok := param["change_kind"]; ok {
		sqlWhere += " AND change_kind = ?"
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
		am := CodeChangeMoneyType{}
		err = rows.Scan(&am.ChangeKind, &am.ID, &am.Name, &am.SortIndex, &am.Status)
		if err != nil {
			return &arl, err
		}
		arl = append(arl, am)
	}
	return &arl, nil
}
