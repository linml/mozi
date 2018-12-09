package lotto

import (
	"fmt"
	"github.com/xiuos/mozi/common"
	"github.com/xiuos/xorm"
	"strings"
)

type LottoResult struct {
	ID         int    `json:"id"`
	LottoID    int    `json:"lotto_id"`
	Issue      string `json:"issue"`
	DrawNumber string `json:"draw_number"`
	Status     int    `json:"status"`
	StartTime  string `json:"start_time"`
	StopTime   string `json:"stop_time"`
	ResultTime string `json:"result_time"`
	IssueDate  string `json:"issue_date"`
	UpdateTime string `json:"update_time"`
}

func (l *LottoResult) TableName(id int) string {
	return fmt.Sprintf("lotto_result_%d", id)
}

func (l *LottoResult) Field() []string {
	return []string{"id", "lotto_id", "issue", "draw_number", "status", "start_time", "stop_time", "result_time", "issue_date", "update_time"}
}

func (l *LottoResult) FieldItem() []interface{} {
	return []interface{}{&l.ID, &l.LottoID, &l.Issue, &l.DrawNumber, &l.Status, &l.StartTime, &l.StopTime, &l.ResultTime, &l.IssueDate, &l.UpdateTime}
}

func GetLottoResult(lid int, issue string) (*LottoResult, error) {
	o := xorm.Orm{}
	l := LottoResult{}
	querySql, err := o.Table(l.TableName(lid)).Query().Where("lotto_id = ? AND issue = ?", lid, issue).Do(&l)
	if err != nil {
		return &l, err
	}
	err = common.BaseDb.QueryRow(querySql, o.Args()...).Scan(l.FieldItem()...)
	return &l, err
}

func CreateIssueInfo(lid int, data *[]LottoResult) error {
	var vals []interface{}
	l := LottoResult{}
	sql := fmt.Sprintf("INSERT IGNORE INTO %s (lotto_id, issue, draw_number, status, start_time, stop_time, result_time, issue_date, update_time) VALUES ", l.TableName(lid))
	for i, _ := range *data {
		sql += "(?, ?, ?, ?, ?, ?, ?, ?, ?),"
		vals = append(vals, (*data)[i].LottoID, (*data)[i].Issue, "", 0, (*data)[i].StartTime, (*data)[i].StopTime, (*data)[i].ResultTime, (*data)[i].IssueDate, "")
	}
	sql = sql[0:len(sql)-1] + ";"
	stmt, _ := common.BaseDb.Prepare(sql)
	_, err := stmt.Exec(vals...)

	return err
}

func PageFindLottoResultList(lottoID int, pageParam common.PageParams) (*common.PageResult, *[]LottoResult, error) {
	t := LottoResult{}
	conditionSql := fmt.Sprintf(" FROM %s WHERE 1=1 ", t.TableName(lottoID))
	countSql := "SELECT COUNT(*) AS count " + conditionSql
	querySql := fmt.Sprintf("SELECT %s ", strings.Join(t.Field(), ",")) + conditionSql
	sqlWhere, args := "", []interface{}{}
	if v, ok := pageParam.Params["id"]; ok {
		sqlWhere += " AND id = ?"
		args = append(args, v)
	}
	if v, ok := pageParam.Params["lotto_id"]; ok {
		sqlWhere += " AND lotto_id = ?"
		args = append(args, v)
	}
	if v, ok := pageParam.Params["issue"]; ok {
		sqlWhere += " AND issue = ?"
		args = append(args, v)
	}
	if v, ok := pageParam.Params["draw_number"]; ok {
		sqlWhere += " AND draw_number = ?"
		args = append(args, v)
	}
	if v, ok := pageParam.Params["issue_date"]; ok {
		sqlWhere += " AND issue_date = ?"
		args = append(args, v)
	}
	if v, ok := pageParam.Params["status"]; ok {
		sqlWhere += " AND status = ?"
		args = append(args, v)
	}

	if v, ok := pageParam.Params["start_time_from"]; ok {
		sqlWhere += " AND start_time >= ?"
		args = append(args, v)
	}
	if v, ok := pageParam.Params["start_time_to"]; ok {
		sqlWhere += " AND start_time <= ?"
		args = append(args, v)
	}
	if v, ok := pageParam.Params["stop_time_from"]; ok {
		sqlWhere += " AND stop_time >= ?"
		args = append(args, v)
	}
	if v, ok := pageParam.Params["stop_time_to"]; ok {
		sqlWhere += " AND stop_time <= ?"
		args = append(args, v)
	}
	if v, ok := pageParam.Params["result_time_from"]; ok {
		sqlWhere += " AND result_time >= ?"
		args = append(args, v)
	}
	if v, ok := pageParam.Params["result_time_to"]; ok {
		sqlWhere += " AND result_time <= ?"
		args = append(args, v)
	}
	if v, ok := pageParam.Params["issue_date_from"]; ok {
		sqlWhere += " AND issue_date >= ?"
		args = append(args, v)
	}
	if v, ok := pageParam.Params["issue_date_to"]; ok {
		sqlWhere += " AND issue_date <= ?"
		args = append(args, v)
	}
	if v, ok := pageParam.Params["update_time_from"]; ok {
		sqlWhere += " AND update_time >= ?"
		args = append(args, v)
	}
	if v, ok := pageParam.Params["update_time_to"]; ok {
		sqlWhere += " AND update_time <= ?"
		args = append(args, v)
	}

	if v, ok := pageParam.Params["sort_type"]; ok {
		if v == "DESC" {
			sqlWhere += " ORDER BY issue DESC "
		} else {
			sqlWhere += " ORDER BY issue ASC "
		}
	}

	var pg common.PageResult
	var data []LottoResult
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
		d := LottoResult{}
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

func GetCurIssue(lid int) (*LottoResult, error) {
	l := LottoResult{}
	querySql := fmt.Sprintf("SELECT %s FROM %s WHERE result_time > ? ORDER BY issue ASC LIMIT 1", strings.Join(l.Field(), ","), l.TableName(lid))
	err := common.BaseDb.QueryRow(querySql, common.GetTimeNowString()).Scan(l.FieldItem()...)
	return &l, err
}
