package lotto

import (
	"fmt"
	"github.com/xiuos/mozi/common"
	"strings"
)

type CMSLottoMethodGroup struct {
	ID         int    `json:"id"`
	LottoID    int    `json:"lotto_id"`
	GroupID    int    `json:"group_id"`
	GroupName  string `json:"group_name"`
	GroupAlias string `json:"group_alias"`
	SortIndex  int    `json:"sort_index"`
	Status     int    `json:"status"`
}

func (d *CMSLottoMethodGroup) TableName() string {
	return "cms_lotto_method_group"
}

func (d *CMSLottoMethodGroup) Field() []string {
	return []string{"id", "lotto_id", "group_id", "group_name", "group_alias", "sort_index", "status"}
}

func (d *CMSLottoMethodGroup) FieldItem() []interface{} {
	return []interface{}{&d.ID, &d.LottoID, &d.GroupID, &d.GroupName, &d.GroupAlias, &d.SortIndex, &d.Status}
}

func CreateLottoMethodGroup(l CMSLottoMethodGroup) error {
	createSql := fmt.Sprintf("INSERT IGNORE INTO %s SET lotto_id=?,group_id=?,group_name=?,group_alias=?,sort_index=?,status=?", l.TableName())
	_, err := common.BaseDb.Exec(createSql, &l.LottoID, &l.GroupID, &l.GroupName, &l.GroupAlias, &l.SortIndex, &l.Status)
	return err
}

func FindCMSLottoMethodGroupList(param map[string]string) (*[]CMSLottoMethodGroup, error) {
	c := CMSLottoMethodGroup{}
	querySql := fmt.Sprintf("SELECT %s FROM %s WHERE 1=1 ", strings.Join(c.Field(), ","), c.TableName())
	sqlWhere, args := "", []interface{}{}

	if v, ok := param["lotto_id"]; ok {
		sqlWhere += " AND lotto_id = ?"
		args = append(args, v)
	}
	if v, ok := param["group_id"]; ok {
		sqlWhere += " AND group_id = ?"
		args = append(args, v)
	}
	if v, ok := param["status"]; ok {
		sqlWhere += " AND status = ?"
		args = append(args, v)
	}

	var data []CMSLottoMethodGroup
	var err error

	querySql += sqlWhere
	rows, err := common.BaseDb.Query(querySql, args...)
	if err != nil {
		return &data, err
	}
	for rows.Next() {
		d := CMSLottoMethodGroup{}
		err = rows.Scan(d.FieldItem()...)
		if err != nil {
			return &data, err
		}
		data = append(data, d)
	}

	return &data, err
}
