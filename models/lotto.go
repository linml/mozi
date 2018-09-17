package models

import (
	"github.com/xiuos/mozi/common"
	"github.com/xiuos/xorm"
)

type Lotto struct {
	LottoID      int    `json:"lotto_id"`
	Name         string `json:"name"`
	LottoType    int    `json:"lotto_type"`
	GameKind     int    `json:"game_kind"`
	GameType     int    `json:"game_type"`
	Status       int    `json:"status"`
	SortIndex    int    `json:"sort_index"`
	Introduction string `json:"introduction"`
}

func (l *Lotto) TableName() string {
	return "code_lotto"
}

func (l *Lotto) Field() []string {
	return []string{"lotto_id", "name", "lotto_type", "game_kind", "game_type", "status", "sort_index", "introduction"}
}

func (l *Lotto) FieldItem() []interface{} {
	return []interface{}{&l.LottoID, &l.Name, &l.LottoType, &l.GameKind, &l.GameType, &l.Status, &l.SortIndex, &l.Introduction}
}

func GetLotto(lid int) (*Lotto, error) {
	o := xorm.Orm{}
	l := Lotto{}
	querySql, err := o.Table(l.TableName()).Query().Where("lotto_id = ?", lid).Do(&l)
	if err != nil {
		return &l, err
	}
	err = common.BaseDb.QueryRow(querySql, o.Args()...).Scan(l.FieldItem()...)
	return &l, err
}
