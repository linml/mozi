package lotto

import (
	"fmt"
	"github.com/xiuos/mozi/common"
	"strings"
)

type LottoMethodTemplate struct {
	LottoType  int    `json:"lotto_type"`
	MethodCode string `json:"method_code"`
	Name       string `json:"name"`
	OddsType   int    `json:"odds_type"`
	Status     int    `json:"status"`
}

func (d *LottoMethodTemplate) TableName() string {
	return "code_method_template"
}

func (d *LottoMethodTemplate) Field() []string {
	return []string{"lotto_type", "method_code", "name", "odds_type", "status"}
}

func (d *LottoMethodTemplate) FieldItem() []interface{} {
	return []interface{}{&d.LottoType, &d.MethodCode, &d.Name, &d.OddsType, &d.Status}
}

func FindLottoMethodTemplateList(param map[string]string) (*[]LottoMethodTemplate, error) {
	c := LottoMethodTemplate{}
	querySql := fmt.Sprintf("SELECT %s FROM %s WHERE 1=1 ", strings.Join(c.Field(), ","), c.TableName())
	sqlWhere, args := "", []interface{}{}

	if v, ok := param["lotto_type"]; ok {
		sqlWhere += " AND lotto_type = ?"
		args = append(args, v)
	}
	if v, ok := param["method_code"]; ok {
		sqlWhere += " AND method_code = ?"
		args = append(args, v)
	}
	if v, ok := param["status"]; ok {
		sqlWhere += " AND status = ?"
		args = append(args, v)
	}

	var data []LottoMethodTemplate
	var err error

	querySql += sqlWhere
	rows, err := common.BaseDb.Query(querySql, args...)
	if err != nil {
		return &data, err
	}
	for rows.Next() {
		d := LottoMethodTemplate{}
		err = rows.Scan(d.FieldItem()...)
		if err != nil {
			return &data, err
		}
		data = append(data, d)
	}

	return &data, err
}

func GEtLottoMethodTemplate(lottoType int, methodCode string) (*LottoMethodTemplate, error) {
	c := LottoMethodTemplate{}
	querySql := fmt.Sprintf("SELECT %s FROM %s WHERE lotto_type=? AND method_code=? ", strings.Join(c.Field(), ","), c.TableName())
	err := common.BaseDb.QueryRow(querySql, lottoType, methodCode).Scan(c.FieldItem()...)
	return &c, err
}
