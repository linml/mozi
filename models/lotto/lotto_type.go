package lotto

import (
	"fmt"
	"github.com/xiuos/mozi/common"
	"strings"
)

type CodeLottoType struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
}

func (c *CodeLottoType) TableName() string {
	return "code_lotto_type"
}

func (c *CodeLottoType) Field() []string {
	return []string{"id", "name"}
}

func (c *CodeLottoType) FieldItem() []interface{} {
	return []interface{}{&c.ID, &c.Name}
}

func FindCodeLottoTypeList(param map[string]string) (*[]CodeLottoType, error) {
	arl := []CodeLottoType{}
	t := CodeLottoType{}
	querySql := fmt.Sprintf("SELECT %s FROM %s WHERE 1=1", strings.Join(t.Field(), ","), t.TableName())
	sqlWhere, args := "", []interface{}{}
	if v, ok := param["id"]; ok {
		sqlWhere += " AND id = ?"
		args = append(args, v)
	}
	if v, ok := param["name"]; ok {
		sqlWhere += " AND name = ?"
		args = append(args, v)
	}

	querySql += sqlWhere
	rows, err := common.BaseDb.Query(querySql, args...)
	if err != nil {
		fmt.Println(err)
		return &arl, err
	}
	for rows.Next() {
		am := CodeLottoType{}
		err = rows.Scan(am.FieldItem()...)
		if err != nil {
			return &arl, err
		}
		arl = append(arl, am)
	}
	return &arl, nil
}
