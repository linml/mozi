package models

import (
	"fmt"
	"github.com/shopspring/decimal"
	"github.com/xiuos/mozi/common"
	"github.com/xiuos/mozi/models/errors"
	"strings"
)

type ReportLottoDayCount struct {
	ID          int             `json:"id"`
	CountDate   string          `json:"count_date"`
	UserID      int             `json:"user_id"`
	Name        string          `json:"name"`
	ParentID    int             `json:"parent_id"`
	Parents     string          `json:"parents"`
	LottoID     int             `json:"lotto_id"`
	GameKind    int             `json:"game_kind"`
	GameType    int             `json:"game_type"`
	TotalCount  int             `json:"total_count"`
	TotalBet    decimal.Decimal `json:"total_bet"`
	TotalPayout decimal.Decimal `json:"total_payout"`
	TotalProfit decimal.Decimal `json:"total_profit"`
	UpdateTime  string          `json:"update_time"`
}

type AgentLottoReport struct {
	groupParent string
	UserID      int             `json:"user_id"`
	Name        string          `json:"name"`
	TotalCount  int             `json:"total_count"`
	TotalBet    decimal.Decimal `json:"total_bet"`
	TotalPayout decimal.Decimal `json:"total_payout"`
	TotalProfit decimal.Decimal `json:"total_profit"`
}

func (r *ReportLottoDayCount) TableName() string {
	return "report_lotto_day_count"
}

func (r *ReportLottoDayCount) Field() []string {
	return []string{"id", "count_date", "user_id", "name", "parent_id", "parents", "lotto_id", "game_kind", "game_type", "total_count", "total_bet", "total_payout", "total_profit", "update_time"}
}

func (r *ReportLottoDayCount) FieldItem() []interface{} {
	return []interface{}{&r.ID, &r.CountDate, &r.UserID, &r.Name, &r.ParentID, &r.Parents, &r.LottoID, &r.GameKind, &r.GameType, &r.TotalCount, &r.TotalBet, &r.TotalPayout, &r.TotalProfit, &r.UpdateTime}
}

func ReplaceLottoDayCount(r *ReportLottoDayCount) error {
	replaceSql := fmt.Sprintf("INSERT INTO %s (count_date, user_id, name, parent_id, parents, game_kind, game_type, lotto_id, total_count, total_bet, total_payout, total_profit, update_time) VALUES (?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?) ON DUPLICATE KEY UPDATE name=VALUES(name),parent_id=VALUES(parent_id),parents=VALUES(parents),game_kind=VALUES(game_kind),game_type=VALUES(game_type),total_count=VALUES(total_count),total_bet=VALUES(total_bet),total_payout=VALUES(total_payout),total_profit=VALUES(total_profit),update_time=VALUES(update_time)", r.TableName())
	_, err := common.BaseDb.Exec(replaceSql, &r.CountDate, &r.UserID, &r.Name, &r.ParentID, &r.Parents, &r.GameKind, &r.GameType, &r.LottoID, &r.TotalCount, &r.TotalBet, &r.TotalPayout, &r.TotalProfit, &r.UpdateTime)
	return err
}

