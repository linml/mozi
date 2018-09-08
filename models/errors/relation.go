package errors

import "fmt"

type RefNotFound struct {
	Ref int
}

func (err RefNotFound)Error() string {
	return fmt.Sprintf("找不到推广:%d 信息", err.Ref)
}

