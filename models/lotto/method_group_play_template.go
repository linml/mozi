package lotto

import (
	"fmt"
	"github.com/xiuos/mozi/common"
	"strings"
)

type CMSLottoMethodGroupPlayTemplate struct {
	ID          int    `json:"id"`
	GroupID     int    `json:"group_id"`
	MethodCode  string `json:"method_code"`
	MethodName  string `json:"method_name"`
	MethodAlias string `json:"method_alias"`
	SortIndex   int    `json:"sort_index"`
	Status      int    `json:"status"`
}

func (d *CMSLottoMethodGroupPlayTemplate) TableName() string {
	return "cms_lotto_method_group_play_template"
}

func (d *CMSLottoMethodGroupPlayTemplate) Field() []string {
	return []string{"id", "group_id", "method_code", "method_name", "method_alias", "sort_index", "status"}
}

func (d *CMSLottoMethodGroupPlayTemplate) FieldItem() []interface{} {
	return []interface{}{&d.ID, &d.GroupID, &d.MethodCode, &d.MethodName, &d.MethodAlias, &d.SortIndex, &d.Status}
}

func FindCMSLottoMethodGroupPlayTemplateList(param map[string]string) (*[]CMSLottoMethodGroupPlayTemplate, error) {
	c := CMSLottoMethodGroupPlayTemplate{}
	querySql := fmt.Sprintf("SELECT %s FROM %s WHERE 1=1 ", strings.Join(c.Field(), ","), c.TableName())
	sqlWhere, args := "", []interface{}{}

	if v, ok := param["group_id"]; ok {
		sqlWhere += " AND group_id = ?"
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

	var data []CMSLottoMethodGroupPlayTemplate
	var err error

	querySql += sqlWhere
	rows, err := common.BaseDb.Query(querySql, args...)
	if err != nil {
		return &data, err
	}
	for rows.Next() {
		d := CMSLottoMethodGroupPlayTemplate{}
		err = rows.Scan(d.FieldItem()...)
		if err != nil {
			return &data, err
		}
		data = append(data, d)
	}

	return &data, err
}
