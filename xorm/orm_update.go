package xorm

import (
	"fmt"
)

func (o *Orm) DoUpdate(m Model) (string, error) {
	if o.table == "" {
		return "", TableMissing{}
	}
	s := fmt.Sprintf("UPDATE %s", o.table)
	if o.whereStatus == 1 {
		s += fmt.Sprintf(" WHERE %s", o.whereCondition)
	}
	return s, nil
}
