package xorm

import "fmt"

type TableMissing struct {
}

func (err TableMissing) Error() string {
	return fmt.Sprintf("Table missing")
}
