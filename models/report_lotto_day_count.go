package models

import (
	"fmt"
	"github.com/shopspring/decimal"
	"github.com/xiuos/mozi/common"
)

type ReportLottoDayCount struct {
	ID          int             `json:"id"`
	CountDate   string          `json:"count_date"`
	UserID      int             `json:"user_id"`
	Name        string          `json:"name"`
	ParentID    int             `json:"parent_id"`
	Parents     string          `json:"parents"`
	GameKind    int             `json:"game_kind"`
	GameType    int             `json:"game_type"`
	TotalCount  int             `json:"total_count"`
	TotalBet    decimal.Decimal `json:"total_bet"`
	TotalPayout decimal.Decimal `json:"total_payout"`
	TotalProfit decimal.Decimal `json:"total_profit"`
	UpdateTime  string          `json:"update_time"`
}

func (r *ReportLottoDayCount) TableName() string {
	return "report_lotto_day_count"
}

func (r *ReportLottoDayCount) Field() []string {
	return []string{"id", "count_date", "user_id", "name", "parent_id", "parents", "game_kind", "game_type", "total_count", "total_bet", "total_payout", "total_profit", "update_time"}
}

func (r *ReportLottoDayCount) FieldItem() []interface{} {
	return []interface{}{&r.ID, &r.CountDate, &r.UserID, &r.Name, &r.ParentID, &r.Parents, &r.GameKind, &r.GameType, &r.TotalCount, &r.TotalBet, &r.TotalPayout, &r.TotalProfit, &r.UpdateTime}
}

func ReplaceLottoDayCount(r *ReportLottoDayCount) error {
	replaceSql := fmt.Sprintf("INSERT INTO %s (count_date, user_id, name, parent_id, parents, game_kind, game_type, total_count, total_bet, total_payout, total_profit, update_time) VALUES (?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?) ON DUPLICATE KEY UPDATE name=VALUES(name),parent_id=VALUES(parent_id),parents=VALUES(parents),game_kind=VALUES(game_kind),game_type=VALUES(game_type),total_count=VALUES(total_count),total_bet=VALUES(total_bet),total_payout=VALUES(total_payout),total_profit=VALUES(total_profit),update_time=VALUES(update_time)", r.TableName())
	_, err := common.BaseDb.Exec(replaceSql, &r.CountDate, &r.UserID, &r.Name, &r.ParentID, &r.Parents, &r.GameKind, &r.GameType, &r.TotalCount, &r.TotalBet, &r.TotalPayout, &r.TotalProfit, &r.UpdateTime)
	return err
}

func PageFindReportLottoDayCountList(pageParam common.PageParams) (*common.PageResult, *[]RecordAdminAction, error) {
	r := RecordAdminAction{}
	conditionSql := fmt.Sprintf(" FROM %s WHERE 1=1 ", r.TableName())
	countSql := "SELECT COUNT(*) AS count " + conditionSql
	querySql := "SELECT id,user_id,name,action_module,action_id,content,ip,record_at,success,message,operator_id,operator_name,operator_type " + conditionSql
	sqlWhere, args := "", []interface{}{}
	if v, ok := pageParam.Params["user_id"]; ok {
		sqlWhere += " AND user_id = ?"
		args = append(args, v)
	}
	if v, ok := pageParam.Params["name"]; ok {
		sqlWhere += " AND name = ?"
		args = append(args, v)
	}
	if v, ok := pageParam.Params["action_module"]; ok {
		sqlWhere += " AND action_module = ?"
		args = append(args, v)
	}
	if v, ok := pageParam.Params["action_id"]; ok {
		sqlWhere += " AND action_id = ?"
		args = append(args, v)
	}
	if v, ok := pageParam.Params["ip"]; ok {
		sqlWhere += " AND ip = ?"
		args = append(args, v)
	}
	if v, ok := pageParam.Params["success"]; ok {
		sqlWhere += " AND success = ?"
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
	if v, ok := pageParam.Params["operator_type"]; ok {
		sqlWhere += " AND operator_type = ?"
		args = append(args, v)
	}
	if v, ok := pageParam.Params["start_at"]; ok {
		sqlWhere += " AND record_at >= ?"
		args = append(args, v)
	}
	if v, ok := pageParam.Params["end_at"]; ok {
		sqlWhere += " AND record_at <= ?"
		args = append(args, v)
	}
	sqlWhere += " ORDER BY record_at DESC "

	var pg common.PageResult
	var data []RecordAdminAction
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
		d := RecordAdminAction{}
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
