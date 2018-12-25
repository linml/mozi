package lotto

import (
	"fmt"
	"github.com/xiuos/mozi/common"
	"strings"
)

type CMSLottoMethodGroupTemplate struct {
	ID         int    `json:"id"`
	LottoType  int    `json:"lotto_type"`
	GroupID    int    `json:"group_id"`
	GroupName  string `json:"group_name"`
	GroupAlias string `json:"group_alias"`
	SortIndex  int    `json:"sort_index"`
	Status     int    `json:"status"`
}

func (d *CMSLottoMethodGroupTemplate) TableName() string {
	return "cms_lotto_method_group_template"
}

func (d *CMSLottoMethodGroupTemplate) Field() []string {
	return []string{"id", "lotto_type", "group_id", "group_name", "group_alias", "sort_index", "status"}
}

func (d *CMSLottoMethodGroupTemplate) FieldItem() []interface{} {
	return []interface{}{&d.ID, &d.LottoType, &d.GroupID, &d.GroupName, &d.GroupAlias, &d.SortIndex, &d.Status}
}

func FindCMSLottoMethodGroupTemplateList(param map[string]string) (*[]CMSLottoMethodGroupTemplate, error) {
	c := CMSLottoMethodGroupTemplate{}
	querySql := fmt.Sprintf("SELECT %s FROM %s WHERE 1=1 ", strings.Join(c.Field(), ","), c.TableName())
	sqlWhere, args := "", []interface{}{}

	if v, ok := param["lotto_type"]; ok {
		sqlWhere += " AND lotto_type = ?"
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

	var data []CMSLottoMethodGroupTemplate
	var err error

	querySql += sqlWhere
	rows, err := common.BaseDb.Query(querySql, args...)
	if err != nil {
		return &data, err
	}
	for rows.Next() {
		d := CMSLottoMethodGroupTemplate{}
		err = rows.Scan(d.FieldItem()...)
		if err != nil {
			return &data, err
		}
		data = append(data, d)
	}

	return &data, err
}
