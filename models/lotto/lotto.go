package lotto

import (
	"fmt"
	"github.com/xiuos/mozi/common"
	"github.com/xiuos/mozi/xorm"
	"strings"
)

const (
	CQSSC   = 1  //	重庆时时彩[信]
	XJSSC   = 2  //	新疆时时彩[信]
	TJSSC   = 3  //	天津时时彩[信]
	MZ1F    = 4  //	分分彩[信][自开]
	MZ2F    = 5  //	两分彩[信][自开]
	MZ3F    = 6  //	三分彩[信][自开]
	MZ5F    = 7  //	五分彩[信][自开]
	SD11X5  = 8  //	山东11选5[信]
	JX11X5  = 9  //	江西11选5[信]
	GD11X5  = 10 //	广东11选5[信]
	JS11X5  = 11 //	江苏11选5[信]
	AH11X5  = 12 //	安徽11选5[信]
	SX11X5  = 13 //	山西11选5[信]
	SH11X5  = 14 //	上海11选5[信]
	MZ11X5  = 15 //	分分11选5[信][自开]
	JSK3    = 16 //	江苏快3[信]
	AHK3    = 17 //	安徽快3[信]
	HBK3    = 18 //	湖北快3[信]
	HNK3    = 19 //	河南快3[信]
	JSTB    = 20 //	江苏骰宝[信]
	MZK3    = 21 //	分分快3[信][自开]
	BJPK10  = 22 //	北京PK10[信]
	XYFT    = 23 //	幸运飞艇[信]
	MZPK10  = 24 //	分分PK10[信][自开]
	FC3D    = 25 //	福彩3D[信]
	PL3     = 26 //	排列3[信]
	PL5     = 27 //	排列5[信]
	GDKLSF  = 28 //	广东快乐十分[信]
	CQKLSF  = 29 //	重庆快乐十分[信]
	TJKLSF  = 30 //	天津快乐十分[信]
	BJKL8   = 31 //	北京快乐8[信]
	JND     = 32 //	加拿大基诺[信]
	MZKLC   = 33 //	分分快乐彩[信][自开]
	TWBG    = 34 //	台湾宾果[信]
	BJXY28  = 35 //	北京幸运28[信]
	JNDXY28 = 36 //	加拿大幸运28[信]
	TWXY28  = 37 //	台湾幸运28[信]
	LHC     = 38 //	六合彩[信]
	MZLHC   = 39 //	五分六合彩[信][自开]
	MZ10LHC = 40 //	十分六合彩[信][自开]
)

type CodeLotto struct {
	LottoID         int    `json:"lotto_id"`
	Name            string `json:"name"`
	LottoType       int    `json:"lotto_type"`
	GameKind        int    `json:"game_kind"`
	GameType        int    `json:"game_type"`
	Status          int    `json:"status"`
	SortIndex       int    `json:"sort_index"`
	IsShow          int    `json:"is_show"`
	Introduction    string `json:"introduction"`
	ImgUrl          string `json:"img_url"`
	Extra1SortIndex int    `json:"extra_1_sort_index"` // 大厅
	Extra2SortIndex int    `json:"extra_2_sort_index"` // 首页
	Extra3SortIndex int    `json:"extra_3_sort_index"` // 结果
}

func (l *CodeLotto) TableName() string {
	return "code_lotto"
}

func (l *CodeLotto) Field() []string {
	return []string{"lotto_id", "name", "lotto_type", "game_kind", "game_type", "status", "sort_index", "is_show", "introduction", "img_url", "extra_1_sort_index", "extra_2_sort_index", "extra_3_sort_index"}
}

func (l *CodeLotto) FieldItem() []interface{} {
	return []interface{}{&l.LottoID, &l.Name, &l.LottoType, &l.GameKind, &l.GameType, &l.Status, &l.SortIndex, &l.IsShow, &l.Introduction, &l.ImgUrl, &l.Extra1SortIndex, &l.Extra2SortIndex, &l.Extra3SortIndex}
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

func SetCodeLottoInfo(lid int, filed string, val interface{}) error {
	u := CodeLotto{}
	updateSql := fmt.Sprintf("UPDATE %s SET %s=? WHERE lotto_id=?", u.TableName(), filed)
	_, err := common.BaseDb.Exec(updateSql, val, lid)
	return err
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
	if v, ok := pageParam.Params["is_show"]; ok {
		sqlWhere += " AND is_show = ?"
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
	if v, ok := param["is_show"]; ok {
		sqlWhere += " AND is_show = ?"
		args = append(args, v)
	}

	if v, ok := param["order_by"]; ok {
		f, val := true, ""
		switch v {
		case "sort_index":
			val = " ORDER BY sort_index "
		case "extra_1_sort_index":
			val = " ORDER BY extra_1_sort_index "
		case "extra_2_sort_index":
			val = " ORDER BY extra_2_sort_index "
		case "extra_3_sort_index":
			val = " ORDER BY extra_3_sort_index "
		default:
			f = false
		}
		sqlWhere += val
		if f == true {
			if v, ok := param["sort_type"]; ok {
				if v == "ASC" {
					sqlWhere += " ASC "
				} else {
					sqlWhere += " DESC "
				}
			}
		}
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
