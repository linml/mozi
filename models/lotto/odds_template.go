package lotto

import (
	"fmt"
	"github.com/shopspring/decimal"
	"github.com/xiuos/mozi/common"
	"strings"
)

type LottoOddsTemplate struct {
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
	SortIndex  int             `json:"sort_index"`
	IsShow     int             `json:"is_show"`
}

func (d *LottoOddsTemplate) TableName() string {
	return "lotto_odds_template"
}

func (d *LottoOddsTemplate) Field() []string {
	return []string{"method_code", "play_code", "method_name", "play_name", "base_odds", "odds", "odds_min", "odds_max", "bet_min", "bet_max", "status", "sort_index", "is_show"}
}

func (d *LottoOddsTemplate) FieldItem() []interface{} {
	return []interface{}{&d.MethodCode, &d.PlayCode, &d.MethodName, &d.PlayName, &d.BaseOdds, &d.Odds, &d.OddsMin, &d.OddsMax, &d.BetMin, &d.BetMax, &d.Status, &d.SortIndex, &d.IsShow}
}

func FindLottoOddsTemplateList(param map[string]string) (*[]LottoOddsTemplate, error) {
	c := LottoOddsTemplate{}
	querySql := fmt.Sprintf("SELECT %s FROM %s WHERE 1=1 ", strings.Join(c.Field(), ","), c.TableName())
	sqlWhere, args := "", []interface{}{}

	if v, ok := param["method_code"]; ok {
		sqlWhere += " AND method_code = ?"
		args = append(args, v)
	}
	if v, ok := param["play_code"]; ok {
		sqlWhere += " AND play_code = ?"
		args = append(args, v)
	}
	if v, ok := param["status"]; ok {
		sqlWhere += " AND status = ?"
		args = append(args, v)
	}
	if v, ok := param["is_show"]; ok {
		sqlWhere += " AND is_show = ?"
		args = append(args, v)
	}

	var data []LottoOddsTemplate
	var err error

	querySql += sqlWhere
	rows, err := common.BaseDb.Query(querySql, args...)
	if err != nil {
		return &data, err
	}
	for rows.Next() {
		d := LottoOddsTemplate{}
		err = rows.Scan(d.FieldItem()...)
		if err != nil {
			return &data, err
		}
		data = append(data, d)
	}

	return &data, err
}
