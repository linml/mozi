package models

import (
	"database/sql"
	"fmt"
	"github.com/shopspring/decimal"
	"github.com/xiuos/mozi/common"
	"strings"
)

type RecordMoneyChange struct {
	ID            int             `json:"id"`
	LinkID        string          `json:"link_id"`
	UserID        int             `json:"user_id"`
	Name          string          `json:"name"`
	GameKind      int             `json:"game_kind"`
	GameType      int             `json:"game_type"`
	ChangeKind    int             `json:"change_kind"`
	ChangeType    int             `json:"change_type"`
	BeforeBalance decimal.Decimal `json:"before_balance"`
	Amount        decimal.Decimal `json:"amount"`
	AfterBalance  decimal.Decimal `json:"after_balance"`
	RecordDate    string          `json:"record_date"`
	RecordAt      string          `json:"record_at"`
	OperateType   int             `json:"operate_type"`
	OperatorID    int             `json:"operator_id"`
	OperatorName  string          `json:"operator_name"`
	Remark        string          `json:"remark"`
}

func (r *RecordMoneyChange) TableName() string {
	return "record_money_change"
}

func (r *RecordMoneyChange) Field() []string {
	return []string{"id", "link_id", "user_id", "name", "game_kind", "game_type", "change_kind", "change_type", "before_balance", "amount", "after_balance", "record_date", "record_at", "operate_type", "operator_id", "operator_name", "remark"}
}

func (r *RecordMoneyChange) FieldItem() []interface{} {
	return []interface{}{&r.ID, &r.LinkID, &r.UserID, &r.Name, &r.GameKind, &r.GameType, &r.ChangeKind, &r.ChangeType, &r.BeforeBalance, &r.Amount, &r.AfterBalance, &r.RecordDate, &r.RecordAt, &r.OperateType, &r.OperatorID, &r.OperatorName, &r.Remark}
}

func CreateRecordMoneyChangeTx(tx *sql.Tx, r *RecordMoneyChange) (sql.Result, error) {
	createSql := "INSERT INTO " + r.TableName() + " SET link_id = ?, user_id = ?, name = ?, game_kind = ?, game_type = ?, change_kind = ?, change_type = ?, before_balance = ?, amount = ?, after_balance = ?, record_date = ?, record_at = ?, operate_type = ?, operator_id = ?, operator_name = ?, remark = ?"
	rs, err := tx.Exec(createSql, r.LinkID, r.UserID, r.Name, r.GameKind, r.GameType, r.ChangeKind, r.ChangeType, r.BeforeBalance, r.Amount, r.AfterBalance, r.RecordDate, r.RecordAt, r.OperateType, r.OperatorID, r.OperatorName, r.Remark)
	if err != nil {
		return nil, err
	}
	return rs, nil
}

func PageFindRecordMoneyList(pageParam common.PageParams) (*common.PageResult, *[]RecordMoneyChange, error) {
	t := RecordMoneyChange{}
	conditionSql := fmt.Sprintf(" FROM %s WHERE 1=1 ", t.TableName())
	countSql := "SELECT COUNT(*) AS count " + conditionSql
	querySql := fmt.Sprintf("SELECT %s ", strings.Join(t.Field(), ",")) + conditionSql
	sqlWhere, args := "", []interface{}{}
	if v, ok := pageParam.Params["id"]; ok {
		sqlWhere += " AND id = ?"
		args = append(args, v)
	}
	if v, ok := pageParam.Params["link_id"]; ok {
		sqlWhere += " AND link_id = ?"
		args = append(args, v)
	}
	if v, ok := pageParam.Params["user_id"]; ok {
		sqlWhere += " AND user_id = ?"
		args = append(args, v)
	}
	if v, ok := pageParam.Params["name"]; ok {
		sqlWhere += " AND name = ?"
		args = append(args, v)
	}
	if v, ok := pageParam.Params["game_kind"]; ok {
		sqlWhere += " AND game_kind = ?"
		args = append(args, v)
	}
	if v, ok := pageParam.Params["game_type"]; ok {
		sqlWhere += " AND game_type = ?"
		args = append(args, v)
	}
	if v, ok := pageParam.Params["change_kind"]; ok {
		sqlWhere += " AND change_kind = ?"
		args = append(args, v)
	}
	if v, ok := pageParam.Params["change_type"]; ok {
		sqlWhere += " AND change_type = ?"
		args = append(args, v)
	}
	if v, ok := pageParam.Params["before_balance"]; ok {
		sqlWhere += " AND before_balance = ?"
		args = append(args, v)
	}
	if v, ok := pageParam.Params["amount"]; ok {
		sqlWhere += " AND amount = ?"
		args = append(args, v)
	}
	if v, ok := pageParam.Params["after_balance"]; ok {
		sqlWhere += " AND after_balance = ?"
		args = append(args, v)
	}
	if v, ok := pageParam.Params["record_date"]; ok {
		sqlWhere += " AND record_date = ?"
		args = append(args, v)
	}
	if v, ok := pageParam.Params["record_at"]; ok {
		sqlWhere += " AND record_at = ?"
		args = append(args, v)
	}
	if v, ok := pageParam.Params["operate_type"]; ok {
		sqlWhere += " AND operate_type = ?"
		args = append(args, v)
	}
	if v, ok := pageParam.Params["operator_id"]; ok {
		sqlWhere += " AND operator_id = ?"
		args = append(args, v)
	}
	if v, ok := pageParam.Params["operator_name"]; ok {
		sqlWhere += " AND operator_name = ?"
		args = append(args, v)
	}
	if v, ok := pageParam.Params["remark"]; ok {
		sqlWhere += " AND remark = ?"
		args = append(args, v)
	}

	if v, ok := pageParam.Params["amount_min"]; ok {
		sqlWhere += " AND amount >= ?"
		args = append(args, v)
	}
	if v, ok := pageParam.Params["amount_max"]; ok {
		sqlWhere += " AND amount <= ?"
		args = append(args, v)
	}

	if v, ok := pageParam.Params["record_date_from"]; ok {
		sqlWhere += " AND record_date >= ?"
		args = append(args, v)
	}
	if v, ok := pageParam.Params["record_date_to"]; ok {
		sqlWhere += " AND record_date <= ?"
		args = append(args, v)
	}
	if v, ok := pageParam.Params["record_at_from"]; ok {
		sqlWhere += " AND record_at >= ?"
		args = append(args, v)
	}
	if v, ok := pageParam.Params["record_at_to"]; ok {
		sqlWhere += " AND record_at <= ?"
		args = append(args, v)
	}

	sqlWhere += " ORDER BY id DESC"

	var pg common.PageResult
	var data []RecordMoneyChange
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
		d := RecordMoneyChange{}
		err = rows.Scan(d.FieldItem()...)
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