func PageFindAgentLottoReportList(pageParam common.PageParams) (*common.PageResult, *[]AgentLottoReport, error) {
	r := ReportLottoDayCount{}
	conditionSql := fmt.Sprintf(" FROM %s WHERE 1=1 ", r.TableName())
	countSql := "SELECT COUNT(*) AS count " + conditionSql
	querySql := "SELECT SUBSTRING_INDEX(parents, ',', %d) AS group_parent, user_id,name,COALESCE(SUM(total_count), 0) AS total_count,COALESCE(SUM(total_bet), 0) AS total_bet,COALESCE(SUM(total_payout), 0) AS total_payout,COALESCE(SUM(total_profit), 0) AS total_profit " + conditionSql

	if v, ok := pageParam.Params["parent_id"]; ok {
		pid := common.GetInt(v)
		pInfo, err := GetUserRelation(pid)
		if err != nil {
			return &common.PageResult{}, &[]AgentLottoReport{}, errors.New("找不到此用户")
		}
		querySql = fmt.Sprintf(querySql, strings.Count(pInfo.Parents, ",")+2)
	} else if v, ok := pageParam.Params["user_id"]; ok {

		pid := common.GetInt(v)
		pInfo, err := GetUserRelation(pid)
		if err != nil {
			return &common.PageResult{}, &[]AgentLottoReport{}, errors.New("找不到此用户")
		}
		querySql = fmt.Sprintf(querySql, strings.Count(pInfo.Parents, ",")+2)
	} else {
		querySql = fmt.Sprintf(querySql, 2)
	}

	sqlWhere, args := "", []interface{}{}
	if v, ok := pageParam.Params["user_id"]; ok {
		sqlWhere += " AND user_id = ?"
		args = append(args, v)
	}

	if v, ok := pageParam.Params["parent_id"]; ok {
		sqlWhere += " AND (FIND_IN_SET(?, parents) OR user_id = ?) "
		args = append(args, v)
		args = append(args, v)
	}

	if v, ok := pageParam.Params["count_date"]; ok {
		sqlWhere += " AND count_date = ?"
		args = append(args, v)
	}
	if v, ok := pageParam.Params["count_date_from"]; ok {
		sqlWhere += " AND count_date >= ?"
		args = append(args, v)
	}
	if v, ok := pageParam.Params["count_date_to"]; ok {
		sqlWhere += " AND count_date <= ?"
		args = append(args, v)
	}
	var pg common.PageResult
	var data []AgentLottoReport
	var err error

	countSql += sqlWhere
	err = common.BaseDb.QueryRow(countSql, args...).Scan(&pg.TotalCount)
	if err != nil {
		return &pg, &data, err
	}

	if pg.TotalCount < 1 {
		return &pg, &data, err
	}
	sqlWhere += " GROUP BY group_parent ASC ORDER BY LENGTH(group_parent) ASC "

	querySql += sqlWhere
	querySql += fmt.Sprintf(" LIMIT %d,%d", (pageParam.CurrentPage-1)*pageParam.PageRow, pageParam.PageRow)
	rows, err := common.BaseDb.Query(querySql, args...)
	if err != nil {
		return &pg, &data, err
	}
	for rows.Next() {
		d := AgentLottoReport{}
		err = rows.Scan(&d.groupParent, &d.UserID, &d.Name, &d.TotalCount, &d.TotalBet, &d.TotalPayout, &d.TotalProfit)
		if err != nil {
			return &pg, &data, err
		}
		pl := strings.Split(d.groupParent, ",")

		pid := common.GetInt(pl[len(pl)-1])
		if pid != 0 {
			groupInfo, _ := GetUserByID(common.GetInt(pl[len(pl)-1]))
			d.UserID = groupInfo.UserID
			d.Name = groupInfo.Name
		}

		data = append(data, d)
		pg.RecordData = append(pg.RecordData, d)
	}
	pg.PageRow = pageParam.PageRow
	pg.CurrentPage = pageParam.CurrentPage
	pg.PageCount = len(data)
	return &pg, &data, err
}

func PageFindReportLottoDayCountList(pageParam common.PageParams) (*common.PageResult, *[]ReportLottoDayCount, error) {
	r := ReportLottoDayCount{}
	conditionSql := fmt.Sprintf(" FROM %s WHERE 1=1 ", r.TableName())
	countSql := "SELECT COUNT(*) AS count " + conditionSql
	querySql := fmt.Sprintf("SELECT %s ", strings.Join(r.Field(), ",")) + conditionSql
	sqlWhere, args := "", []interface{}{}
	if v, ok := pageParam.Params["id"]; ok {
		sqlWhere += " AND id = ?"
		args = append(args, v)
	}
	if v, ok := pageParam.Params["count_date"]; ok {
		sqlWhere += " AND count_date = ?"
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
	if v, ok := pageParam.Params["parent_id"]; ok {
		sqlWhere += " AND parent_id = ?"
		args = append(args, v)
	}
	if v, ok := pageParam.Params["lotto_id"]; ok {
		sqlWhere += " AND lotto_id = ?"
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

	if v, ok := pageParam.Params["count_date_from"]; ok {
		sqlWhere += " AND count_date >= ?"
		args = append(args, v)
	}
	if v, ok := pageParam.Params["count_date_to"]; ok {
		sqlWhere += " AND count_date <= ?"
		args = append(args, v)
	}
	sqlWhere += " ORDER BY count_date DESC "

	var pg common.PageResult
	var data []ReportLottoDayCount
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
		d := ReportLottoDayCount{}
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
