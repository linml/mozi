package models

import (
	"database/sql"
	"fmt"
	"github.com/xiuos/mozi/common"
)

type Bank struct {
	ID        int    `json:"id"`
	BankCode  string `json:"bank_code"`
	BankName  string `json:"bank_name"`
	SortIndex int    `json:"sort_index"`
	GroupType int    `json:"group_type"`
	Status    string `json:"status"`
}

func (b *Bank) TableName() string {
	return "code_bank"
}

func (b *Bank) Field() []string {
	return []string{"id", "bank_code", "bank_name", "sort_index", "group", "status"}
}

func (b *Bank) FieldItem() []interface{} {
	return []interface{}{&b.ID, &b.BankCode, &b.BankName, &b.SortIndex, &b.GroupType, &b.Status}
}

func CreatebankTx(tx *sql.Tx, b *Bank) (sql.Result, error) {
	createSql := "INSERT INTO " + b.TableName() + " SET bank_code = ?, bank_name = ?, sort_index = ?, group_type = ?, status = ?"
	rs, err := tx.Exec(createSql, &b.BankCode, &b.BankName, &b.SortIndex, &b.GroupType, &b.Status)
	if err != nil {
		return nil, err
	}
	return rs, nil
}

func PageFindCodeBankList(pageParam PageParams) (*PageResult, *[]Bank, error) {
	conditionSql := " FROM code_bank WHERE 1=1 "
	countSql := "SELECT COUNT(*) AS count " + conditionSql
	querySql := "SELECT id, bank_code, bank_name, sort_index, group_type, status " + conditionSql
	sqlWhere, args := "", []interface{}{}
	if v, ok := pageParam.Params["id"]; ok {
		sqlWhere += " AND id = ?"
		args = append(args, v)
	}
	if v, ok := pageParam.Params["bank_code"]; ok {
		sqlWhere += " AND bank_code = ?"
		args = append(args, v)
	}
	if v, ok := pageParam.Params["bank_name"]; ok {
		sqlWhere += " AND bank_name = ?"
		args = append(args, v)
	}
	if v, ok := pageParam.Params["group"]; ok {
		sqlWhere += " AND group = ?"
		args = append(args, v)
	}
	if v, ok := pageParam.Params["status"]; ok {
		sqlWhere += " AND status = ?"
		args = append(args, v)
	}

	var pg PageResult
	var data []Bank
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
		d := Bank{}
		err = rows.Scan(&d.ID, &d.BankCode, &d.BankName, &d.SortIndex, &d.GroupType, &d.Status)
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
