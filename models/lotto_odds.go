package models

import (
	"github.com/shopspring/decimal"
	"github.com/xiuos/mozi/common"
	"github.com/xiuos/xorm"
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

func (lo *LottoOdds) TableName() string {
	return "lotto_odds"
}

func (lo *LottoOdds) Field() []string {
	return []string{"lotto_id", "method_code", "play_code", "method_name", "play_name", "base_odds", "odds", "odds_min", "odds_max", "bet_min", "bet_max", "status", "is_show"}
}

func (lo *LottoOdds) FieldItem() []interface{} {
	return []interface{}{&lo.LottoID, &lo.MethodCode, &lo.PlayCode, &lo.MethodName, &lo.PlayName, &lo.BaseOdds, &lo.Odds, &lo.OddsMin, &lo.OddsMax, &lo.BetMin, &lo.BetMax, &lo.Status, &lo.IsShow}
}

func GetPlayInfo(lid int, mid string, pid string) (*LottoOdds, error) {
	o := xorm.Orm{}
	lo := LottoOdds{}
	querySql, err := o.Table(lo.TableName()).Query().Where("lotto_id = ? AND method_code = ? AND play_code = ?", lid, mid, pid).Do(&lo)
	if err != nil {
		return &lo, err
	}
	err = common.BaseDb.QueryRow(querySql, o.Args()...).Scan(lo.FieldItem()...)
	return &lo, err
}
