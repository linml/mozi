package xorm

import (
	"fmt"
	"strings"
)

func (o *Orm) DoCreate(m Model) (string, error) {
	if o.table == "" {
		return "", TableMissing{}
	}
	fieldList := []string{}
	for _, v := range m.Field() {
		fieldList = append(fieldList, fmt.Sprintf("%s = ?", v))
	}
	s := fmt.Sprintf("INSERT INTO %s SET %s", o.table, strings.Join(fieldList, ","))

	return s, nil
}
