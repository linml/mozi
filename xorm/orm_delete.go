package xorm

import (
	"fmt"
)

func (o *Orm) DoDelete(m Model) (string, error) {
	if o.table == "" {
		return "", TableMissing{}
	}
	s := fmt.Sprintf("DELETE FROM %s", o.table)
	if o.whereStatus == 1 {
		s += fmt.Sprintf(" WHERE %s", o.whereCondition)
	}
	return s, nil
}
