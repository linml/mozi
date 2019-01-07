package models

import (
	"github.com/xiuos/mozi/common"
	"github.com/xiuos/mozi/xorm"
)

const (
	SysDefaultParentID = "default_parent_id"
	SysLottoCalcMyself = "lotto_calc_myself"
)

type SysSettings struct {
	SysKey   string
	SysValue string
}

func (ss *SysSettings) TableName() string {
	return "sys_settings"
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
	querySql, err := o.Table(ss.TableName()).Query().Where("sys_key = ?", key).Do(&ss)
	if err != nil {
		return &ss, err
	}
	err = common.BaseDb.QueryRow(querySql, o.Args()...).Scan(ss.FieldItem()...)
	return &ss, err
}
