package xorm

import "errors"

type Orm struct {
	table  string
	method int

	whereStatus    int
	whereCondition interface{}
	whereArgs      []interface{}
}

func (o *Orm) Table(t string) *Orm {
	o.table = t
	return o
}

func (o *Orm) Args() []interface{} {
	return o.whereArgs
}

func (o *Orm) Create() *Orm {
	o.method = MethodCreate
	return o
}

func (o *Orm) Delete() *Orm {
	o.method = MethodDelete
	return o
}

func (o *Orm) Update() *Orm {
	o.method = MethodUpdate
	return o
}

func (o *Orm) Query() *Orm {
	o.method = MethodQuery
	return o
}

func (o *Orm) Where(query interface{}, args ...interface{}) *Orm {
	o.whereStatus = 1
	o.whereCondition = query
	o.whereArgs = args
	return o
}

// return sql format string.
func (o *Orm) Do(m Model) (string, error) {
	switch o.method {
	case MethodCreate:
		return o.DoCreate(m)
	case MethodDelete:
		return o.DoDelete(m)
	case MethodUpdate:
		return o.DoUpdate(m)
	case MethodQuery:
		return o.DoQuery(m)
	default:
		return "", errors.New("err")
	}
}
