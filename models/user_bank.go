package models

import (
	"fmt"
	"github.com/xiuos/mozi/common"
)

type UserBank struct {
	ID       int    `json:"id"`
	UserID   int    `json:"user_id"`
	BankCode int    `json:"bank_code"`
	Address  string `json:"address"`
	CardNo   string `json:"card_no"`
	CreateAt string `json:"create_at"`
	UpdateAt string `json:"update_at"`
}

func (u *UserBank) TableName() string {
	return "users"
}

func (u *UserBank) Field() []string {
	return []string{"id", "user_id", "bank_code", "address", "card_no", "create_at", "update_at"}
}

func (u *UserBank) FieldItem() []interface{} {
	return []interface{}{&u.ID, &u.UserID, &u.BankCode, &u.Address, &u.CardNo, &u.CreateAt, &u.UpdateAt}
}

type UserBank4Admin struct {
	ID       int    `json:"id"`
	UserID   int    `json:"user_id"`
	Name     string `json:"name"`
	RealName string `json:"real_name"`
	BankCode string `json:"bank_code"`
	BankName string `json:"bank_name"`
	Address  string `json:"address"`
	CardNo   string `json:"card_no"`
	CreateAt string `json:"create_at"`
	UpdateAt string `json:"update_at"`
}

func PageFindUserBank4AdminList(pageParam PageParams) (*PageResult, *[]UserBank4Admin, error) {
	conditionSql := " FROM user_bank a LEFT JOIN users b ON a.user_id=b.user_id LEFT JOIN user_profile c ON a.user_id=c.user_id LEFT JOIN code_bank d ON a.bank_code=d.bank_code  WHERE 1=1 "
	countSql := "SELECT COUNT(*) AS count " + conditionSql
	querySql := "SELECT a.id,a.user_id,b.name,c.real_name,a.bank_code,d.bank_name,a.address,a.card_no,a.create_at,a.update_at " + conditionSql
	sqlWhere, args := "", []interface{}{}
	if v, ok := pageParam.Params["user_id"]; ok {
		sqlWhere += " AND a.user_id = ?"
		args = append(args, v)
	}
	if v, ok := pageParam.Params["name"]; ok {
		sqlWhere += " AND b.name = ?"
		args = append(args, v)
	}
	if v, ok := pageParam.Params["real_name"]; ok {
		sqlWhere += " AND c.real_name = ?"
		args = append(args, v)
	}
	if v, ok := pageParam.Params["card_no"]; ok {
		sqlWhere += " AND a.card_no = ?"
		args = append(args, v)
	}
	var pg PageResult
	var data []UserBank4Admin
	var err error
	countSql += sqlWhere
	err = common.BaseDb.QueryRow(countSql, args...).Scan(&pg.TotalCount)
	if err != nil {
		return &pg, &data, err
	}

	if pg.TotalCount < 1 {
		return &pg, &data, err
	}

	querySql += sqlWhere
	querySql += fmt.Sprintf(" LIMIT %d,%d", (pageParam.CurrentPage-1)*pageParam.PageRow, pageParam.PageRow)
	rows, err := common.BaseDb.Query(querySql, args...)
	if err != nil {
		return &pg, &data, err
	}
	for rows.Next() {
		d := UserBank4Admin{}
		err = rows.Scan(&d.ID, &d.UserID, &d.Name, &d.RealName, &d.BankCode, &d.BankName, &d.Address, &d.CardNo, &d.CreateAt, &d.UpdateAt)
		if err != nil {
			return &pg, &data, err
		}
		data = append(data, d)
		pg.RecordData = append(pg.RecordData, d)
	}
	pg.PageRow = pageParam.PageRow
	pg.CurrentPage = pageParam.CurrentPage
	pg.PageCount = len(data)
	return &pg, &data, err
}
