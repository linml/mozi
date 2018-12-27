package lotto

import (
	"fmt"
	"github.com/shopspring/decimal"
	"github.com/xiuos/mozi/common"
	"strings"
)

type LottoOdds struct {
	LottoID    int             `json:"lotto_id"`
	MethodCode string          `json:"method_code"`
	PlayCode   string          `json:"play_code"`
	MethodName string          `json:"method_name"`
	PlayName   string          `json:"play_name"`
	BaseOdds   decimal.Decimal `json:"base_odds"`
	Odds       decimal.Decimal `json:"odds"`
	OddsMin    decimal.Decimal `json:"odds_min"`
	OddsMax    decimal.Decimal `json:"odds_max"`
	BetMin     decimal.Decimal `json:"bet_min"`
	BetMax     decimal.Decimal `json:"bet_max"`
	Status     int             `json:"status"`
	IsShow     int             `json:"is_show"`
}

func (d *LottoOdds) TableName() string {
	return "lotto_odds"
}

func (d *LottoOdds) Field() []string {
	return []string{"lotto_id", "method_code", "play_code", "method_name", "play_name", "base_odds", "odds", "odds_min", "odds_max", "bet_min", "bet_max", "status", "is_show"}
}

func (d *LottoOdds) FieldItem() []interface{} {
	return []interface{}{&d.LottoID, &d.MethodCode, &d.PlayCode, &d.MethodName, &d.PlayName, &d.BaseOdds, &d.Odds, &d.OddsMin, &d.OddsMax, &d.BetMin, &d.BetMax, &d.Status, &d.IsShow}
}

func CreateLottoOdds(lo LottoOdds) error {
	createSql := fmt.Sprintf("INSERT IGNORE INTO %s SET lotto_id=?,method_code=?,play_code=?,method_name=?,play_name=?,base_odds=?,odds=?,odds_min=?,odds_max=?,bet_min=?,bet_max=?,status=?,is_show=?", lo.TableName())
	_, err := common.BaseDb.Exec(createSql, lo.LottoID, lo.MethodCode, lo.PlayCode, lo.MethodName, lo.PlayName, lo.BaseOdds, lo.Odds, lo.OddsMin, lo.OddsMax, lo.BetMin, lo.BetMax, lo.Status, lo.IsShow)
	return err
}

func GetOdds(lid int, mCode string, pCode string) (*LottoOdds, error) {
	lo := LottoOdds{}
	querySql := fmt.Sprintf("SELECT %s FROM %s WHERE lotto_id=? AND method_code=? AND play_code=?", strings.Join(lo.Field(), ","), lo.TableName())
	err := common.BaseDb.QueryRow(querySql, lid, mCode, pCode).Scan(lo.FieldItem()...)
	return &lo, err
}

func PageFindLottoOddsList(pageParam common.PageParams) (*common.PageResult, *[]LottoOdds, error) {
	c := LottoOdds{}
	conditionSql := fmt.Sprintf(" FROM %s WHERE 1=1 ", c.TableName())
	countSql := "SELECT COUNT(*) AS count " + conditionSql
	querySql := fmt.Sprintf("SELECT %s ", strings.Join(c.Field(), ",")) + conditionSql
	sqlWhere, args := "", []interface{}{}
	if v, ok := pageParam.Params["lotto_id"]; ok {
		sqlWhere += " AND lotto_id = ?"
		args = append(args, v)
	}
	if v, ok := pageParam.Params["method_code"]; ok {
		sqlWhere += " AND method_code = ?"
		args = append(args, v)
	}
	if v, ok := pageParam.Params["play_code"]; ok {
		sqlWhere += " AND play_code = ?"
		args = append(args, v)
	}
	if v, ok := pageParam.Params["status"]; ok {
		sqlWhere += " AND status = ?"
		args = append(args, v)
	}
	if v, ok := pageParam.Params["is_show"]; ok {
		sqlWhere += " AND is_show = ?"
		args = append(args, v)
	}

	var pg common.PageResult
	var data []LottoOdds
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
		d := LottoOdds{}
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

func FindLottoOddsList(params map[string]string) (*[]LottoOdds, error) {
	c := LottoOdds{}
	querySql := fmt.Sprintf("SELECT %s FROM %s WHERE 1=1 ", strings.Join(c.Field(), ","), c.TableName())
	sqlWhere, args := "", []interface{}{}
	if v, ok := params["lotto_id"]; ok {
		sqlWhere += " AND lotto_id = ?"
		args = append(args, v)
	}
	if v, ok := params["method_code"]; ok {
		sqlWhere += " AND method_code = ?"
		args = append(args, v)
	}
	if v, ok := params["play_code"]; ok {
		sqlWhere += " AND play_code = ?"
		args = append(args, v)
	}
	if v, ok := params["status"]; ok {
		sqlWhere += " AND status = ?"
		args = append(args, v)
	}
	if v, ok := params["is_show"]; ok {
		sqlWhere += " AND is_show = ?"
		args = append(args, v)
	}

	if v, ok := params["order_by"]; ok {
		if v == "sort_index" {
			sqlWhere += " ORDER BY play_code "
		} else {
			sqlWhere += " ORDER BY lotto_id  "
		}
	} else {
		sqlWhere += " ORDER BY lotto_id "
	}

	if v, ok := params["sort_type"]; ok {
		if v == "ASC" {
			sqlWhere += " ASC "
		} else {
			sqlWhere += " DESC "
		}
	} else {
		sqlWhere += " DESC "
	}

	var data []LottoOdds
	var err error

	querySql += sqlWhere
	rows, err := common.BaseDb.Query(querySql, args...)
	if err != nil {
		fmt.Println(err)
		return &data, err
	}
	for rows.Next() {
		d := LottoOdds{}
		err = rows.Scan(d.FieldItem()...)
		if err != nil {
			return &data, err
		}
		data = append(data, d)
	}
	return &data, err
}

func SetLottoOddsInfo(lid int, mCode string, pCode string, filed string, val interface{}) error {
	u := LottoOdds{}
	updateSql := fmt.Sprintf("UPDATE %s SET %s=? WHERE lotto_id=? AND method_code=? AND play_code=?", u.TableName(), filed)
	_, err := common.BaseDb.Exec(updateSql, val, lid, mCode, pCode)
	return err

}
