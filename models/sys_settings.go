package models

import (
	"github.com/xiuos/mozi/common"
	"github.com/xiuos/xorm"
)

const (
	TableSysSettings = "sys_settings"
)

const (
	SysDefaultParentID = "default_parent_id"
)

type SysSettings struct {
	SysKey   string
	SysValue string
}

func (ss *SysSettings) Field() []string {
	return []string{"sys_key", "sys_value"}
}

func (ss *SysSettings) FieldItem() []interface{} {
	return []interface{}{&ss.SysKey, &ss.SysValue}
}

func GetSysSettings(key string) (*SysSettings, error) {
	o := xorm.Orm{}
	ss := SysSettings{}
	querySql, err := o.Table(TableSysSettings).Query().Where("sys_key = ?", key).Do(&ss)
	if err != nil {
		return &ss, err
	}
	err = common.BaseDb.QueryRow(querySql, o.Args()...).Scan(ss.FieldItem()...)
	return &ss, err
}
