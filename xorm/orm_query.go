package xorm

import (
	"fmt"
	"strings"
)

func (o *Orm) DoQuery(m Model) (string, error) {
	if o.table == "" {
		return "", TableMissing{}
	}

	s := fmt.Sprintf("SELECT %s FROM %s", strings.Join(m.Field(), ","), o.table)
	if o.whereStatus == 1 {
		s += fmt.Sprintf(" WHERE %s", o.whereCondition)
	}
	return s, nil
}
