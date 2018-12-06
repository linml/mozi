package lotto

import (
	"fmt"
	"github.com/xiuos/mozi/common"
	"github.com/xiuos/xorm"
	"strings"
)

type CodeLotto struct {
	LottoID      int    `json:"lotto_id"`
	Name         string `json:"name"`
	LottoType    int    `json:"lotto_type"`
	GameKind     int    `json:"game_kind"`
	GameType     int    `json:"game_type"`
	Status       int    `json:"status"`
	SortIndex    int    `json:"sort_index"`
	Introduction string `json:"introduction"`
}

func (l *CodeLotto) TableName() string {
	return "code_lotto"
}

func (l *CodeLotto) Field() []string {
	return []string{"lotto_id", "name", "lotto_type", "game_kind", "game_type", "status", "sort_index", "introduction"}
}

func (l *CodeLotto) FieldItem() []interface{} {
	return []interface{}{&l.LottoID, &l.Name, &l.LottoType, &l.GameKind, &l.GameType, &l.Status, &l.SortIndex, &l.Introduction}
}

func GetLotto(lid int) (*CodeLotto, error) {
	o := xorm.Orm{}
	l := CodeLotto{}
	querySql, err := o.Table(l.TableName()).Query().Where("lotto_id = ?", lid).Do(&l)
	if err != nil {
		return &l, err
	}
	err = common.BaseDb.QueryRow(querySql, o.Args()...).Scan(l.FieldItem()...)
	return &l, err
}

func PageFindCodeLottoList(pageParam common.PageParams) (*common.PageResult, *[]CodeLotto, error) {
	c := CodeLotto{}
	conditionSql := fmt.Sprintf(" FROM %s WHERE 1=1 ", c.TableName())
	countSql := "SELECT COUNT(*) AS count " + conditionSql
	querySql := fmt.Sprintf("SELECT %s ", strings.Join(c.Field(), ",")) + conditionSql
	sqlWhere, args := "", []interface{}{}
	if v, ok := pageParam.Params["lotto_id"]; ok {
		sqlWhere += " AND lotto_id = ?"
		args = append(args, v)
	}
	if v, ok := pageParam.Params["name"]; ok {
		sqlWhere += " AND name = ?"
		args = append(args, v)
	}
	if v, ok := pageParam.Params["lotto_type"]; ok {
		sqlWhere += " AND lotto_type = ?"
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
	if v, ok := pageParam.Params["status"]; ok {
		sqlWhere += " AND status = ?"
		args = append(args, v)
	}

	var pg common.PageResult
	var data []CodeLotto
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
		d := CodeLotto{}
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

func FindCodeLottoList(param map[string]string) (*[]CodeLotto, error) {
	arl := []CodeLotto{}
	t := CodeLotto{}
	querySql := fmt.Sprintf("SELECT %s FROM %s WHERE 1=1", strings.Join(t.Field(), ","), t.TableName())
	sqlWhere, args := "", []interface{}{}
	if v, ok := param["lotto_id"]; ok {
		sqlWhere += " AND lotto_id = ?"
		args = append(args, v)
	}
	if v, ok := param["name"]; ok {
		sqlWhere += " AND name = ?"
		args = append(args, v)
	}
	if v, ok := param["lotto_type"]; ok {
		sqlWhere += " AND lotto_type = ?"
		args = append(args, v)
	}
	if v, ok := param["game_kind"]; ok {
		sqlWhere += " AND game_kind = ?"
		args = append(args, v)
	}
	if v, ok := param["game_type"]; ok {
		sqlWhere += " AND game_type = ?"
		args = append(args, v)
	}
	if v, ok := param["status"]; ok {
		sqlWhere += " AND status = ?"
		args = append(args, v)
	}
	querySql += sqlWhere
	rows, err := common.BaseDb.Query(querySql, args...)
	if err != nil {
		fmt.Println(err)
		return &arl, err
	}
	for rows.Next() {
		am := CodeLotto{}
		err = rows.Scan(am.FieldItem()...)
		if err != nil {
			return &arl, err
		}
		arl = append(arl, am)
	}
	return &arl, nil
}
