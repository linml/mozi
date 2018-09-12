package models

import (
	"github.com/xiuos/lotto"
	"github.com/xiuos/mozi/common"
	"github.com/xiuos/xorm"
)

const (
	TableLottoResult = "lotto_result"
)

type IssueInfo = lotto.IssueInfo

type LottoResult struct {
	LottoID    int    `json:"lotto_id"`
	Issue      string `json:"issue"`
	DrawNumber string `json:"draw_number"`
	Status     int    `json:"status"`
	StartTime  string `json:"start_time"`
	CloseTime  string `json:"close_time"`
	EndTime    string `json:"end_time"`
	DrawTime   string `json:"draw_time"`
}

func (l *LottoResult) Field() []string {
	return []string{"lotto_id", "issue", "draw_number", "status", "start_time", "close_time", "end_time", "draw_time"}
}

func (l *LottoResult) FieldItem() []interface{} {
	return []interface{}{l.LottoID, l.Issue, l.DrawNumber, l.Status, l.StartTime, l.CloseTime, l.EndTime, l.DrawTime}
}

func GetLottoResult(lid int, issue string) (*LottoResult, error) {
	o := xorm.Orm{}
	l := LottoResult{}
	querySql, err := o.Table(TableLottoResult).Query().Where("lotto_id = ? AND issue = ?", lid, issue).Do(&l)
	if err != nil {
		return &l, err
	}
	err = common.BaseDb.QueryRow(querySql, o.Args()...).Scan(l.FieldItem()...)
	return &l, err
}

func GenIssueInfo(data *[]IssueInfo) error {
	var vals []interface{}
	sql := "INSERT IGNORE INTO lotto_result (lotto_id, issue, draw_number, status, start_time, close_time, end_time, draw_time) VALUES "
	for i, _ := range *data {
		sql += "(?, ?, ?, ?, ?, ?, ?, ?),"
		vals = append(vals, (*data)[i].LottoID, (*data)[i].Issue, "", 0, (*data)[i].StartTime, (*data)[i].CloseTime, (*data)[i].EndTime, "")
	}
	sql = sql[0:len(sql)-1] + ";"
	stmt, _ := common.BaseDb.Prepare(sql)
	_, err := stmt.Exec(vals...)

	return err
}
